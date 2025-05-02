package app

import (
	"database/sql"
	"gh_static_portfolio/internal/app/handlers"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/features/auth"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/home"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/termoccasion"
	"gh_static_portfolio/internal/features/unit"
	"gh_static_portfolio/internal/features/user"
	markdownrenderer "gh_static_portfolio/internal/infrastructure/goldmarkrenderer"
	"gh_static_portfolio/internal/infrastructure/localfilesystem"
	"gh_static_portfolio/internal/infrastructure/sqlite"
	authentication "gh_static_portfolio/internal/newauthentication"
	"gh_static_portfolio/internal/shared/routes"
	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
	db   *sql.DB
}

func New() (*App, error) {
	e := echo.New()
	e.Use(logger)
	e.Debug = true
	assets.RegisterStatic(e)
	queries, db, err := sqlite.InitDB("./course_manager.db")
	if err != nil {
		return nil, err
	}
	root := e.Group("")
	protected := root.Group(
		"",
		authentication.AddCookieToHeader,
		authentication.JWTMiddlewareProtectedNew(e.Reverse(routes.GetSignin.String())),
		authentication.GetClaims,
	)

	homeService := home.NewService(struct{}{})
	homeHandler := home.NewHandler(*homeService, e)
	err = home.RegisterRoutes(root, homeHandler)
	if err != nil {
		return nil, err
	}
	// repositories
	userRepo := sqlite.NewUserRepo(queries)
	termRepo := sqlite.NewTermRepo(queries)
	courseRepo := sqlite.NewCourseRepo(queries)
	unitRepo := sqlite.NewUnitRepo(queries)
	lessonRepo := sqlite.NewLessonRepo(queries)
	termOccasionRepo := sqlite.NewTermOccasionRepo(queries)

	// feature-level services
	authService := auth.NewService(struct{}{}, userRepo)
	userService := user.NewService(userRepo)
	termService := term.NewService(termRepo)
	courseService := course.NewService(courseRepo)
	unitService := unit.NewService(unitRepo)
	lessonService := lesson.NewService(lessonRepo)
	termOccasionService := termoccasion.New(termOccasionRepo)

	// feature-level handlers
	authHandler := auth.NewHandler(authService, e)
	userHandler := user.NewHandler(userService, e)
	termHandler := term.NewHandler(termService)
	courseHandler := course.NewHandler(courseService)
	unitHandler := unit.NewHandler(unitService)
	lessonHandler := lesson.NewHandler(lessonService)

	// register feature-level routes
	err = auth.RegisterRoutes(root, authHandler)
	if err != nil {
		return nil, err
	}
	err = user.RegisterRoutes(protected, userHandler)
	if err != nil {
		return nil, err
	}
	err = term.RegisterRoutes(protected, termHandler)
	if err != nil {
		return nil, err
	}
	err = course.RegisterRoutes(protected, courseHandler)
	if err != nil {
		return nil, err
	}
	err = unit.RegisterRoutes(protected, unitHandler)
	if err != nil {
		return nil, err
	}
	err = lesson.RegisterRoutes(protected, lessonHandler)
	if err != nil {
		return nil, err
	}

	// file system and markdown renderer
	filesRepo := localfilesystem.New()
	markdownRenderer := markdownrenderer.New()

	// this is used by other application services, though it is also in the application layer (not ideal)
	fileService := services.NewFileService(filesRepo, markdownRenderer)

	// application-level services
	lessonAppService := services.NewLessonService(lessonService)
	unitAppService := services.NewUnitService(unitService, lessonService)
	courseAppService := services.NewCourseService(courseService, unitAppService.ByCourseID)
	termAppService := services.NewTermService(termService, termOccasionService, courseService)
	userAppService := services.NewUserService(userService, termService)
	nodeAppService := services.NewNodeService(userAppService.ByID, termAppService.ByID, courseAppService.ByID, unitAppService.ByID, lessonAppService.ByID)
	termCalService := services.NewTermCalendarService(termAppService.ByID, termOccasionService)
	courseCalAppService := services.NewCourseCalendarService(courseAppService.ByID, lessonAppService.ByUnitID, termAppService.ByID)

	// application-level handlers
	lessonAppHandler := handlers.NewLessonHandler(lessonAppService, nodeAppService, e.Reverse)
	unitAppHandler := handlers.NewUnitHandler(unitAppService, nodeAppService, e.Reverse)
	courseAppHandler := handlers.NewCourseHandler(courseAppService, nodeAppService, e.Reverse)
	courseCalAppHandler := handlers.NewCourseCalHandler(courseCalAppService, nodeAppService, e.Reverse)
	termAppHandler := handlers.NewTermHandler(termAppService, nodeAppService, fileService, e.Reverse)
	termCalHandler := handlers.NewTermCalHandler(termCalService, nodeAppService, e.Reverse)
	userAppHandler := handlers.NewUserHandler(userAppService, nodeAppService, e.Reverse)

	// register application-level routes
	err = handlers.RegisterUserRoutes(protected, userAppHandler)
	if err != nil {
		return nil, err
	}
	err = handlers.RegisterTermRoutes(protected, termAppHandler)
	if err != nil {
		return nil, err
	}
	err = handlers.RegisterTermCalRoutes(protected, termCalHandler)
	if err != nil {
		return nil, err
	}
	err = handlers.RegisterCourseRoutes(protected, courseAppHandler)
	if err != nil {
		return nil, err
	}
	err = handlers.RegisterCourseCalRoutes(protected, courseCalAppHandler)
	if err != nil {
		return nil, err
	}
	err = handlers.RegisterUnitRoutes(protected, unitAppHandler)
	if err != nil {
		return nil, err
	}
	err = handlers.RegisterLessonRoutes(protected, lessonAppHandler)
	if err != nil {
		return nil, err
	}
	return &App{
		Echo: e,
		db:   db,
	}, nil

}

func (app *App) Start(url string) error {
	defer app.db.Close()
	for _, route := range app.Echo.Routes() {
		log.Printf("\nMethod: %s\nName: %s\nPath: %s\n", route.Method, route.Name, route.Path)
	}
	return app.Echo.Start(url)
}

package app

import (
	"database/sql"
	"gh_static_portfolio/internal/app/handlers"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/features/auth"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/files"
	"gh_static_portfolio/internal/features/home"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/features/markdown"
	"gh_static_portfolio/internal/features/sitegen"
	"gh_static_portfolio/internal/features/slides"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/termoccasion"
	"gh_static_portfolio/internal/features/unit"
	"gh_static_portfolio/internal/features/user"
	"gh_static_portfolio/internal/infrastructure/hugo"
	"gh_static_portfolio/internal/infrastructure/localfilesystem"
	"gh_static_portfolio/internal/infrastructure/pathing"
	slidesrenderer "gh_static_portfolio/internal/infrastructure/slides"
	"gh_static_portfolio/internal/infrastructure/sqlite"
	"gh_static_portfolio/internal/shared/routes"
	"log"

	"github.com/labstack/echo/v4"
)

const dbPath = "./course_manager.db"
const dataFilesRoot = "./internal/data/users"
const staticSitesRoot = "./hugosites"

type App struct {
	Echo *echo.Echo
	db   *sql.DB
}

type NewAppParams struct {
	MarpBaseURL string
}

func New(params NewAppParams) (*App, error) {
	e := echo.New()
	e.Use(logger)
	e.Debug = true
	assets.RegisterStatic(e)
	queries, db, err := sqlite.InitDB(dbPath)
	if err != nil {
		return nil, err
	}

	// repositories
	filesRepo := localfilesystem.New()
	userRepo := sqlite.NewUserRepo(queries)
	termRepo := sqlite.NewTermRepo(queries)
	courseRepo := sqlite.NewCourseRepo(queries)
	unitRepo := sqlite.NewUnitRepo(queries)
	lessonRepo := sqlite.NewLessonRepo(queries)
	termOccasionRepo := sqlite.NewTermOccasionRepo(queries)

	// infrastructure
	dataFilesPathingSvc := pathing.NewNodePathService(dataFilesRoot)
	staticSiteDataPathingSvc := pathing.NewNodePathService(staticSitesRoot)
	markdownRenderer := markdown.NewService()
	slidesRenderer := slidesrenderer.New(params.MarpBaseURL, dataFilesPathingSvc)

	// feature services
	authService := auth.NewService(userRepo)
	userService := user.NewService(userRepo)
	termService := term.NewService(termRepo)
	courseService := course.NewService(courseRepo)
	unitService := unit.NewService(unitRepo)
	lessonService := lesson.NewService(lessonRepo)
	termOccasionService := termoccasion.NewService(termOccasionRepo)
	fileSystem := files.NewFileService(filesRepo, dataFilesPathingSvc)
	slidesService := slides.New(params.MarpBaseURL, slidesRenderer, dataFilesPathingSvc, filesRepo)

	// application-level services
	lessonAppService := services.NewLessonService(lessonService, slidesService)
	unitAppService := services.NewUnitService(unitService)
	courseAppService := services.NewCourseService(courseService)
	termAppService := services.NewTermService(termService, termOccasionService)
	userAppService := services.NewUserService(userService)
	nodeAppService := services.NewNodeService(userAppService.ByID, termAppService.WithOccasions, courseAppService.ByID, unitAppService.ByID, lessonAppService.ByID)
	termCalService := services.NewTermCalendarService(termAppService.WithOccasions, termOccasionService)
	courseCalAppService := services.NewCourseCalendarService(courseAppService.ByID, unitAppService.ByCourseID, lessonAppService.ByUnitID, termAppService.WithOccasions)

	// infrastructure init
	hugoParams := hugo.Params{
		HugoURL:                  "hugo",
		SitesRootDir:             "hugosites",
		DataFilesRoot:            dataFilesRoot,
		StaticSitePathingService: staticSiteDataPathingSvc,
		DataPathingService:       dataFilesPathingSvc,
		GetUser:                  userAppService.ByID,
		GetTerm:                  termAppService.ByID,
		GetCourses:               courseAppService.ListByTerm,
		GetUnits:                 unitAppService.ByCourseID,
		GetLessons:               lessonAppService.ByUnitID,
	}
	hugoGenerator, err := hugo.New(hugoParams)
	if err != nil {
		return nil, err
	}

	// site generator service
	siteGenerator := sitegen.New(hugoGenerator)

	// route groups
	root := e.Group("")
	homeHandler := home.NewHandler(e.Reverse)
	err = home.RegisterRoutes(root, homeHandler)
	if err != nil {
		return nil, err
	}
	authHandler := handlers.NewAuthHandler(authService, e.Reverse)
	err = handlers.RegisterAuthRoutes(root, authHandler)
	if err != nil {
		return nil, err
	}

	protected := root.Group(
		"",
		authService.AddCookieToHeader,
		authService.JWTMiddlewareProtectedNew(e.Reverse(routes.GetSignin.String())),
		authService.GetClaims,
	)

	// application-level handlers
	lessonAppHandler := handlers.NewLessonHandler(lessonAppService, nodeAppService, e.Reverse)
	unitAppHandler := handlers.NewUnitHandler(unitAppService, nodeAppService, e.Reverse)
	courseAppHandler := handlers.NewCourseHandler(courseAppService, nodeAppService, fileSystem, e.Reverse)
	courseCalAppHandler := handlers.NewCourseCalHandler(courseCalAppService, nodeAppService, e.Reverse)
	termAppHandler := handlers.NewTermHandler(termAppService, nodeAppService, fileSystem, markdownRenderer, e.Reverse)
	termCalHandler := handlers.NewTermCalHandler(termCalService, nodeAppService, e.Reverse)
	dashboardAppHandler := handlers.NewDashboardHandler(siteGenerator, userAppService, nodeAppService, e.Reverse)

	// register application-level routes
	err = handlers.RegisterDashboardRoutes(protected, dashboardAppHandler)
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

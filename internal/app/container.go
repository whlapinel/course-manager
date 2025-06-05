package app

import (
	"database/sql"
	"gh_static_portfolio/internal/app/handlers"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/courseoccasion"
	"gh_static_portfolio/internal/features/home"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/termoccasion"
	"gh_static_portfolio/internal/features/unit"
	"gh_static_portfolio/internal/features/user"
	"gh_static_portfolio/internal/infrastructure/hugo"
	"gh_static_portfolio/internal/infrastructure/localfilesystem"
	"gh_static_portfolio/internal/infrastructure/markdown"
	"gh_static_portfolio/internal/infrastructure/marpclient"
	"gh_static_portfolio/internal/infrastructure/pathing"
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
	Domain      string
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
	courseOccasionRepo := sqlite.NewCourseOccasionRepo(queries)

	// infrastructure
	dataFilesPathingSvc := pathing.NewNodePathService(dataFilesRoot)
	staticSiteDataPathingSvc := pathing.NewNodePathService(staticSitesRoot)
	markdownRenderer := markdown.New()
	marp := marpclient.New(params.MarpBaseURL)

	// feature services
	userService := user.NewService(userRepo)
	termService := term.NewService(termRepo)
	courseService := course.NewService(courseRepo)
	unitService := unit.NewService(unitRepo)
	lessonService := lesson.NewService(lessonRepo)
	termOccasionService := termoccasion.NewService(termOccasionRepo)
	courseOccasionService := courseoccasion.NewService(courseOccasionRepo)

	// application-level services
	authService := services.NewAuthService(userRepo)
	fileService := services.NewFileService(filesRepo, dataFilesPathingSvc)
	slidesService := services.NewSlidesService(params.MarpBaseURL, marp, dataFilesPathingSvc, filesRepo)
	lessonAppService := services.NewLessonService(lessonService)
	unitAppService := services.NewUnitService(unitService)
	courseAppService := services.NewCourseService(courseService)
	termAppService := services.NewTermService(termService, termOccasionService)
	userAppService := services.NewUserService(userService)
	nodeAppService := services.NewNodeService(userAppService.ByID, termAppService.WithOccasions, courseAppService.ByID, unitAppService.ByID, lessonAppService.ByID)
	termCalService := services.NewTermCalendarService(termAppService.WithOccasions, termOccasionService)
	courseCalAppService := services.NewCourseCalendarService(
		termAppService,
		courseAppService,
		unitAppService,
		lessonAppService,
		courseOccasionService,
		lessonService,
	)
	markdownService := services.NewMarkdownService(markdownRenderer, dataFilesPathingSvc, filesRepo)

	// infrastructure init
	hugoParams := hugo.Params{
		Domain:                   params.Domain,
		HugoURL:                  "hugo",
		SitesRootDir:             "hugosites",
		DataFilesRoot:            dataFilesRoot,
		StaticSitePathingService: staticSiteDataPathingSvc,
		DataPathingService:       dataFilesPathingSvc,
		CalendarService:          courseCalAppService,
		GetUnits:                 unitAppService.ByParentID,
		GetLessons:               lessonAppService.ByParentID,
	}
	hugoGenerator, err := hugo.New(hugoParams)
	if err != nil {
		return nil, err
	}

	// site generator service
	siteGenerator := services.NewSiteGeneratorService(hugoGenerator)

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
	lessonAppHandler := handlers.NewLessonHandler(lessonAppService, nodeAppService, e.Reverse, dataFilesPathingSvc, slidesService, fileService, markdownService)
	unitAppHandler := handlers.NewUnitHandler(unitAppService, nodeAppService, fileService, markdownService, e.Reverse)
	courseAppHandler := handlers.NewCourseHandler(siteGenerator, courseAppService, nodeAppService, fileService, markdownService, e.Reverse)
	courseCalAppHandler := handlers.NewCourseCalHandler(courseCalAppService, courseOccasionService, nodeAppService, lessonAppService, unitAppService, e.Reverse)
	termAppHandler := handlers.NewTermHandler(termAppService, nodeAppService, fileService, markdownService, e.Reverse)
	termCalHandler := handlers.NewTermCalHandler(termCalService, nodeAppService, termOccasionService, e.Reverse)
	dashboardAppHandler := handlers.NewDashboardHandler(userAppService, termAppService, nodeAppService, e.Reverse)

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

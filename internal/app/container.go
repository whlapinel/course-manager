package app

import (
	"database/sql"
	"gh_static_portfolio/internal/app/handlers"
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/features/auth"
	"gh_static_portfolio/internal/features/home"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/user"
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
	protected := root.Group("", authentication.AddCookieToHeader, authentication.JWTMiddlewareProtectedNew(e.Reverse(routes.GetSignin.String())), authentication.GetClaims)
	homeService := home.NewService(struct{}{})
	homeHandler := home.NewHandler(*homeService, e)
	err = home.RegisterRoutes(root, homeHandler)
	if err != nil {
		return nil, err
	}
	// repositories
	userRepo := sqlite.NewUserRepo(queries)
	termRepo := sqlite.NewTermRepo(queries)
	// courseRepo := sqlite.NewCourseRepo(queries)
	// unitRepo := sqlite.NewUnitRepo(queries)
	// lessonRepo := sqlite.NewLessonRepo(queries)

	// services
	authService := auth.NewService(struct{}{}, userRepo)
	userService := user.NewService(userRepo)
	termService := term.NewService(termRepo)

	// handlers
	authHandler := auth.NewHandler(authService, e)
	userAppHandler := handlers.NewUserHandler(userService, termService, e)
	userHandler := user.NewHandler(userService, e)
	termHandler := term.NewHandler(termService)

	// register routes
	err = auth.RegisterRoutes(root, authHandler)
	if err != nil {
		return nil, err
	}
	handlers.RegisterUserRoutes(protected, userAppHandler)
	user.RegisterRoutes(protected, userHandler)
	term.RegisterRoutes(protected, termHandler)

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

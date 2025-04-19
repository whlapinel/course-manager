package app

import (
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/features/auth"
	"gh_static_portfolio/internal/features/home"
	"gh_static_portfolio/internal/features/user"
	"gh_static_portfolio/internal/infrastructure/sqlite"
	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
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
	defer db.Close()
	root := e.Group("")
	homeService := home.NewService(struct{}{})
	homeHandler := home.NewHandler(*homeService, e)
	err = home.RegisterRoutes(root, homeHandler)
	if err != nil {
		return nil, err
	}
	authService := auth.NewService(struct{}{})
	authHandler := auth.NewHandler(*authService, e)
	err = auth.RegisterRoutes(root, authHandler)
	if err != nil {
		return nil, err
	}

	termRepo := sqlite.NewTermRepo(*queries)
	userRepo := sqlite.NewUserRepo(*queries)
	svc := user.NewService(userRepo, termRepo)
	userHandler := user.NewHandler(*svc, e)
	user.RegisterRoutes(root, userHandler)
	return &App{
		Echo: e,
	}, nil

}

func (app *App) Start(url string) error {

	for _, route := range app.Echo.Routes() {
		log.Printf("\nMethod: %s\nName: %s\nPath: %s\n", route.Method, route.Name, route.Path)
	}
	return app.Echo.Start(url)

}

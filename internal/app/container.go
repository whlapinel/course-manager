package app

import (
	"gh_static_portfolio/internal/features/user"
	"gh_static_portfolio/internal/infrastructure/sqlite"

	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
}

func New() (*App, error) {
	e := echo.New()
	queries, db, err := sqlite.InitDB("./course_manager.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	root := e.Group("/")
	userRepo := sqlite.NewUserRepo(*queries)
	svc := user.NewService(userRepo)
	userHandler := user.NewHandler(*svc, e)
	user.RegisterRoutes(root, userHandler)
	return &App{
		Echo: e,
	}, nil

}

func (app *App) Start(url string) error {
	return app.Echo.Start(url)

}

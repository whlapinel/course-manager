package main

import (
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/handlers"
	"gh_static_portfolio/internal/service"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(logger)
	queries, db, err := data.InitDB("internal/data/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	courseHandler := handlers.NewCourseHandler(e, service.NewCourseService(courseRepo))
	courseHandler.Mount()
	assets.RegisterStatic(e)
	e.Logger.Fatal(e.Start(":1323"))

}

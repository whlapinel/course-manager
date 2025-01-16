package main

import (
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/service"
	"gh_static_portfolio/cmd/web_app/assets"
	"gh_static_portfolio/cmd/web_app/handlers"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(logger)
	queries, db, err := data.InitDB("course_manager.db")
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

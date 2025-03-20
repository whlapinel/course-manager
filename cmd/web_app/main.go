package main

import (
	"fmt"
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/handlers"
	"gh_static_portfolio/internal/service"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(logger)
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	// moved to starting marp with script rather than in go app
	// startMarp := exec.Command("marp", "-s", "internal/data/users")
	// err = startMarp.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// go func() {
	// 	if err := startMarp.Wait(); err != nil {
	// 		log.Printf("Marp exited with error: %v", err)
	// 	} else {
	// 		log.Printf("Marp exited normally")
	// 	}
	// }()
	// defer func() {
	// 	startMarp.Process.Signal(os.Interrupt)
	// }()
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	service := service.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(e, service)
	handlers.Mount(service, e) // new way
	// Log all NEW registered routes
	for _, route := range e.Routes() {
		fmt.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
	courseHandler.Mount() // old way
	assets.RegisterStatic(e)
	host := os.Getenv("ECHO_HOST")
	log.Println("Host:", host)
	port := os.Getenv("ECHO_PORT")
	log.Println("Port:", port)
	startString := fmt.Sprintf("%s:%s", host, port)
	e.Logger.Fatal(e.Start(startString))
	// if environment == development {
	// 	e.Logger.Fatal(e.Start(startString))
	// } else if environment == production {
	// 	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	// 	e.Logger.Fatal(e.StartAutoTLS(startString))
	// }

}

package main

import (
	"fmt"
	"gh_static_portfolio/internal/assets"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/handlers"
	"gh_static_portfolio/internal/service"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const production = "production"
const development = "development"

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(logger)
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	startMarp := exec.Command("marp", "-s", "internal/data")
	err = startMarp.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		startMarp.Process.Signal(os.Interrupt)
	}()
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
	_, err = LoadEnvironment()
	if err != nil {
		log.Fatal("error loading environment")
	}
	host := os.Getenv("HOST")
	log.Println("Host:", host)
	port := os.Getenv("PORT")
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

func LoadEnvironment() (string, error) {
	if os.Getenv("ENV") == development {
		godotenv.Load(".env.development")
		return development, nil
	}
	if os.Getenv("ENV") == production {
		godotenv.Load(".env.production")
		return production, nil
	}
	return "", fmt.Errorf("environment not expected:'%s'", os.Getenv("ENV"))
}

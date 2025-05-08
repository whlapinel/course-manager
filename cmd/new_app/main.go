package main

import (
	"fmt"
	"gh_static_portfolio/internal/app"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// feature-level services
	marpHost := os.Getenv("MARP_HOST")
	marpPort := os.Getenv("MARP_PORT")
	baseURL := fmt.Sprintf("http://%s:%s", marpHost, marpPort)
	app, err := app.New(app.NewAppParams{MarpBaseURL: baseURL})
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("ECHO_HOST")
	if host == "" {
		host = "localhost"
	}
	log.Println("Host:", host)
	port := os.Getenv("ECHO_PORT")
	if port == "" {
		port = "1323"
	}
	log.Println("Port:", port)
	startString := fmt.Sprintf("%s:%s", host, port)
	log.Fatal(app.Start(startString))
}

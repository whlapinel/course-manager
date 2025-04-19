package main

import (
	"fmt"
	"gh_static_portfolio/internal/app"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app, err := app.New()
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

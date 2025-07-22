package main

import (
	"fmt"
	"gh_static_portfolio/internal/app"
	"log"
	"os"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// feature-level services
	marpHost := os.Getenv("MARP_HOST")
	marpPort := os.Getenv("MARP_PORT")
	baseURL := fmt.Sprintf("http://%s:%s", marpHost, marpPort)
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
	production := os.Getenv("PRODUCTION") == "true"
	var noAuth bool
	domain := "localhost"
	if production {
		domain = os.Getenv("PROD_DOMAIN")
	} else {
		noAuth = os.Getenv("NO_AUTH") == "true"
	}
	log.Println("Domain: ", domain)
	app, err := app.New(app.NewAppParams{
		Domain:        domain,
		MarpBaseURL:   baseURL,
		DevModeNoAuth: noAuth,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(app.Start(startString))
}

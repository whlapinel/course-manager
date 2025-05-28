package main

import (
	"context"
	"database/sql"
	"flag"
	_ "gh_static_portfolio/internal/infrastructure/sqlite/migrations/schema"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "internal/infrastructure/sqlite/migrations/schema", "directory with migration files")
)

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("goose: failed to parse flags: %v", err)
	}

	args := flags.Args()
	if len(args) < 2 {
		log.Fatalf("usage: go run main.go [database] [command] [args...]")
	}

	dbstring := args[0]
	command := args[1]
	commandArgs := args[2:]

	db, err := sql.Open("sqlite3", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v", err)
	}
	log.Println("migration script: directory is ", *dir)
	defer db.Close()

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, *dir, commandArgs...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

package sqlite

import (
	"context"
	"database/sql"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"log"
)

func InitDB(fileName string) (*database.Queries, *sql.DB, error) {
	var queries *database.Queries
	ctx := context.Background()
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, nil, err
	}
	// Enable foreign keys
	_, err = db.ExecContext(ctx, "PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}

	// Check if foreign keys are enabled
	var foreignKeysEnabled int
	err = db.QueryRow("PRAGMA foreign_keys;").Scan(&foreignKeysEnabled)
	if err != nil {
		log.Fatal("Failed to check foreign_keys status:", err)
	}

	log.Println("Foreign keys enabled:", foreignKeysEnabled)
	queries = database.New(db)
	return queries, db, nil

}

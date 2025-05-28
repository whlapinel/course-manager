package migrations

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed 20250527180816/down.sql
var down string

//go:embed 20250527180816/up.sql
var up string

func init() {
	goose.AddMigrationContext(upOccasions, downOccasions)
}

func upOccasions(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	result, err := tx.ExecContext(ctx, up)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func downOccasions(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	result, err := tx.ExecContext(ctx, down)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil

}

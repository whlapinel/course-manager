package migrations

import (
	"context"
	"database/sql"

	_ "embed"

	"github.com/pressly/goose/v3"
)

//go:embed 20250528151924/up.sql
var up20250528 string

func init() {
	goose.AddMigrationContext(upDropassessments, downDropassessments)
}

func upDropassessments(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, up20250528)
	if err != nil {
		return err
	}
	return nil
}

func downDropassessments(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

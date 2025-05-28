package migrations

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/pressly/goose/v3"
)

//go:embed 20250527184953/up.sql
var courseOccasionUp string

//go:embed 20250527184953/down.sql
var courseOccasionDown string

func init() {
	goose.AddMigrationContext(upCourseOccasion, downCourseOccasion)
}

func upCourseOccasion(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, courseOccasionUp)
	if err != nil {
		return err
	}
	return nil
}

func downCourseOccasion(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, courseOccasionDown)
	if err != nil {
		return err
	}
	return nil

}

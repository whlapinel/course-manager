package migrations

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/pressly/goose/v3"
)

//go:embed 20250405170401/first.sql
var first string

//go:embed 20250405170401/second.sql
var second string

//go:embed 20250405170401/third.sql
var third string

//go:embed 20250405170401/fourth.sql
var fourth string

//go:embed 20250405170401/fifth.sql
var fifth string

var commands = []string{first, second, third, fourth, fifth}

func init() {
	goose.AddMigrationContext(upAssessmentParent, downAssessmentParent)
}

func upAssessmentParent(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	for _, cmd := range commands {
		_, err := tx.ExecContext(ctx, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func downAssessmentParent(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

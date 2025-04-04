package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upModCategory, downModCategory)
}

type CategoryMap map[int]string

var catMap = CategoryMap{
	0: "perform",
	1: "rehearse",
	2: "prepare",
	3: "midterm",
	4: "final",
}

func upModCategory(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	// Add new text column
	_, err := tx.ExecContext(ctx, `ALTER TABLE assessments ADD COLUMN category_text TEXT`)
	if err != nil {
		return err
	}

	// Populate new column with mapped text
	for k, v := range catMap {
		_, err := tx.ExecContext(ctx, `UPDATE assessments SET category_text = ? WHERE category = ?`, v, k)
		if err != nil {
			return err
		}
	}

	// Drop old column and rename new column
	_, err = tx.ExecContext(ctx, `ALTER TABLE assessments RENAME TO assessments_old`)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
			CREATE TABLE assessments (
			id INTEGER PRIMARY KEY,
			lesson_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			instructions TEXT NOT NULL,
			file TEXT,
			category TEXT NOT NULL,
			date_assigned TEXT NOT NULL,
			date_due TEXT NOT NULL,
			dropped INTEGER NOT NULL,
			FOREIGN KEY (lesson_id) REFERENCES lessons (id)
			)
		`)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
			INSERT INTO assessments (id, lesson_id, name, instructions, file, category, date_assigned, date_due, dropped)
			SELECT id, lesson_id, name, instructions, file, category_text, date_assigned, date_due, dropped FROM assessments_old
		`)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DROP TABLE assessments_old`)
	return err

}

func downModCategory(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

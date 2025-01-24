package data

import (
	"context"
	"database/sql"
	"gh_static_portfolio/internal/data/database"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	_ "embed"
)

type CourseRepo struct {
	queries *database.Queries
}

func NewCourseRepo(db *database.Queries) CourseRepo {
	return CourseRepo{queries: db}
}

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

// for copying all files in the lesson, unit, or course directory
func CopyNodeDir(srcRoot, destRoot string) error {
	log.Println("copying", srcRoot, "to", destRoot)
	// if directory doesn't exist, early return
	_, err := os.Stat(srcRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	files, err := os.ReadDir(srcRoot)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}
	return filepath.WalkDir(srcRoot, func(srcPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destRoot, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}
		return copyFile(srcPath, destPath)
	})
}

func copyFile(srcPath, destPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

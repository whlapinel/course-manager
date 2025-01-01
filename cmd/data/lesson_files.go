package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func (c CourseRepo) SaveFile(file domain.File) (int, error) {
	dbFile, err := c.queries.SaveFile(context.Background(), database.SaveFileParams{
		Name: file.Name,
		Description: sql.NullString{
			Valid:  file.Description != "",
			String: file.Description,
		},
		FileName: filepath.Base(file.SourcePath),
		Modified: time.Now().Format(time.DateOnly),
	})
	if err != nil {
		return 0, err
	}
	file.ID = int(dbFile.ID)
	err = AddFile(file)
	if err != nil {
		return 0, err
	}
	return int(dbFile.ID), nil
}

func FileName(baseDir string, file domain.File) (string, error) {
	basePath, err := filepath.Abs(baseDir)
	if err != nil {
		return "", err
	}
	dstPath := filepath.Join(basePath, fmt.Sprintf("file_id_%d_%s", file.ID, filepath.Base(file.SourcePath)))
	return dstPath, nil
}

func AddFile(file domain.File) error {
	srcFile, err := os.Open(file.SourcePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	err = os.MkdirAll("./cmd/data/files", 0777)
	if err != nil {
		return err
	}
	dstPath, err := FileName("./cmd/data/files", file)
	if err != nil {
		return err
	}
	newFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	bytes, err := io.Copy(newFile, srcFile)
	if err != nil {
		return err
	}
	closeErr := newFile.Close()
	if closeErr != nil {
		return closeErr
	}
	log.Println("Bytes written: ", bytes)
	return nil

}

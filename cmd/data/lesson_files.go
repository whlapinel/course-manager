package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func (c CourseRepo) getLessonFiles(lesson domain.Lesson) (domain.FilesDir, error) {
	dbFile, err := c.queries.GetLessonFilesDir(context.Background(), int64(lesson.ID))
	if err != nil {
		return domain.FilesDir{}, err
	}
	file := domain.FilesDir{
		ID:          int(dbFile.ID),
		Name:        dbFile.Name,
		Description: dbFile.Description.String,
	}
	return file, nil
}

func LessonFilesDirPath(file domain.FilesDir) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/files"
	return filepath.Join(basePath, fmt.Sprintf("files_id_%d", file.ID))
}

// temp function to save files to app file system, V2
func (cr CourseRepo) SaveFilesDir(filesDir string, lesson domain.Lesson) error {
	newFilesDir := domain.NewFile(lesson.Name, "", filesDir)
	dbFilesDir, err := cr.queries.SaveFilesDir(context.Background(), database.SaveFilesDirParams{
		Name: newFilesDir.Name,
		Description: sql.NullString{
			Valid:  newFilesDir.Description != "",
			String: newFilesDir.Description,
		},
	})
	if err != nil {
		return err
	}
	newFilesDir.ID = int(dbFilesDir.ID)
	destRoot := LessonFilesDirPath(newFilesDir)
	_, err = cr.queries.SaveLessonFilesDir(context.Background(), database.SaveLessonFilesDirParams{
		FileID:   int64(newFilesDir.ID),
		LessonID: int64(lesson.ID),
	})
	if err != nil {
		return err
	}
	return filepath.WalkDir(filesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(filesDir, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destRoot, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}
		return copyFile(path, destPath)
	})
}

func copyFile(srcPath, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	return err
}


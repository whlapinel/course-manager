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
func (cr CourseRepo) SaveFilesDir(filesDir domain.FilesDir, lesson domain.Lesson) (domain.FilesDir, error) {
	dbFilesDir, err := cr.queries.SaveFilesDir(context.Background(), database.SaveFilesDirParams{
		Name: filesDir.Name,
		Description: sql.NullString{
			Valid:  filesDir.Description != "",
			String: filesDir.Description,
		},
	})
	if err != nil {
		return domain.FilesDir{}, err
	}
	filesDir.ID = int(dbFilesDir.ID)
	destRoot := LessonFilesDirPath(filesDir)
	_, err = cr.queries.SaveLessonFilesDir(context.Background(), database.SaveLessonFilesDirParams{
		FileID:   int64(filesDir.ID),
		LessonID: int64(lesson.ID),
	})
	if err != nil {
		return domain.FilesDir{}, err
	}
	return filesDir, filepath.WalkDir(filesDir.SourcePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(filesDir.SourcePath, path)
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

func (cr CourseRepo) NewFileDir(lesson domain.Lesson) (domain.FilesDir, error) {
	fileDir := domain.NewFileDir(lesson.Name, lesson.Description, "")
	dbFilesDir, err := cr.queries.SaveFilesDir(context.Background(), database.SaveFilesDirParams{
		Name: fileDir.Name,
		Description: sql.NullString{
			Valid:  fileDir.Description != "",
			String: fileDir.Description,
		},
	})
	if err != nil {
		return domain.FilesDir{}, err
	}
	fileDir.ID = int(dbFilesDir.ID)
	destRoot := LessonFilesDirPath(fileDir)
	err = os.MkdirAll(destRoot, os.ModePerm)
	if err != nil {
		return domain.FilesDir{}, err
	}
	_, err = cr.queries.SaveLessonFilesDir(context.Background(), database.SaveLessonFilesDirParams{
		FileID:   int64(fileDir.ID),
		LessonID: int64(lesson.ID),
	})
	if err != nil {
		return domain.FilesDir{}, err
	}
	return fileDir, nil
}

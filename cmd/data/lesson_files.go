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
)

func (c CourseRepo) SaveFile(file domain.File) (int, error) {
	dbFile, err := c.queries.SaveFile(context.Background(), database.SaveFileParams{
		Name: file.Name,
		Description: sql.NullString{
			Valid:  file.Description != "",
			String: file.Description,
		},
		BasePath: file.BasePath,
	})
	if err != nil {
		return 0, err
	}
	file.ID = int(dbFile.ID)
	err = addFile(file)
	if err != nil {
		return 0, err
	}
	return int(dbFile.ID), nil
}

func (c CourseRepo) GetLessonFiles(lesson domain.Lesson) ([]domain.File, error) {
	dbLessonFiles, err := c.queries.GetLessonFiles(context.Background(), int64(lesson.ID))
	if err != nil {
		return nil, err
	}
	var files []domain.File
	for _, dbFile := range dbLessonFiles {
		file := domain.File{
			Name:        dbFile.Name,
			Description: dbFile.Description.String,
			BasePath:    dbFile.BasePath,
		}
		files = append(files, file)
	}
	return files, nil
}

func LessonFilePath(file domain.File) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/files"
	dstPath := filepath.Join(basePath, fmt.Sprintf("file_id_%d_%s", file.ID, file.BasePath))
	return dstPath
}

func addFile(file domain.File) error {
	srcFile, err := os.Open(file.SourcePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstPath := LessonFilePath(file)
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

// if file ID is 0, the file will be saved before adding to the lesson
func (cr CourseRepo) AddFileToLesson(file domain.File, lesson domain.Lesson) error {
	if file.ID == 0 {
		// cannot save without source path
		if file.SourcePath == "" {
			return fmt.Errorf("no source path provided")
		}
		id, err := cr.SaveFile(file)
		if err != nil {
			return err
		}
		file.ID = id
	}
	_, err := cr.queries.SaveLessonFile(context.Background(), database.SaveLessonFileParams{
		LessonID: int64(lesson.ID),
		FileID:   int64(file.ID),
	})
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(file domain.File) error {
	path := LessonFilePath(file)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func SlidesHTMLFilePath(lesson domain.Lesson) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/slides/html"
	return filepath.Join(basePath, fmt.Sprintf("lesson_%d_slides.html", lesson.ID))
}

func SlidesMarkdownFilePath(lesson domain.Lesson) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/slides/markdown"
	return filepath.Join(basePath, fmt.Sprintf("lesson_%d_slides.md", lesson.ID))
}

package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"io"
	"os"
	"path/filepath"
)

func (cr CourseRepo) SaveSlides(slides domain.Slides, lesson domain.Lesson) (int, error) {
	// save in DB
	dbSlides, err := cr.queries.SaveSlides(context.Background(), database.SaveSlidesParams{
		Name: slides.Name,
		Description: sql.NullString{
			Valid:  slides.Description != "",
			String: slides.Description,
		},
	})
	if err != nil {
		return 0, err
	}
	slides.ID = int(dbSlides.ID)
	srcPath := slides.SourcePath
	destPath := OldSlidesMarkdownFilePath(slides)
	src, err := os.Open(srcPath)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	dst, err := os.Create(destPath)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return 0, err
	}
	_, err = cr.queries.SaveLessonSlides(context.Background(), database.SaveLessonSlidesParams{
		SlidesID: int64(slides.ID),
		LessonID: int64(lesson.ID),
	})
	if err != nil {
		return 0, err
	}
	return slides.ID, nil
}

func OldSlidesHTMLFilePath(slides domain.Slides) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/slides/html"
	return filepath.Join(basePath, fmt.Sprintf("slides_%d.html", slides.ID))
}

func OldSlidesMarkdownFilePath(slides domain.Slides) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/slides/markdown"
	return filepath.Join(basePath, fmt.Sprintf("slides_%d.md", slides.ID))
}

func NewSlidesMarkdownFilePath(lesson domain.Lesson) string {
	return filepath.Join(LessonDirPath(lesson.ID), "slides.md")
}

func NewSlidesHTMLFilePath(lesson domain.Lesson) string {
	return filepath.Join(LessonDirPath(lesson.ID), "slides.html")
}

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

func (c CourseRepo) SaveImage(image domain.Image) (int, error) {
	dbImage, err := c.queries.SaveImage(context.Background(), database.SaveImageParams{
		Name: image.Name,
		Description: sql.NullString{
			Valid:  image.Description != "",
			String: image.Description,
		},
		BasePath: image.BasePath,
	})
	if err != nil {
		return 0, err
	}
	image.ID = int(dbImage.ID)
	err = addImage(image)
	if err != nil {
		return 0, err
	}
	return int(dbImage.ID), nil
}

func (c CourseRepo) GetLessonImage(lesson domain.Lesson) (domain.Image, error) {
	dbImage, err := c.queries.GetLessonImage(context.Background(), int64(lesson.ID))
	if err != nil {
		return domain.Image{}, err
	}
	image := domain.Image{
		ID:          int(dbImage.ID),
		Name:        dbImage.Name,
		Description: dbImage.Description.String,
		BasePath:    dbImage.BasePath,
	}
	return image, nil
}

// This returns the path for an image, either for retrieving or storing the actual file
func ImagesPath(image domain.Image) string {
	basePath := "/home/whlapinel/personal_projects/course_manager/cmd/data/images"
	dstPath := filepath.Join(basePath, fmt.Sprintf("image_id_%d_%s", image.ID, image.BasePath))
	return dstPath
}

// this copies the image to the app's internal file system
func addImage(image domain.Image) error {
	srcFile, err := os.Open(image.SourcePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstPath := ImagesPath(image)
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

// if image ID is 0, the image will be saved before adding to the lesson
func (cr CourseRepo) AddImageToLesson(image domain.Image, lesson domain.Lesson) error {
	if image.ID == 0 {
		return fmt.Errorf("image has not been saved; id is 0")
	}
	_, err := cr.queries.SaveLessonImage(context.Background(), database.SaveLessonImageParams{
		LessonID: int64(lesson.ID),
		ImageID:  int64(image.ID),
	})
	if err != nil {
		return err
	}
	return nil
}

// if image ID is 0, the image will be saved before adding to the lesson
func (cr CourseRepo) AddImageToCourse(image domain.Image, course domain.Course) error {
	if image.ID == 0 {
		return fmt.Errorf("image has not been saved; id is 0")
	}
	_, err := cr.queries.SaveCourseImage(context.Background(), database.SaveCourseImageParams{
		CourseID: int64(course.ID),
		ImageID:  int64(image.ID),
	})
	if err != nil {
		return err
	}
	return nil
}

// if image ID is 0, the image will be saved before adding to the lesson
func (cr CourseRepo) AddImageToUnit(image domain.Image, unit domain.Unit) error {
	if image.ID == 0 {
		return fmt.Errorf("image has not been saved; id is 0")
	}
	_, err := cr.queries.SaveUnitImage(context.Background(), database.SaveUnitImageParams{
		UnitID:  int64(unit.ID),
		ImageID: int64(image.ID),
	})
	if err != nil {
		return err
	}
	return nil
}

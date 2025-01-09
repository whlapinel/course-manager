package service

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"time"
)

func (svc CourseService) GetLesson(lessonID int) (*domain.Lesson, error) {
	lesson, err := svc.repo.GetLesson(lessonID)
	if err != nil {
		return nil, fmt.Errorf("CourseService.GetLesson: %d", lessonID)
	}
	slides, err := svc.repo.GetSlides(lessonID)
	if err != nil {
		return nil, err
	}
	lesson.Slides = slides
	lessonDates, err := svc.repo.GetLessonDates(lessonID)
	if err != nil {
		return nil, err
	}
	lesson.Dates = lessonDates
	lessonFiles, err := svc.repo.GetLessonFiles(*lesson)
	if err != nil {
		return nil, err
	}
	lesson.Files = lessonFiles
	return lesson, nil

}

func (svc CourseService) GetLessons(unitID int) ([]*domain.Lesson, error) {
	lessons, err := svc.repo.GetLessons(unitID)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

func (svc CourseService) UpdateLesson(l domain.Lesson) error {
	log.Println("CourseService.UpdateLesson: ")
	log.Println("Descr:", l.Description)
	err := svc.repo.UpdateLesson(l)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) Shift(lesson domain.Lesson, term domain.Term, direction domain.CalendarDirection) (domain.Lesson, time.Time, error) {
	if lesson.ID == 0 {
		log.Println("lesson id: ", lesson.ID)
	}
	dates, err := svc.repo.GetLessonDates(lesson.ID)
	if err != nil {
		return domain.Lesson{}, time.Time{}, err
	}
	if len(dates) == 0 {
		log.Println("service: length dates 0!")
	}
	for _, date := range dates {
		log.Println("service: date: ", date.Format(time.DateOnly))
	}
	lesson.Dates = dates
	if len(lesson.Dates) == 0 {
		log.Println("Lesson dates length 0!")
	}
	for _, date := range lesson.Dates {
		log.Println("Service: Before shifting:")
		log.Println(date.Format(time.DateOnly))
	}
	shiftedLesson, newTime, err := lesson.Shift(direction, term)
	if err != nil {
		return domain.Lesson{}, time.Time{}, err
	}
	err = svc.UpdateLesson(shiftedLesson)
	if err != nil {
		return domain.Lesson{}, time.Time{}, err
	}

	for _, date := range shiftedLesson.Dates {
		log.Println("Service: After shifting:")

		log.Println(date.Format(time.DateOnly))
	}
	return shiftedLesson, newTime, nil
}

func (svc CourseService) Extend(lesson domain.Lesson, term domain.Term, direction domain.CalendarDirection) (domain.Lesson, error) {
	shifted, err := lesson.Extend(direction, term)
	if err != nil {
		return domain.Lesson{}, err
	}
	err = svc.UpdateLesson(lesson)
	if err != nil {
		return domain.Lesson{}, err
	}
	return shifted, nil
}

func (svc CourseService) CreateNewLessonSlides(lesson *domain.Lesson) (domain.Slides, error) {
	tempPath := "./temp_files"
	err := os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		return domain.Slides{}, err
	}
	tempFile, err := os.CreateTemp(tempPath, "slides*.md")
	if err != nil {
		return domain.Slides{}, err
	}
	defer os.Remove(tempFile.Name())
	slides := domain.NewSlides(lesson.Name, lesson.Description, tempFile.Name())
	id, err := svc.repo.SaveSlides(slides, *lesson)
	if err != nil {
		return domain.Slides{}, err
	}
	slides.ID = id
	lesson.Slides = slides
	return slides, nil
}

// lesson param is a pointer because this updates the Lesson.Files field
func (svc CourseService) CreateNewLessonFileDir(lesson *domain.Lesson) (domain.FilesDir, error) {
	fileDir, err := svc.repo.NewFileDir(*lesson)
	if err != nil {
		return domain.FilesDir{}, err
	}
	lesson.Files = fileDir
	return fileDir, nil
}

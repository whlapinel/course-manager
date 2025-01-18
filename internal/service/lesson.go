package service

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"log"
	"os"
	"time"
)

func (svc CourseService) CreateNewLessonSlides(lessonID int) error {
	path := data.NewSlidesMarkdownFilePath(lessonID)
	// make sure file does not exist
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("CourseService.CreateNewLessonSlides: file exists already at %s", path)
	}
	_, err = os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) CreateNewLessonFileDir(lessonID int) error {
	path := data.NewLessonFilesDirPath(lessonID)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) GetLesson(lessonID int) (*domain.Lesson, error) {
	lesson, err := svc.repo.GetLesson(lessonID)
	if err != nil {
		return nil, fmt.Errorf("CourseService.GetLesson: %d", lessonID)
	}
	lessonDates, err := svc.repo.GetLessonDates(lessonID)
	if err != nil {
		return nil, err
	}
	lesson.Dates = lessonDates
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

// This version is for the web app
func (svc CourseService) WebShift(termID, courseID, lessonID int, cd domain.CalendarDirection) error {
	lesson, err := svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	term, err := svc.repo.GetTermByID(termID)
	if err != nil {
		return err
	}
	termWithDates, err := svc.repo.GetTermDates(termID)
	if err != nil {
		return err
	}
	term.InstructionalDays = termWithDates.InstructionalDays
	newLesson, _, err := lesson.Shift(cd, term)
	if err != nil {
		return err
	}
	err = svc.UpdateLesson(newLesson)
	if err != nil {
		return err
	}
	return err
}

// This version is for the fyne app
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

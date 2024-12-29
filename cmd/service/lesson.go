package service

import (
	"gh_static_portfolio/cmd/domain"
	"log"
	"time"
)

func (svc CourseService) UpdateLesson(l domain.Lesson) error {
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

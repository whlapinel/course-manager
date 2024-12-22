package data

import (
	"log"
	"testing"
	"time"
)

func TestCancelAndRescheduleLesson(t *testing.T) {
	lessons, err := cr.GetLessons(13)
	if err != nil {
		t.Error()
	}
	for _, lesson := range lessons {
		log.Println("lesson:", lesson.Name)
	}
	lesson := lessons[0]
	log.Println("dates before canceling:", lesson.Dates)
	date := lesson.Dates[0]
	cr.DeleteLessonDate(lesson, date)
	lessons, err = cr.GetLessons(13)
	if err != nil {
		t.Error(err)
	}
	lesson = lessons[0]
	log.Println("dates after canceling:", lesson.Dates)
	lessons, err = cr.GetLessons(13)
	if err != nil {
		t.Error(err)
	}
	lesson = lessons[0]
	log.Println("lesson:", lesson.Name)
	log.Println("dates before scheduling:", lesson.Dates)
	date = time.Date(2024, 8, 26, 0, 0, 0, 0, time.Local)
	cr.AddLessonDate(lesson, date)
	lessons, err = cr.GetLessons(13)
	if err != nil {
		t.Error(err)
	}
	lesson = lessons[0]
	log.Println("dates after scheduling:", lesson.Dates)

}

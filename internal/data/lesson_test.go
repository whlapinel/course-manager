package data

import (
	"log"
	"testing"
)

func TestGetLesson(t *testing.T) {
	lesson, err := cr.GetLesson(1)
	if err != nil {
		t.Error(err)
	}
	log.Println(lesson.Name)
}

func TestGetLessonDates(t *testing.T) {
	dates, err := cr.GetLessonDates(1)
	if err != nil {
		t.Error(err)
	}
	for _, date := range dates {
		log.Println(date)
	}

}

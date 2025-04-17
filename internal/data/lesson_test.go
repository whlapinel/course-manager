package data

import (
	"context"
	"gh_static_portfolio/internal/data/database"
	"log"
	"testing"
	"time"
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

func TestGetLessonsOnDate(t *testing.T) {
	lessons, err := cr.queries.GetLessonsOnDate(context.Background(), database.GetLessonsOnDateParams{
		Date:   "2025-01-21",
		TermID: 1,
	})
	if err != nil {
		t.Error(err)
	}
	for _, lesson := range lessons {
		log.Println(lesson)
	}
}

func TestGetLessonsOnDateForCourse(t *testing.T) {
	date := time.Date(2025, time.April, 29, 0, 0, 0, 0, time.Local)
	courseID := 5
	lessons, err := cr.GetLessonsOnDateForCourse(date, courseID)
	if err != nil {
		t.Error(err)
	}
	for _, lesson := range lessons {
		log.Println(lesson)
	}
}

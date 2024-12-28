package service

import (
	"gh_static_portfolio/cmd/domain"
	"log"
	"testing"
	"time"
)

var testTermStart = time.Date(2025, time.January, 2, 0, 0, 0, 0, time.Local)
var testTermEnd = testTermStart.AddDate(0, 3, 0)
var testLessonDates = []time.Time{time.Date(2025, time.January, 3, 0, 0, 0, 0, time.Local)}
var testNonInstructDays = []time.Time{time.Date(2025, time.January, 6, 0, 0, 0, 0, time.Local)}

func TestShift(t *testing.T) {
	svc := NewCourseService(cr)
	terms, err := svc.GetTerms()
	if err != nil {
		t.Error(err)
	}
	courses, err := svc.GetCourses(terms[0].ID)
	if err != nil {
		t.Error(err)
	}
	lesson := courses[0].Units[0].Lessons[0]
	log.Println("before shifting: ", lesson.Name)
	lesson.Dates = []time.Time{}
	for _, date := range lesson.Dates {
		log.Println(date.Format(time.DateOnly))
	}
	shifted, newTime, err := svc.Shift(lesson, terms[0], domain.Left)
	if err != nil {
		t.Error(err)
	}
	log.Println("after shifting: ", newTime.Format(time.DateOnly))
	got := shifted.Dates
	log.Println("shifted dates", got)

}

func TestExtend(t *testing.T) {
	svc := NewCourseService(cr)
	terms, err := svc.GetTerms()
	if err != nil {
		t.Error(err)
	}
	courses, err := svc.GetCourses(terms[0].ID)
	if err != nil {
		t.Error(err)
	}
	lesson := courses[0].Units[0].Lessons[0]
	extended, err := svc.Extend(lesson, terms[0], domain.Right)
	if err != nil {
		t.Error(err)
	}
	got := extended.Dates
	log.Println("extended dates")
	for _, date := range got {
		log.Println(date.Format(time.DateOnly))
	}

}

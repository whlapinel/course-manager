package domain

import (
	"log"
	"testing"
	"time"
)

var testTermStart = time.Date(2025, time.January, 2, 0, 0, 0, 0, time.Local)
var testTermEnd = testTermStart.AddDate(0, 3, 0)
var testLessonDates = []time.Time{time.Date(2025, time.January, 3, 0, 0, 0, 0, time.Local)}
var testNonInstructDays = []time.Time{time.Date(2025, time.January, 6, 0, 0, 0, 0, time.Local)}

func TestShift(t *testing.T) {
	lesson := NewLesson(1, 1, 1, "Lesson 0.1", "How to turn on your computer", testLessonDates)
	term, err := NewTerm(testTermStart, testTermEnd, testNonInstructDays, Semester, 1, "Spring 2025")
	if err != nil {
		t.Error()
	}
	shiftedLesson, newTime, err := lesson.Shift(Right, term)
	if err != nil {
		t.Error()
	}
	log.Println(newTime.Format(time.DateOnly))
	got := shiftedLesson.Dates[0]
	expected := testLessonDates[0].AddDate(0, 0, 4)
	if got != expected {
		t.Error("expected: ", expected, "got: ", got)
	}

}

func TestExtend(t *testing.T) {
	lesson := NewLesson(1, 1, 1, "Lesson 0.1", "How to turn on your computer", testLessonDates)
	term, err := NewTerm(testTermStart, testTermEnd, testNonInstructDays, Semester, 1, "Spring 2025")
	if err != nil {
		t.Error()
	}
	extended, err := lesson.Extend(Right, term)
	if err != nil {
		t.Error()
	}
	if !(len(extended.Dates) > 1) {
		t.Error("length of lesson.Dates 1 or less after extending")
	}
	got := extended.Dates[0]
	expected := testLessonDates[0]
	if got != expected {
		t.Error("expected: ", expected, "got: ", got)
	}
	got = extended.Dates[1]
	expected = testLessonDates[0].AddDate(0, 0, 4)
	if got != expected {
		t.Error("expected: ", expected, "got: ", got)
	}

}

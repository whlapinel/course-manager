package domain

import (
	"time"
)

type Assessment struct {
	ID           int
	LessonID     int
	Name         string
	Instructions string
	File         string
	Category     AssessmentCategory
	DateAssigned time.Time
	DateDue      time.Time
	Dropped      bool
}

type AssessmentCategory string

const (
	Prepare  AssessmentCategory = "prepare"
	Rehearse AssessmentCategory = "rehearse"
	Perform  AssessmentCategory = "perform"
	Midterm  AssessmentCategory = "midterm"
	Final    AssessmentCategory = "final"
)

var Categories = []AssessmentCategory{
	Prepare,
	Rehearse,
	Perform,
	Midterm,
	Final,
}

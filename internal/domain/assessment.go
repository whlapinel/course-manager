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
	Rehearse                    = "rehearse"
	Perform                     = "perform"
	Midterm                     = "midterm"
	Final                       = "final"
)

var Categories = []AssessmentCategory{
	Prepare,
	Rehearse,
	Perform,
	Midterm,
	Final,
}

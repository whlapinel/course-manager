package domain

import (
	"fmt"
	"log"
	"time"
)

type Assessment struct {
	ID           int
	LessonID     int
	Name         string
	Instructions string
	Category     AssessmentCategory
	DateAssigned time.Time
	DateDue      time.Time
	Dropped      bool
}

type AssessmentCategory int

const (
	Perform AssessmentCategory = iota
	Rehearse
	Prepare
	Midterm
	Final
)

var Categories = []AssessmentCategory{
	Perform,
	Rehearse,
	Prepare,
	Midterm,
	Final,
}

var catStrings = []string{"perform", "rehearse", "prepare", "midterm", "final"}

func (cat AssessmentCategory) String() string {
	log.Println("category:", int(cat))
	return catStrings[cat]
}

func ParseCategories(cat string) (AssessmentCategory, error) {
	for i, category := range Categories {
		if cat == category.String() {
			return Categories[i], nil
		}
	}
	return -1, fmt.Errorf("invalid category")
}

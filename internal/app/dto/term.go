package dto

import (
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/ports"
	"time"
)

type Term struct {
	term.Term `json:"term"`
	Courses   []Course            `json:"courses"`
	Occasions []occasion.Occasion `json:"occasions"`
}

func (t Term) GetChildren() []ports.Node {
	var courses []ports.Node
	for _, c := range t.Courses {
		courses = append(courses, c)
	}
	return courses
}

func (t Term) GetParentTypeName() string {
	return UserTypeName.String()
}

func (t Term) GetTypeName() string {
	return TermTypeName.String()
}

func (t Term) GetChildTypeName() string {
	return CourseTypeName.String()
}

func (t Term) GetNumber() int {
	return -1
}

type NonInstructionalDays struct {
	TermID []int
	Dates  []time.Time
}

const (
	Semester = "semester"
	YearLong = "year_long"
)

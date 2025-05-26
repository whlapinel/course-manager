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

func (t Term) GetName() string {
	return t.Name
}

func (t Term) GetNumber() int {
	return -1
}

func (t Term) GetDescription() string {
	return t.Description
}

func (t Term) GetID() any {
	return t.ID
}

func (t Term) GetParentID() any {
	return t.UserID
}

func (t Term) Children() []ports.Node {
	var courses []ports.Node
	for _, c := range t.Courses {
		courses = append(courses, c)
	}
	return courses
}

func (t Term) TypeName() string {
	return TermTypeName.String()
}

func (t Term) ParentTypeName() string {
	return UserTypeName.String()
}

func (t Term) ChildTypeName() string {
	return CourseTypeName.String()
}

type NonInstructionalDays struct {
	TermID []int
	Dates  []time.Time
}

const (
	Semester = "semester"
	YearLong = "year_long"
)

func (t Term) Designation() string {
	return ""
}

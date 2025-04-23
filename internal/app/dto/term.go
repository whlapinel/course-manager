package dto

import (
	"fmt"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/core/term"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"slices"
	"time"
)

type Term struct {
	term.Term
	Courses   []Course
	Occasions []occasion.Occasion
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

func (t Term) GetParentID() int {
	return t.UserID
}

func (t Term) Children() []templates.Node {
	var courses []templates.Node
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

func NonInstructionDayRange(start time.Time, end time.Time) []time.Time {
	dateRange := []time.Time{}
	currDate := start
	for !currDate.After(end) {
		dateRange = append(dateRange, currDate)
		currDate = currDate.Add(24 * time.Hour)
	}
	return dateRange
}

func InstructionDays(term Term) []time.Time {
	currDate := term.Start
	instructionDays := []time.Time{}
	for !currDate.After(term.End) {
		isInstructionDay := true
		if currDate.Weekday() == 0 || currDate.Weekday() == 6 {
			isInstructionDay = false
		}
		for _, day := range term.NonInstructionalDays {
			if IsSameDate(currDate, day) {
				isInstructionDay = false
			}
		}
		if isInstructionDay {
			instructionDays = append(instructionDays, currDate)
		}
		currDate = currDate.Add(24 * time.Hour)
	}
	return instructionDays
}

func IsSameDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		return true
	}
	return false
}

func (t Term) IsInstructionDay(queryDate time.Time) bool {
	if queryDate.Before(t.Start) || queryDate.After(t.End) {
		return false
	}
	for _, termDate := range t.InstructionalDays {
		if IsSameDate(queryDate, termDate) {
			return true
		}
	}
	return false
}

// returns a slice consisting of the first of each month that is included in the term
func (t Term) TermMonths() ([]time.Time, error) {
	if t.Start.IsZero() || t.End.IsZero() {
		return nil, fmt.Errorf("term not initialized")
	}
	var dates []time.Time

	currDate := t.Start
	for !currDate.After(t.End.AddDate(0, 1, 0)) {
		first := time.Date(currDate.Year(), currDate.Month(), 1, 0, 0, 0, 0, time.Local)
		dates = append(dates, first)
		currDate = currDate.AddDate(0, 1, 0)
	}
	return dates, nil
}

func (t Term) Designation() string {
	return ""
}

type Date time.Time

type Dates []Date

func (dates Dates) SortAscending() {
	slices.SortFunc(dates, func(a, b Date) int {
		if IsSameDate(time.Time(a), time.Time(b)) {
			return 0
		}
		if time.Time(a).Before(time.Time(b)) {
			return -1
		} else {
			return 1
		}
	})
}

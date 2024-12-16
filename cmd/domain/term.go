package domain

import (
	"fmt"
	"time"
)

func NewTerm(start, end time.Time, nonInstructionalDays []time.Time, termType TermType, termID int, name string) (*Term, error) {
	if start.After(end) {
		return nil, fmt.Errorf("start must come before end: %s is after %s", start.String(), end.String())
	}
	term := Term{Start: start, End: end, NonInstructionalDays: nonInstructionalDays, TermType: termType, ID: termID, Name: name}
	instructionDates := InstructionDays(term)
	term.InstructionalDays = instructionDates
	return &term, nil
}

type TermType string

type Term struct {
	Start                time.Time
	End                  time.Time
	NonInstructionalDays []time.Time
	InstructionalDays    []time.Time
	TermType             TermType
	ID                   int
	Name                 string
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

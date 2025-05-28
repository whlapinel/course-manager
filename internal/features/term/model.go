package term

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	datespkg "gh_static_portfolio/internal/shared/dates"
	"slices"
	"time"
)

type Term struct {
	Start                time.Time   `json:"start"`
	End                  time.Time   `json:"end"`
	NonInstructionalDays []time.Time `json:"nonInstructionalDays"`
	InstructionalDays    []time.Time `json:"instructionalDays"`
	ports.BaseNode[string, int]
}

func (t Term) Designation() string {
	return ""
}

// this sets the instruction days field
func New(term Term) Term {
	term.InstructionalDays = InstructionDays(term)
	return term
}

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
	instructionDays := []time.Time{}
	for currDate := term.Start; !currDate.After(term.End); currDate = currDate.Add((24 * time.Hour)) {
		if currDate.Weekday() == time.Sunday || currDate.Weekday() == time.Saturday {
			continue
		}
		if slices.ContainsFunc(term.NonInstructionalDays, func(date time.Time) bool {
			return datespkg.IsSameDate(date, currDate)
		}) {
			continue
		}
		instructionDays = append(instructionDays, currDate)
	}
	return instructionDays
}

func (t Term) IsInstructionDay(queryDate time.Time) bool {
	if queryDate.Before(t.Start) || queryDate.After(t.End) {
		return false
	}
	if queryDate.Weekday() == time.Sunday || queryDate.Weekday() == time.Saturday {
		return false
	}
	for _, termDate := range t.InstructionalDays {
		if datespkg.IsSameDate(queryDate, termDate) {
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

type Date time.Time

type Dates []Date

func (dates Dates) SortAscending() {
	slices.SortFunc(dates, func(a, b Date) int {
		if datespkg.IsSameDate(time.Time(a), time.Time(b)) {
			return 0
		}
		if time.Time(a).Before(time.Time(b)) {
			return -1
		} else {
			return 1
		}
	})
}

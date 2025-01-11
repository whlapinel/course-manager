package domain

import (
	"fmt"
	"slices"
	"time"
)

func NewLesson(number, unitID int, name string, descr string, dates []time.Time) Lesson {
	return Lesson{Number: number, UnitID: unitID, Name: name, Description: descr, Dates: dates}
}

type Lesson struct {
	ID          int
	UnitID      int
	Number      int
	Name        string
	Description string
	Dates       []time.Time
	Image       Image
}

func (l Lesson) GetTitle() string {
	return l.Name
}

type CalendarDirection int

const (
	Right CalendarDirection = iota
	Left
)

var dirStringList = []string{"right", "left"}

func (d CalendarDirection) String() string {
	return dirStringList[d]
}

func ParseDirection(cd string) (CalendarDirection, error) {
	for i, word := range dirStringList {
		if cd == word {
			return CalendarDirection(i), nil
		}
	}
	return 0, fmt.Errorf("invalid direction value")
}

// removes all current instructional days and adds subsequent or previous instructional days from/to lesson.Dates,
// depending on direction argument (Left/Right).
// should work for lessons already spanning multiple dates.
// Returns modified Lesson and new date after shifting
// Returns an error if there is no subsequent or previous day
func (l Lesson) Shift(direction CalendarDirection, term Term) (Lesson, time.Time, error) {
	var shiftedDates []time.Time
	var shifted time.Time
	for _, date := range l.Dates {
		index := slices.Index(term.InstructionalDays, date)
		if direction == Right && index+1 < len(term.InstructionalDays) {
			shifted = term.InstructionalDays[index+1]
		} else if direction == Left && index-1 >= 0 {
			shifted = term.InstructionalDays[index-1]
		} else {
			return Lesson{}, time.Time{}, fmt.Errorf("lesson already occurs on first or last day of instruction, cannot shift (%s)", direction.String())
		}
		shiftedDates = append(shiftedDates, shifted)
	}
	l.Dates = shiftedDates
	return l, shifted, nil
}

// keeps all current instructional days and adds subsequent or previous instructional days from/to lesson.Dates,
// depending on direction argument (Left/Right).
// should work for lessons already spanning multiple dates.
// Returns an error if there is no subsequent or previous day
func (l Lesson) Extend(direction CalendarDirection, term Term) (Lesson, error) {
	var extendedDates []time.Time
	for _, date := range l.Dates {
		index := slices.Index(term.InstructionalDays, date)
		var extended time.Time
		if direction == Right && index+1 < len(term.InstructionalDays) {
			// append original date before extended date since the original is earlier
			extended = term.InstructionalDays[index+1]
			extendedDates = append(extendedDates, date)
			extendedDates = append(extendedDates, extended)
		} else if direction == Left && index-1 > 0 {
			// append original date after extended date since the original is later
			extended = term.InstructionalDays[index-1]
			extendedDates = append(extendedDates, extended)
			extendedDates = append(extendedDates, date)
		} else {
			return Lesson{}, fmt.Errorf("lesson already occurs on last day of instruction")
		}
	}
	l.Dates = extendedDates
	return l, nil
}

func (l Lesson) DedupeDates() Lesson {
	var dateMap = make(map[string]time.Time)
	for _, date := range l.Dates {
		dateMap[date.Format(time.DateOnly)] = date
	}
	var deduped []time.Time
	for _, date := range dateMap {
		deduped = append(deduped, date)
	}
	l.Dates = deduped
	return l
}

func (l Lesson) SortDates() Lesson {
	slices.SortFunc(l.Dates, func(d1 time.Time, d2 time.Time) int {
		if d1.Before(d2) {
			return 1
		}
		return -1
	})
	return l
}

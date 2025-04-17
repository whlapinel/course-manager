package domain

import (
	"slices"
	"time"
)

// for indicating special events such as teacher workdays, SAT testing, etc...
type Occasion struct {
	ID int
	// required paramater; all occasions will be part of a term
	TermID int
	Date   time.Time // required
	Name   string    // required
}

type Occasions []Occasion

// sorts in ascending order by date
func (o Occasions) Sort() {
	slices.SortFunc(o, func(a, b Occasion) int {
		if IsSameDate(a.Date, b.Date) {
			return 0
		}
		if a.Date.Before(b.Date) {
			return -1
		} else {
			return 1
		}
	})
}

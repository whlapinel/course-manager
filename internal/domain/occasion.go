package domain

import "time"

// for indicating special events such as teacher workdays, SAT testing, etc...
type Occasion struct {
	ID int
	// required paramater; all occasions will be part of a term
	TermID int
	Date   time.Time // required
	Name   string    // required
}

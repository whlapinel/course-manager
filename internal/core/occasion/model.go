package occasion

import "time"

// for indicating special events such as teacher workdays, SAT testing, etc...
type Occasion struct {
	ID       int
	ParentID int       // required paramater; all occasions will be part of a term or course
	Date     time.Time // required
	Name     string    // required
}

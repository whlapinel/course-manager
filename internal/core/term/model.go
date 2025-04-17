package term

import "time"

type Term struct {
	UserID               int
	Start                time.Time
	End                  time.Time
	NonInstructionalDays []time.Time
	InstructionalDays    []time.Time
	ID                   int
	Name                 string
	Description          string
}

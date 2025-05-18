package term

import "time"

type Term struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	Description          string      `json:"description"`
	UserID               string      `json:"userID"`
	Start                time.Time   `json:"start"`
	End                  time.Time   `json:"end"`
	NonInstructionalDays []time.Time `json:"nonInstructionalDays"`
	InstructionalDays    []time.Time `json:"instructionalDays"`
}

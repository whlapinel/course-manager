package lesson

import "time"

type Lesson struct {
	ID          int
	UnitID      int
	UnitNum     int
	UnitName    string
	Number      int
	Name        string
	Description string
	Dates       []time.Time
}

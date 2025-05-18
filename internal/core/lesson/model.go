package lesson

import "time"

type Lesson struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	UnitID      int         `json:"unitID"`
	UnitNumber  int         `json:"unitNumber"`
	Number      int         `json:"number"`
	UnitName    string      `json:"unitName"`
	Dates       []time.Time `json:"dates"`
}

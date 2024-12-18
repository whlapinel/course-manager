package domain

import "time"

func NewLesson(number, unitID int, name string, descr string, dates []time.Time) Lesson {
	return Lesson{Number: number, UnitID: unitID, Name: name, Description: descr, Dates: dates}
}

type Lesson struct {
	ID          int
	UnitID      int
	TemplateID  int
	Number      int
	Name        string
	Description string
	Dates       []time.Time
}

func (l Lesson) GetTitle() string {
	return l.Name
}

package domain

import "time"

func NewLesson(number, unitID int, name string, descr string, date time.Time) Lesson {
	return Lesson{Number: number, UnitID: unitID, Name: name, Description: descr, Date: date}
}

type Lesson struct {
	ID          int
	UnitID      int
	TemplateID  int
	Number      int
	Name        string
	Description string
	Date        time.Time
}

func (l Lesson) GetTitle() string {
	return l.Name
}

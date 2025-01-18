package domain

import "time"

type Assignment struct {
	LessonID         int
	Name             string
	Description      string
	InstructionsPath string
	RubricPath       string
	Category         AssignmentCategory
	DueDate          time.Time
}

type AssignmentCategory int

const (
	Perform AssignmentCategory = iota
	Rehearse
	Prepare
)

var catStrings = []string{"Perform, Rehearse, Prepare"}

func (cat AssignmentCategory) String() string {
	return catStrings[cat]
}

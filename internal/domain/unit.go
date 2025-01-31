package domain

import "fmt"

type NewUnitParams struct {
	Unit
}

func NewUnit(params NewUnitParams) Unit {
	return Unit{
		CourseID:    params.CourseID,
		Number:      params.Number,
		SequenceNum: params.SequenceNum,
		Name:        params.Name,
		Description: params.Description,
		Lessons:     params.Lessons,
	}

}

func (u Unit) GetName() string {
	return u.Name
}

func (u Unit) GetDescription() string {
	return u.Description
}

func (u Unit) GetID() int {
	return u.ID
}

func (u Unit) GetParentID() int {
	return u.CourseID
}

func (u Unit) Children() []CourseNode {
	var nodes []CourseNode
	for _, l := range u.Lessons {
		nodes = append(nodes, l)
	}
	return nodes
}

func (u Unit) TypeName() string {
	return "Unit"

}

func (u Unit) ParentTypeName() string {
	return "Course"
}

func (u Unit) ChildTypeName() string {
	return "Lesson"
}

// Always associated with a particular course
type Unit struct {
	ID          int
	CourseID    int
	Number      int
	SequenceNum int
	Name        string
	Description string
	Lessons     []Lesson
	Image       Image
}

func (u Unit) AddLesson(lesson Lesson) {
	u.Lessons = append(u.Lessons, lesson)
}

func (u Unit) Designation() string {
	return fmt.Sprintf("Unit %d", u.Number)
}

package dto

import (
	"fmt"
	"gh_static_portfolio/internal/app/traversal"
	"gh_static_portfolio/internal/core/unit"
)

type Unit struct {
	unit.Unit
	Lessons []Lesson
}

func (u Unit) GetName() string {
	return u.Name
}

func (u Unit) GetDescription() string {
	return u.Description
}

func (u Unit) GetID() any {
	return u.ID
}

func (u Unit) GetParentID() int {
	return u.CourseID
}

func (u Unit) Children() []traversal.CourseNode {
	var nodes []traversal.CourseNode
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

func (u Unit) Designation() string {
	if u.Number < 0 {
		return "N/A"
	}
	return fmt.Sprintf("Unit %d", u.Number)
}

func (u Unit) GetNumber() int {
	return u.Number
}

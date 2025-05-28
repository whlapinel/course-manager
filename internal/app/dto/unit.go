package dto

import (
	"gh_static_portfolio/internal/features/unit"
	"gh_static_portfolio/internal/ports"
)

type Unit struct {
	unit.Unit `json:"unit"`
	Lessons   []Lesson `json:"lessons"`
}

func (u Unit) GetChildren() []ports.Node {
	var nodes []ports.Node
	for _, l := range u.Lessons {
		nodes = append(nodes, l)
	}
	return nodes
}

func (u Unit) GetParentTypeName() string {
	return CourseTypeName.String()
}

func (u Unit) GetTypeName() string {
	return UnitTypeName.String()
}

func (t Unit) GetChildTypeName() string {
	return LessonTypeName.String()
}

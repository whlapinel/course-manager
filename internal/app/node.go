package app

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/traversal"
	"slices"
)

type NodePath struct {
	UserID   string
	TermID   int
	CourseID int
	UnitID   int
	LessonID int
}

func (path NodePath) ToSlice() []any {
	var pathSlice []any
	params := []any{path.UserID, path.TermID, path.CourseID, path.UnitID, path.LessonID}
	for _, param := range params {
		switch v := param.(type) {
		case int:
			if v != 0 {
				pathSlice = append(pathSlice, param)
			}
		case string:
			if v != "" {
				pathSlice = append(pathSlice, param)

			}
		}
	}
	return pathSlice
}

// for storing the actual nodes
type Nodes struct {
	User   dto.User
	Term   dto.Term
	Course dto.Course
	Unit   dto.Unit
	Lesson dto.Lesson
}

func (nodes Nodes) ToSlice(additional ...traversal.CourseNode) []traversal.CourseNode {
	var nodeSlice []traversal.CourseNode
	params := []traversal.CourseNode{nodes.User, nodes.Term, nodes.Course, nodes.Unit, nodes.Lesson}
	for _, param := range params {
		switch v := param.GetID().(type) {
		case string:
			if v != "" {
				nodeSlice = append(nodeSlice, param)
			}
		case int:
			if v != 0 {
				nodeSlice = append(nodeSlice, param)
			}
		}
	}
	nodeSlice = append(nodeSlice, additional...)
	return nodeSlice
}

// last node in node slice represents current node
func (path Nodes) CurrentNode() traversal.CourseNode {
	nodeSlice := path.ToSlice()
	length := len(nodeSlice)
	if length == 0 {
		return path.User
	}
	return nodeSlice[length-1]
}

type NodeSorter interface {
	GetNumber() int
}

// Generic version that works with any slice of T where T implements NodeSorter
func SortByNumber[T NodeSorter](nodes []T) {
	slices.SortStableFunc(nodes, func(a, b T) int {
		if a.GetNumber() == b.GetNumber() {
			return 0
		}
		if a.GetNumber() < b.GetNumber() {
			return -1
		}
		return 1
	})
}

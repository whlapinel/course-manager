package domain

import (
	"slices"
)

type NodeRepository[T CourseNode] interface {
	GetByID(id any) (T, error)
	GetByParentID(parentID any) (T, error)
	Save(node T) (id any, err error)
	Delete(id any) error
}

// Tree node implemented by Term, Course, Unit, Lesson
type CourseNode interface {
	GetID() any // could be string or int
	GetName() string
	GetNumber() int
	GetParentID() int
	GetDescription() string
	Children() []CourseNode
	TypeName() string
	ParentTypeName() string
	ChildTypeName() string
	Designation() string
}

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
	User   User
	Term   Term
	Course Course
	Unit   Unit
	Lesson Lesson
}

func (nodes Nodes) ToSlice(additional ...CourseNode) []CourseNode {
	var nodeSlice []CourseNode
	params := []CourseNode{nodes.User, nodes.Term, nodes.Course, nodes.Unit, nodes.Lesson}
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
func (path Nodes) CurrentNode() CourseNode {
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

type NodeTypeName string

const (
	RootTypeName   NodeTypeName = "Root"
	UserTypeName   NodeTypeName = "User"
	TermTypeName   NodeTypeName = "Term"
	CourseTypeName NodeTypeName = "Course"
	UnitTypeName   NodeTypeName = "Unit"
	LessonTypeName NodeTypeName = "Lesson"
)

func (n NodeTypeName) String() string {
	return string(n)
}

type RootCourseNode struct {
	Users []User
}

func (root RootCourseNode) GetNumber() int {
	return -1
}

func (root RootCourseNode) GetID() interface{} {
	return 0
}

func (r RootCourseNode) GetName() string {
	return ""
}
func (r RootCourseNode) GetParentID() int {
	return 0
}
func (r RootCourseNode) GetDescription() string {
	return ""
}
func (r RootCourseNode) Children() []CourseNode {
	var nodes []CourseNode
	for _, user := range r.Users {
		nodes = append(nodes, &user)
	}
	return nodes
}

func (r RootCourseNode) TypeName() string {
	return RootTypeName.String()

}
func (r RootCourseNode) ParentTypeName() string {
	return ""
}
func (r RootCourseNode) ChildTypeName() string {
	return TermTypeName.String()
}
func (r RootCourseNode) Designation() string {
	return ""
}

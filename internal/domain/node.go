package domain

import "slices"

type NodeRepository[T CourseNode] interface {
	GetByID(id interface{}) (T, error)
	GetByParentID(parentID interface{}) (T, error)
	Save(node T) (id interface{}, err error)
	Delete(id interface{}) error
}

// Tree node implemented by Term, Course, Unit, Lesson
type CourseNode interface {
	GetID() interface{} // could be string or int
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

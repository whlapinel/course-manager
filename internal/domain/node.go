package domain

// Tree node implemented by Term, Course, Unit, Lesson
type CourseNode interface {
	GetID() int
	GetName() string
	GetParentID() int
	GetDescription() string
	Children() []CourseNode
	TypeName() string
	ParentTypeName() string
	ChildTypeName() string
	Designation(parentNum int) string
}

type NodeTypeName string

const (
	RootTypeName   NodeTypeName = "Root"
	TermTypeName   NodeTypeName = "Term"
	CourseTypeName NodeTypeName = "Course"
	UnitTypeName   NodeTypeName = "Unit"
	LessonTypeName NodeTypeName = "Lesson"
)

func (n NodeTypeName) String() string {
	return string(n)
}

func NodeDesignation(node, parent CourseNode) string {
	if unit, ok := node.(Unit); ok {
		return unit.Designation(0)
	}
	if lesson, ok := node.(Lesson); ok {
		if unit, ok := parent.(Unit); ok {
			return lesson.Designation(unit.Number)
		}
		return "Error: parent is not a unit"
	}
	return ""
}

type RootCourseNode struct {
	Terms []Term
}

func (root RootCourseNode) GetNumber() int {
	return -1
}

func (root RootCourseNode) GetID() int {
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
	return []CourseNode{}
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
func (r RootCourseNode) Designation(_ int) string {
	return ""
}

package dto

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

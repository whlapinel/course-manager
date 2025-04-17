package traversal

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

package templates

type Node interface {
	GetID() any // could be string or int
	GetName() string
	GetNumber() int
	GetParentID() int
	GetDescription() string
	Children() []Node
	TypeName() string
	ParentTypeName() string
	ChildTypeName() string
	Designation() string
}

type Nodes struct {
	User   Node
	Term   Node
	Course Node
	Unit   Node
	Lesson Node
}

func (nodes Nodes) ToSlice() []Node {
	return []Node{nodes.User, nodes.Term, nodes.Course, nodes.Unit, nodes.Lesson}
}

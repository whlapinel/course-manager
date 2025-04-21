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

package ports

type Node interface {
	GetID() any // could be string or int
	GetName() string
	GetNumber() int
	GetParentID() any
	GetDescription() string
	GetChildren() []Node
	GetTypeName() string
	GetParentTypeName() string
	GetChildTypeName() string
	Designation() string
	GetSequence() int
}

type Intorstring interface {
	~int | ~string
}

type BaseNode[ParentID Intorstring, ID Intorstring] struct {
	ID          ID       `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Number      int      `json:"number"`
	Sequence    int      `json:"sequence"`
	ParentID    ParentID `json:"parentID"`
}

func (b BaseNode[ParentID, ID]) GetID() any {
	return b.ID
}

func (b BaseNode[ParentID, ID]) GetName() string {
	return b.Name
}

func (b BaseNode[ParentID, ID]) GetNumber() int {
	return b.Number
}

func (b BaseNode[ParentID, ID]) GetParentID() any {
	return b.ParentID
}

func (b BaseNode[ParentID, ID]) GetDescription() string {
	return b.Description
}

func (b BaseNode[ParentID, ID]) Designation() string {
	return ""
}

func (b BaseNode[ParentID, ID]) GetTypeName() string {
	return ""
}
func (b BaseNode[ParentID, ID]) GetParentTypeName() string {
	return ""
}
func (b BaseNode[ParentID, ID]) GetChildTypeName() string {
	return ""
}
func (b BaseNode[ParentID, ID]) GetSequence() int {
	return -1
}

type Nodes struct {
	User   Node
	Term   Node
	Course Node
	Unit   Node
	Lesson Node
}

func (nodes Nodes) ToSlice(addParams ...any) []Node {
	return []Node{nodes.User, nodes.Term, nodes.Course, nodes.Unit, nodes.Lesson}
}

func (nodes Nodes) CurrentNode() Node {
	nodeSlice := nodes.ToSlice()
	for i, node := range nodeSlice {
		if node == nil {
			return nodeSlice[i-1]
		}
	}
	return nodeSlice[len(nodeSlice)-1]
}

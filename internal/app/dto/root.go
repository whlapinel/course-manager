package dto

import (
	"gh_static_portfolio/internal/shared/node"
)

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
func (r RootCourseNode) Children() []node.Node {
	var nodes []node.Node
	for _, user := range r.Users {
		nodes = append(nodes, user)
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

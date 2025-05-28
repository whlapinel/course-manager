package dto

import (
	"gh_static_portfolio/internal/features/user"
	"gh_static_portfolio/internal/ports"
)

type User struct {
	user.User `json:"user"`
	Terms     []Term `json:"terms"`
}

// GetChildren implements CourseNode.
func (u User) GetChildren() []ports.Node {
	var nodes []ports.Node
	for _, term := range u.Terms {
		nodes = append(nodes, term)
	}
	return nodes
}

func (u User) GetParentTypeName() string {
	return RootTypeName.String()
}

func (u User) GetTypeName() string {
	return UserTypeName.String()
}

func (u User) GetChildTypeName() string {
	return TermTypeName.String()
}

func (u User) GetNumber() int {
	return -1
}

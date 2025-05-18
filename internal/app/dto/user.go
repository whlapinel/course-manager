package dto

import (
	"fmt"
	"gh_static_portfolio/internal/core/user"
	"gh_static_portfolio/internal/ports"
	"strings"
)

type User struct {
	user.User `json:"user"`
	Terms     []Term `json:"terms"`
}

func (u User) Username() string {
	return strings.ToLower(u.FirstName[:1] + u.LastName)
}

// ChildTypeName implements CourseNode.
func (u User) ChildTypeName() string {
	return TermTypeName.String()
}

// Children implements CourseNode.
func (u User) Children() []ports.Node {
	var nodes []ports.Node
	for _, term := range u.Terms {
		nodes = append(nodes, term)
	}
	return nodes
}

// Designation implements CourseNode.
func (u User) Designation() string {
	return ""
}

// GetDescription implements CourseNode.
func (u User) GetDescription() string {
	return ""
}

// GetID implements CourseNode.
func (u User) GetID() any {
	return u.ID
}

// GetName implements CourseNode.
func (u User) GetName() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}

// GetNumber implements CourseNode.
func (u User) GetNumber() int {
	return -1
}

// GetParentID implements CourseNode.
func (u User) GetParentID() any {
	return -1
}

// ParentTypeName implements CourseNode.
func (u User) ParentTypeName() string {
	return UserTypeName.String()
}

// TypeName implements CourseNode.
func (u User) TypeName() string {
	return UserTypeName.String()
}

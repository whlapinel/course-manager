package domain

import "fmt"

type User struct {
	ID                         int
	Email, FirstName, LastName string
	Terms                      []Term
}

// ChildTypeName implements CourseNode.
func (u *User) ChildTypeName() string {
	return TermTypeName.String()
}

// Children implements CourseNode.
func (u *User) Children() []CourseNode {
	var nodes []CourseNode
	for _, term := range u.Terms {
		nodes = append(nodes, term)
	}
	return nodes
}

// Designation implements CourseNode.
func (u *User) Designation() string {
	return ""
}

// GetDescription implements CourseNode.
func (u *User) GetDescription() string {
	return ""
}

// GetID implements CourseNode.
func (u *User) GetID() int {
	return u.ID
}

// GetName implements CourseNode.
func (u *User) GetName() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}

// GetNumber implements CourseNode.
func (u *User) GetNumber() int {
	return -1
}

// GetParentID implements CourseNode.
func (u *User) GetParentID() int {
	return -1
}

// ParentTypeName implements CourseNode.
func (u *User) ParentTypeName() string {
	return RootTypeName.String()
}

// TypeName implements CourseNode.
func (u *User) TypeName() string {
	return UserTypeName.String()
}

package user

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	"strings"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Picture   string `json:"picture"`
	ports.BaseNode[int, string]
}

func (u User) GetName() string {
	return fmt.Sprintf("%s. %s", u.FirstName[:1], u.LastName)
}

func (u User) Username() string {
	return strings.ToLower(u.FirstName[:1] + u.LastName)
}

// Designation implements CourseNode.
func (u User) Designation() string {
	return ""
}

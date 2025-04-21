package user

import (
	"gh_static_portfolio/internal/core/user"
)

type Repository interface {
	ByID(id string) (user.User, error)
	Save(user user.User) (string, error) // returns id
	All() ([]user.User, error)
}

package user

import (
	"gh_static_portfolio/internal/core/user"
)

type Repository interface {
	All() ([]user.User, error)
}

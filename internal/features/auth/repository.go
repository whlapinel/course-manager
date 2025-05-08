package auth

import "gh_static_portfolio/internal/core/user"

type Repository interface {
	ByID(id string) (user.User, error)
}

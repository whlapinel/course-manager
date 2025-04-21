package auth

import "gh_static_portfolio/internal/core/user"

type UserReader interface {
	ByID(id string) (user.User, error)
}

type Service struct {
	repo      Repository
	userQuery UserReader
}

func NewService(repo Repository, userQuery UserReader) *Service {
	return &Service{
		repo:      repo,
		userQuery: userQuery,
	}
}

func (svc *Service) GetUser(id string) (user.User, error) {
	return svc.userQuery.ByID(id)
}



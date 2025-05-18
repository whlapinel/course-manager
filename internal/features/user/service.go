package user

import (
	"gh_static_portfolio/internal/core/user"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) ByID(id string) (user.User, error) {
	return svc.repo.ByID(id)
}

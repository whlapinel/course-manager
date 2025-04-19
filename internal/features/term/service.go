package term

import "gh_static_portfolio/internal/core/term"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) TermsByUser(userID string) ([]term.Term, error) {
	return svc.repo.ByUserID(userID)
}

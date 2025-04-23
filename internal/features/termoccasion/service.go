package termoccasion

import "gh_static_portfolio/internal/core/occasion"

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) ByTermID(termID int) ([]occasion.Occasion, error) {
	return svc.repo.ByTermID(termID)
}

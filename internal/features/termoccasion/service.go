package termoccasion

import "gh_static_portfolio/internal/core/occasion"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) ByTermID(termID int) ([]occasion.Occasion, error) {
	return svc.repo.ByTermID(termID)
}

func (svc *Service) ByID(occasionID int) (occasion.Occasion, error) {
	return svc.repo.ByID(occasionID)
}

func (svc *Service) Create(occ occasion.Occasion) (int, error) {
	return svc.repo.Save(occ)
}

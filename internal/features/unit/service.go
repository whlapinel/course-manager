package unit

import "gh_static_portfolio/internal/core/unit"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) ByID(unitID int) (unit.Unit, error) {
	return svc.repo.ByID(unitID)
}

func (svc *Service) ByCourseID(courseID int) ([]unit.Unit, error) {
	return svc.repo.ByCourseID(courseID)
}

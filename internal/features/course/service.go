package course

import "gh_static_portfolio/internal/core/course"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) ByID(courseID int) (course.Course, error) {
	return svc.repo.ByID(courseID)
}

func (svc *Service) ByTermID(termID int) ([]course.Course, error) {
	return svc.repo.ByTermID(termID)
}

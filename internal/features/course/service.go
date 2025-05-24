package course

import (
	"gh_static_portfolio/internal/core/course"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) Update(course course.Course) error {
	return svc.repo.Update(course)
}

func (svc *Service) ByID(courseID int) (course.Course, error) {
	return svc.repo.ByID(courseID)
}

func (svc *Service) ByTermID(termID int) ([]course.Course, error) {
	return svc.repo.ByTermID(termID)
}

func (svc *Service) Delete(courseID int) error {
	return svc.repo.Delete(courseID)
}

func (svc *Service) Save(course course.Course) error {
	_, err := svc.repo.Save(course)
	if err != nil {
		return err
	}
	return nil
}

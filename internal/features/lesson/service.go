package lesson

import (
	"gh_static_portfolio/internal/core/lesson"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) Update(lesson lesson.Lesson) error{
	return svc.repo.Update(lesson)
}

func (svc *Service) ByID(lessonID int) (lesson.Lesson, error) {
	return svc.repo.ByID(lessonID)
}

func (svc *Service) ByUnitID(unitID int) ([]lesson.Lesson, error) {
	return svc.repo.ByUnitID(unitID)
}

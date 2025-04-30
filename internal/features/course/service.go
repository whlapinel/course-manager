package course

import (
	"gh_static_portfolio/internal/app/dto"
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

func (svc *Service) ByID(courseID int) (course.Course, error) {
	return svc.repo.ByID(courseID)
}

func (svc *Service) ByTermID(termID int) ([]dto.Course, error) {
	courses, err := svc.repo.ByTermID(termID)
	if err != nil {
		return nil, err
	}
	var courseDTOs []dto.Course
	for _, course := range courses {
		courseDTO := dto.Course{
			Course: course,
		}
		courseDTOs = append(courseDTOs, courseDTO)
	}
	return courseDTOs, nil
}

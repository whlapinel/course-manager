package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/term"
)

type TermService struct {
	termService   *term.Service
	courseService *course.Service
}

func NewTermService(termSvc *term.Service, courseSvc *course.Service) *TermService {
	return &TermService{
		termService:   termSvc,
		courseService: courseSvc,
	}
}
func (svc *TermService) ByID(termID int) (dto.Term, error) {
	term, err := svc.termService.ByID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	return dto.Term{
		Term: term,
	}, nil
}

func (svc *TermService) ListCourses(termID int) (dto.Term, error) {
	term, err := svc.termService.ByID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	courses, err := svc.courseService.ByTermID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	var courseDTOs []dto.Course
	for _, course := range courses {
		courseDTO := dto.Course{
			Course: course,
		}
		courseDTOs = append(courseDTOs, courseDTO)
	}
	termDTO := dto.Term{
		Term:    term,
		Courses: courseDTOs,
	}
	return termDTO, nil
}

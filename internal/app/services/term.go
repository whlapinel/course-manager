package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/termoccasion"
)

type TermService struct {
	termService   *term.Service
	occasionSvc   *termoccasion.Service
	courseService *course.Service
	fileService   *FileService
}

func NewTermService(
	termSvc *term.Service,
	occasionSvc *termoccasion.Service,
	courseSvc *course.Service,
) *TermService {
	return &TermService{
		termService:   termSvc,
		occasionSvc:   occasionSvc,
		courseService: courseSvc,
	}
}

func (svc *TermService) ByID(termID int) (dto.Term, error) {
	term, err := svc.termService.ByID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	occasions, err := svc.occasionSvc.ByTermID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	courses, err := svc.courseService.ByTermID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	return dto.Term{
		Term:      term,
		Occasions: occasions,
		Courses:   courses,
	}, nil
}

func (svc *TermService) ListCourses(termID int) (dto.Term, error) {
	term, err := svc.termService.ByID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	courseDTOs, err := svc.courseService.ByTermID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	termDTO := dto.Term{
		Term:    term,
		Courses: courseDTOs,
	}
	return termDTO, nil
}

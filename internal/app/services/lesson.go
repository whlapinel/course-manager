package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/features/slides"
)

type LessonService struct {
	lessonService *lesson.Service
	slidesService *slides.Service
}

func NewLessonService(lessonSvc *lesson.Service, slidesService *slides.Service) *LessonService {
	return &LessonService{
		lessonService: lessonSvc,
		slidesService: slidesService,
	}
}

func (svc *LessonService) Update(lesson dto.Lesson) error {
	return svc.lessonService.Update(lesson.Lesson)
}

func (svc *LessonService) ByID(lessonID int) (dto.Lesson, error) {
	lesson, err := svc.lessonService.ByID(lessonID)
	if err != nil {
		return dto.Lesson{}, err
	}
	return dto.Lesson{
		Lesson: lesson,
	}, nil
}

func (svc *LessonService) ByUnitID(unitID int) ([]dto.Lesson, error) {
	lessons, err := svc.lessonService.ByUnitID(unitID)
	if err != nil {
		return nil, err
	}
	var lessonDTOs []dto.Lesson
	for _, lesson := range lessons {
		lessonDTO := dto.Lesson{
			Lesson: lesson,
		}
		lessonDTOs = append(lessonDTOs, lessonDTO)
	}
	return lessonDTOs, nil
}

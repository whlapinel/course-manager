package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/lesson"
)

type LessonService struct {
	lessonService *lesson.Service
}

func NewLessonService(lessonSvc *lesson.Service) *LessonService {
	return &LessonService{
		lessonService: lessonSvc,
	}
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

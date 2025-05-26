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

func (svc *LessonService) ByParentID(unitID int) ([]dto.Lesson, error) {
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
func (svc *LessonService) Delete(lessonID int) error {
	return svc.lessonService.Delete(lessonID)
}

func (svc *LessonService) Save(lesson dto.Lesson) error {
	return svc.lessonService.Save(lesson.Lesson)
}

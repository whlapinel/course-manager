package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/features/unit"
)

type UnitService struct {
	unitService   *unit.Service
	lessonService *lesson.Service
}

func NewUnitService(unitSvc *unit.Service, lessonSvc *lesson.Service) *UnitService {
	return &UnitService{
		unitService:   unitSvc,
		lessonService: lessonSvc,
	}
}

func (svc *UnitService) ListLessons(unitID int) (dto.Unit, error) {
	unit, err := svc.unitService.ByID(unitID)
	if err != nil {
		return dto.Unit{}, err
	}
	lessons, err := svc.lessonService.ByUnitID(unitID)
	if err != nil {
		return dto.Unit{}, err
	}
	var lessonDTOs []dto.Lesson
	for _, lesson := range lessons {
		lessonDTO := dto.Lesson{
			Lesson: lesson,
		}
		lessonDTOs = append(lessonDTOs, lessonDTO)
	}
	unitDTO := dto.Unit{
		Unit:    unit,
		Lessons: lessonDTOs,
	}
	return unitDTO, nil

}

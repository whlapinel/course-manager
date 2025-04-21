package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/unit"
)

type CourseService struct {
	courseService *course.Service
	unitService   *unit.Service
}

func NewCourseService(courseSvc *course.Service, unitSvc *unit.Service) *CourseService {
	return &CourseService{
		courseService: courseSvc,
		unitService:   unitSvc,
	}
}

func (svc *CourseService) ByID(courseID int) (dto.Course, error) {
	course, err := svc.courseService.ByID(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	return dto.Course{
		Course: course,
	}, nil
}

func (svc *CourseService) ListUnits(courseID int) (dto.Course, error) {
	course, err := svc.courseService.ByID(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	units, err := svc.unitService.ByCourseID(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	var unitDTOs []dto.Unit
	for _, unit := range units {
		unitDTO := dto.Unit{
			Unit: unit,
		}
		unitDTOs = append(unitDTOs, unitDTO)
	}
	courseDTO := dto.Course{
		Course: course,
		Units:  unitDTOs,
	}
	return courseDTO, nil

}

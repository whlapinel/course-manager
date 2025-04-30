package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/course"
)

type CourseService struct {
	courseService *course.Service
	getUnits      func(courseID int) ([]dto.Unit, error)
}

func NewCourseService(
	courseSvc *course.Service,
	getUnits func(courseID int) ([]dto.Unit, error),

) *CourseService {
	return &CourseService{
		courseService: courseSvc,
		getUnits:      getUnits,
	}
}

func (svc *CourseService) ByID(courseID int) (dto.Course, error) {
	course, err := svc.courseService.ByID(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	units, err := svc.getUnits(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	return dto.Course{
		Course: course,
		Units:  units,
	}, nil
}

func (svc *CourseService) ListUnits(courseID int) (dto.Course, error) {
	course, err := svc.courseService.ByID(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	units, err := svc.getUnits(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	courseDTO := dto.Course{
		Course: course,
		Units:  units,
	}
	return courseDTO, nil

}

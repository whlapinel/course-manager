package services

import (
	"gh_static_portfolio/internal/app/dto"
)

type CourseCalendarService struct {
	course func(courseID int) (dto.Course, error)
}

func NewCourseCalendarService(course func(courseID int) (dto.Course, error)) *CourseCalendarService {
	return &CourseCalendarService{
		course: course,
	}
}

func (svc *CourseCalendarService) Course(courseID int) (dto.Course, error) {
	courseDTO, err := svc.course(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	return courseDTO, nil
}

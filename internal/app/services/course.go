package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/course"
)

type CourseService struct {
	courseService *course.Service
}

func NewCourseService(
	courseSvc *course.Service,

) *CourseService {
	return &CourseService{
		courseService: courseSvc,
	}
}

func (svc *CourseService) ListByTerm(termID int) ([]dto.Course, error) {
	return svc.courseService.ByTermID(termID)
}

func (svc *CourseService) Update(course dto.Course) error {
	return svc.courseService.Update(course.Course)
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

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

func (svc *CourseService) Save(course dto.Course) error {
	return svc.courseService.Save(course.Course)
}

func (svc *CourseService) ByParentID(termID int) ([]dto.Course, error) {
	courses, err := svc.courseService.ByTermID(termID)
	if err != nil {
		return nil, err
	}
	var courseDTOs []dto.Course
	for _, course := range courses {
		courseDTO := dto.Course{
			Course: course,
		}
		courseDTOs = append(courseDTOs, courseDTO)
	}
	return courseDTOs, nil
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

func (svc *CourseService) Delete(courseID int) error {
	return svc.courseService.Delete(courseID)

}

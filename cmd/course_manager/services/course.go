package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseService interface {
	CreateTemplate(template *domain.CourseTemplate) error
	CreateInstance(course *domain.CourseInstance) error
	GetTemplates() ([]*domain.CourseTemplate, error)
	GetInstances() ([]*domain.CourseInstance, error)
	ReadFromCSV() ([]*domain.CourseInstance, error)
}

type courseService struct {
	repo domain.CourseRepo
}

// ReadFromCSV implements CourseService.
func (svc courseService) ReadFromCSV() ([]*domain.CourseInstance, error) {
	panic("unimplemented")
}

func NewCourseService(courseRepo domain.CourseRepo) CourseService {
	return courseService{repo: courseRepo}

}

func (svc courseService) CreateTemplate(course *domain.CourseTemplate) error {
	_, err := svc.repo.SaveTemplate(course)
	return err
}
func (svc courseService) CreateInstance(course *domain.CourseInstance) error {
	instance := course.CreateInstance()
	return svc.repo.SaveInstance(instance)
}

func (svc courseService) GetTemplates() ([]*domain.CourseTemplate, error) {
	return svc.repo.GetTemplates()
}
func (svc courseService) GetInstances() ([]*domain.CourseInstance, error) {
	return svc.repo.GetInstances()
}

func (svc courseService) ReadInstancesFromCSV() ([]*domain.CourseTemplate, error) {
	courses, err := svc.repo.ReadFromCSV()
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		svc.repo.SaveTemplate(course)
	}
	return courses, nil
}

func (svc courseService) Update(course *domain.CourseTemplate) error {
	return fmt.Errorf("not implemented")
}

func (svc courseService) Delete(course *domain.CourseTemplate) error {
	return fmt.Errorf("not implemented")
}

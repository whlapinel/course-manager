package services

import (
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
)

type CourseService struct {
	courses data.CourseRepo
}

func NewCourseService(courseRepo data.CourseRepo) CourseService {
	return CourseService{courses: courseRepo}

}

func (svc CourseService) CreateTemplate(course domain.CourseTemplate) error {
	_, err := svc.courses.SaveTemplate(course)
	return err
}
func (svc CourseService) CreateInstance(course domain.CourseTemplate, term domain.Term) error {
	instance := course.CreateInstance(term)
	return svc.courses.SaveInstance(instance)
}

func (svc CourseService) ImportTemplate(filename string) ([]domain.CourseTemplate, error) {
	templates, err := svc.courses.ImportTemplatesFromCSV()
	if err != nil {
		return []domain.CourseTemplate{}, err
	}
	return templates, nil
}

func (svc CourseService) GetTemplates() ([]domain.CourseTemplate, error) {
	return svc.courses.GetTemplates()
}
func (svc CourseService) GetInstances(term domain.Term) ([]domain.CourseInstance, error) {
	return svc.courses.GetInstances(term)
}

func (svc CourseService) Update(course domain.CourseTemplate) error {
	return fmt.Errorf("not implemented")
}

func (svc CourseService) Delete(course domain.CourseTemplate) error {
	return fmt.Errorf("not implemented")
}

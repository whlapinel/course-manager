package service

import (
	"gh_static_portfolio/cmd/domain"
	"log"
)

func (svc CourseService) GetTemplates() ([]domain.Course, error) {
	templates, err := svc.repo.GetTemplates()
	if err != nil {
		return nil, err
	}
	log.Println(len(templates))
	return templates, nil
}
func (svc CourseService) UpdateCourseTemplate(tpl domain.Course) error {
	err := svc.repo.UpdateTemplate(tpl)
	if err != nil {
		return err
	}
	return nil
}

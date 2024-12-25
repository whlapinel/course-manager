package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) UpdateLesson(l domain.Lesson) error {
	err := svc.repo.UpdateLesson(l)
	if err != nil {
		return err
	}
	return nil
}

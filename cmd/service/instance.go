package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) GetInstances(termID int) ([]domain.CourseInstance, error) {
	return svc.repo.GetInstances(termID)
}

func (svc CourseService) UpdateCourseInstance(instance domain.CourseInstance) error {
	return svc.repo.UpdateInstance(instance)
}

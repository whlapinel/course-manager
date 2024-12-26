package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) GetCourses(termID int) (domain.Courses, error) {
	return svc.repo.GetCourses(termID)
}

func (svc CourseService) UpdateCourse(instance domain.Course) error {
	return svc.repo.UpdateInstance(instance)
}

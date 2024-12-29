package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) GetSchedule(course domain.Course) (domain.CourseSchedule, error) {
	var schedule domain.CourseSchedule
	schedule, err := svc.repo.GetSchedule(course)
	if err != nil {
		return schedule, err
	}
	return schedule, nil
}

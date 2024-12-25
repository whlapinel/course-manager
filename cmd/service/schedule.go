package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) GetSchedule(instance domain.CourseInstance) (domain.CourseSchedule, error) {
	var schedule domain.CourseSchedule
	schedule, err := svc.repo.GetSchedule(instance)
	if err != nil {
		return schedule, err
	}
	return schedule, nil
}

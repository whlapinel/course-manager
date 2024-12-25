package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) UpdateUnit(u domain.Unit) error {
	err := svc.repo.UpdateUnit(u)
	if err != nil {
		return err
	}
	return nil
}

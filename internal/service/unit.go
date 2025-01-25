package service

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
)

func (svc CourseService) GetUnit(unitID int) (domain.Unit, error) {
	unit, err := svc.repo.GetUnit(unitID)
	if err != nil {
		return domain.Unit{}, fmt.Errorf("CourseService.repo.GetUnit: %s", err)
	}
	lessons, err := svc.repo.GetLessons(unitID)
	if err != nil {
		return domain.Unit{}, fmt.Errorf("CourseService.repo.GetLessons: %s", err)
	}
	unit.Lessons = lessons
	return unit, nil
}
func (svc CourseService) GetUnits(courseID int) ([]domain.Unit, error) {
	units, err := svc.repo.GetUnits(courseID)
	if err != nil {
		return nil, err
	}
	return units, nil
}
func (svc CourseService) UpdateUnit(u domain.Unit) error {
	err := svc.repo.UpdateUnit(u)
	if err != nil {
		return err
	}
	return nil
}

type SaveUnitParams struct {
	domain.Unit
}

func (svc CourseService) SaveUnit(params SaveUnitParams) (domain.Unit, error) {
	newUnit := domain.NewUnit(domain.NewUnitParams{
		Unit: params.Unit,
	})
	unit, err := svc.repo.SaveUnit(newUnit)
	if err != nil {
		return domain.Unit{}, err
	}
	return unit, nil

}

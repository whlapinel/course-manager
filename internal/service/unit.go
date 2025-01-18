package service

import (
	"database/sql"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/domain"
)

func (svc CourseService) GetUnit(unitID int) (*domain.Unit, error) {
	unit, err := svc.repo.GetUnit(unitID)
	if err != nil {
		return nil, fmt.Errorf("CourseService.repo.GetUnit: %s", err)
	}
	lessons, err := svc.repo.GetLessons(unitID)
	if err != nil {
		return nil, fmt.Errorf("CourseService.repo.GetLessons: %s", err)
	}
	unit.Lessons = lessons
	image, err := svc.repo.GetUnitImage(unitID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("CourseService.repo.GetUnitImage")
		}
	} else if image != nil {
		unit.Image = *image
	}
	return unit, nil
}
func (svc CourseService) GetUnits(courseID int) ([]*domain.Unit, error) {
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

func (svc CourseService) CreateUnit(u domain.Unit) (domain.Unit, error) {
	unit, err := svc.repo.SaveUnit(u)
	if err != nil {
		return domain.Unit{}, err
	}
	return unit, nil

}

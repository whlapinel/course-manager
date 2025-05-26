package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/unit"
)

type UnitService struct {
	unitService *unit.Service
}

func NewUnitService(unitSvc *unit.Service) *UnitService {
	return &UnitService{
		unitService: unitSvc,
	}
}

func (svc *UnitService) Delete(unitID int) error {
	return svc.unitService.Delete(unitID)
}
func (svc *UnitService) Save(unit dto.Unit) error {
	return svc.unitService.Save(unit.Unit)
}
func (svc *UnitService) Update(unit dto.Unit) error {
	return svc.unitService.Update(unit.Unit)
}

func (svc *UnitService) ByParentID(courseID int) ([]dto.Unit, error) {
	units, err := svc.unitService.ByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	var unitDTOs []dto.Unit
	for _, unit := range units {
		unitDTO := dto.Unit{
			Unit: unit,
		}
		unitDTOs = append(unitDTOs, unitDTO)
	}
	return unitDTOs, nil

}

func (svc *UnitService) ByID(unitID int) (dto.Unit, error) {
	unit, err := svc.unitService.ByID(unitID)
	if err != nil {
		return dto.Unit{}, err
	}
	return dto.Unit{
		Unit: unit,
	}, nil
}

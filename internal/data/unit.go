package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"os"
	"path/filepath"
)

func UnitDirPath(unitID int) string {
	return fmt.Sprintf("./internal/data/units/unit_%d", unitID)
}

func UnitFilesDirPath(unitID int) string {
	return filepath.Join(UnitDirPath(unitID), "files")
}

func UnitImagePath(unitID int) string {
	return filepath.Join(UnitDirPath(unitID), "image.png")
}

func (c CourseRepo) SaveUnit(unit domain.Unit) (domain.Unit, error) {
	var hasDescr = unit.Description != ""
	currUnit := database.Unit{
		CourseID: int64(unit.CourseID),
		Number:   int64(unit.Number),
		Sequence: int64(unit.SequenceNum),
		Name:     unit.Name,
		Description: sql.NullString{
			String: unit.Description,
			Valid:  hasDescr,
		},
	}
	if currUnit.Sequence == 0 {
		return domain.Unit{}, fmt.Errorf("currUnit sequence is 0")
	}
	currUnit, err := c.queries.SaveUnit(context.Background(), database.SaveUnitParams{
		Number:      currUnit.Number,
		Sequence:    currUnit.Sequence,
		Name:        currUnit.Name,
		Description: currUnit.Description,
		CourseID:    currUnit.CourseID,
	})
	if err != nil {
		return domain.Unit{}, fmt.Errorf("courseRepo.SaveUnit(): %s", err)
	}
	unit.ID = int(currUnit.ID)
	if unit.SequenceNum == 0 {
		return domain.Unit{}, fmt.Errorf("unit sequence is 0")
	}
	return unit, nil

}

func (cr CourseRepo) GetUnit(unitID int) (domain.Unit, error) {
	dbUnit, err := cr.queries.GetUnit(context.Background(), int64(unitID))
	if err != nil {
		return domain.Unit{}, err
	}
	unit := domain.Unit{
		ID:          int(dbUnit.ID),
		CourseID:    int(dbUnit.CourseID),
		Number:      int(dbUnit.Number),
		SequenceNum: int(dbUnit.Sequence),
		Name:        dbUnit.Name,
		Description: dbUnit.Description.String,
	}
	return unit, nil
}
func (cr CourseRepo) GetUnits(courseID int) ([]domain.Unit, error) {
	var units []domain.Unit
	dbUnits, err := cr.queries.GetUnits(context.Background(), int64(courseID))
	if err != nil {
		return nil, err
	}
	for _, dbUnit := range dbUnits {
		unit := domain.Unit{
			ID:          int(dbUnit.ID),
			CourseID:    int(dbUnit.CourseID),
			Number:      int(dbUnit.Number),
			SequenceNum: int(dbUnit.Sequence),
			Name:        dbUnit.Name,
			Description: dbUnit.Description.String,
		}
		units = append(units, unit)
	}
	return units, nil
}

func (c CourseRepo) UpdateUnit(u domain.Unit) error {
	err := c.queries.UpdateUnit(context.Background(), database.UpdateUnitParams{
		ID:       int64(u.ID),
		Number:   int64(u.Number),
		Sequence: int64(u.SequenceNum),
		Name:     u.Name,
		Description: sql.NullString{
			Valid:  u.Description != "",
			String: u.Description,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (lr CourseRepo) DeleteUnit(unit domain.Unit) error {
	err := lr.deleteUnitDir(unit.ID)
	if err != nil {
		return err
	}
	_, err = lr.queries.DeleteUnit(context.Background(), int64(unit.ID))
	return err
}
func (cr CourseRepo) deleteUnitDir(unitID int) error {
	path := UnitDirPath(unitID)
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil

}

package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
)

func (c CourseRepo) SaveUnit(unit domain.Unit) (*domain.Unit, error) {
	log.Println("SaveUnit(): ", "templateID", unit.TemplateID, "ID", unit.ID)
	var hasDescr = unit.Description != ""
	currUnit := database.Unit{
		CourseID: int64(unit.CourseID),
		TemplateID: sql.NullInt64{
			Int64: int64(unit.TemplateID),
			Valid: unit.TemplateID != 0,
		},
		Number:   int64(unit.Number),
		Sequence: int64(unit.SequenceNum),
		Name:     unit.Name,
		Description: sql.NullString{
			String: unit.Description,
			Valid:  hasDescr,
		},
	}
	if currUnit.Sequence == 0 {
		return nil, fmt.Errorf("currUnit sequence is 0")
	}
	currUnit, err := c.queries.SaveUnit(context.Background(), database.SaveUnitParams{
		Number:      currUnit.Number,
		Sequence:    currUnit.Sequence,
		TemplateID:  currUnit.TemplateID,
		Name:        currUnit.Name,
		Description: currUnit.Description,
		CourseID:    currUnit.CourseID,
	})
	if err != nil {
		return nil, fmt.Errorf("courseRepo.SaveUnit(): %s", err)
	}
	unit.ID = int(currUnit.ID)
	log.Println("unit sequence:", unit.SequenceNum)
	if unit.SequenceNum == 0 {
		return nil, fmt.Errorf("unit sequence is 0")
	}
	return &unit, nil

}

func (ur CourseRepo) GetUnits(courseID int) ([]domain.Unit, error) {
	var units []domain.Unit
	dbUnits, err := ur.queries.GetUnits(context.Background(), int64(courseID))
	if err != nil {
		return nil, err
	}
	for _, dbUnit := range dbUnits {
		unit := domain.Unit{
			ID:          int(dbUnit.ID),
			CourseID:    int(dbUnit.CourseID),
			TemplateID:  int(dbUnit.TemplateID.Int64),
			Number:      int(dbUnit.Number),
			SequenceNum: int(dbUnit.Sequence),
			Name:        dbUnit.Name,
			Description: dbUnit.Description.String,
		}
		units = append(units, unit)
	}
	return units, nil
}

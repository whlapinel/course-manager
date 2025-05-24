package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/core/unit"
	unitFeature "gh_static_portfolio/internal/features/unit"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
)

type unitRepo struct {
	queries *database.Queries
}

func NewUnitRepo(queries *database.Queries) unitFeature.Repository {
	return &unitRepo{
		queries: queries,
	}

}

func (u *unitRepo) convertFromDB(dbUnit database.Unit) unit.Unit {
	return unit.Unit{
		ID:          int(dbUnit.ID),
		CourseID:    int(dbUnit.CourseID),
		Number:      int(dbUnit.Number),
		SequenceNum: int(dbUnit.Sequence),
		Name:        dbUnit.Name,
		Description: dbUnit.Description.String,
	}
}

// ByCourseID implements unit.Repository.
func (u *unitRepo) ByCourseID(courseID int) ([]unit.Unit, error) {
	dbUnits, err := u.queries.GetUnits(context.Background(), int64(courseID))
	if err != nil {
		return nil, err
	}
	var units []unit.Unit
	for _, dbUnit := range dbUnits {
		unit := u.convertFromDB(dbUnit)
		units = append(units, unit)
	}
	return units, nil
}

// ByID implements unit.Repository.
func (u *unitRepo) ByID(unitID int) (unit.Unit, error) {
	dbUnit, err := u.queries.GetUnit(context.Background(), int64(unitID))
	if err != nil {
		return unit.Unit{}, err
	}
	return u.convertFromDB(dbUnit), nil
}

// Delete implements unit.Repository.
func (u *unitRepo) Delete(unitID int) error {
	_, err := u.queries.DeleteUnit(context.Background(), int64(unitID))
	if err != nil {
		return err
	}
	return nil
}

// Save implements unit.Repository.
func (u *unitRepo) Save(newUnit unit.Unit) (int, error) {
	unitParams := database.SaveUnitParams{
		CourseID: int64(newUnit.CourseID),
		Name:     newUnit.Name,

		Description: sql.NullString{
			Valid:  newUnit.Description != "",
			String: newUnit.Description,
		},
	}
	dbUnit, err := u.queries.SaveUnit(context.Background(), unitParams)
	if err != nil {
		return 0, fmt.Errorf("unitRepo.SaveUnit: %s", err)
	}
	return int(dbUnit.ID), nil

}

// Update implements unit.Repository.
func (u *unitRepo) Update(updated unit.Unit) error {
	err := u.queries.UpdateUnit(context.Background(), database.UpdateUnitParams{
		ID:   int64(updated.ID),
		Name: updated.Name,
		Description: sql.NullString{
			Valid:  updated.Description != "",
			String: updated.Description,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

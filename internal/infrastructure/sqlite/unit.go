package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/features/unit"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"gh_static_portfolio/internal/ports"
)

type unitRepo struct {
	queries *database.Queries
}

func NewUnitRepo(queries *database.Queries) unit.Repository {
	return &unitRepo{
		queries: queries,
	}

}

func (u *unitRepo) convertFromDB(dbUnit database.Unit) unit.Unit {
	return unit.Unit{
		BaseNode: ports.BaseNode[int, int]{

			ID:          int(dbUnit.ID),
			ParentID:    int(dbUnit.CourseID),
			Number:      int(dbUnit.Number),
			Sequence:    int(dbUnit.Sequence),
			Name:        dbUnit.Name,
			Description: dbUnit.Description.String,
		},
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
		Number:   int64(newUnit.Number),
		CourseID: int64(newUnit.ParentID),
		Name:     newUnit.Name,
		Sequence: int64(newUnit.Sequence),
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
		Number:   int64(updated.Number),
		Sequence: int64(updated.Sequence),
		ID:       int64(updated.ID),
		Name:     updated.Name,
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

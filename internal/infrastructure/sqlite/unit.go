package sqlite

import (
	"context"
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
	panic("unimplemented")
}

// Save implements unit.Repository.
func (u *unitRepo) Save(newUnit unit.Unit) (int, error) {
	panic("unimplemented")
}

// Update implements unit.Repository.
func (u *unitRepo) Update(updated unit.Unit) error {
	panic("unimplemented")
}

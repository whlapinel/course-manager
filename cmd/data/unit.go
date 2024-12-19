package data

import (
	"context"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
)

type UnitRepo struct {
	queries *database.Queries
}

func NewUnitRepo(queries *database.Queries) UnitRepo {
	return UnitRepo{queries: queries}
}

func (ur UnitRepo) All(courseID int) ([]domain.Unit, error) {
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

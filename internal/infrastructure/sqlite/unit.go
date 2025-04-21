package sqlite

import (
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

// ByCourseID implements unit.Repository.
func (u *unitRepo) ByCourseID(courseID int) ([]unit.Unit, error) {
	panic("unimplemented")
}

// ByID implements unit.Repository.
func (u *unitRepo) ByID(unitID int) (unit.Unit, error) {
	panic("unimplemented")
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

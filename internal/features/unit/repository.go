package unit

import (
	"gh_static_portfolio/internal/core/unit"
)

type Repository interface {
	ByCourseID(courseID int) ([]unit.Unit, error)
	ByID(unitID int) (unit.Unit, error)
	Save(newUnit unit.Unit) (int, error)
	Update(updated unit.Unit) error
	Delete(unitID int) error
}

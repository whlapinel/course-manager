package sqlite

import (
	"gh_static_portfolio/internal/core/course"
	courseFeature "gh_static_portfolio/internal/features/course"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
)

type courseRepo struct {
	queries *database.Queries
}

func NewCourseRepo(queries *database.Queries) courseFeature.Repository {
	return &courseRepo{
		queries: queries,
	}

}

// ByID implements course.Repository.
func (c *courseRepo) ByID(courseID int) (course.Course, error) {
	panic("unimplemented")
}

// ByTermID implements course.Repository.
func (c *courseRepo) ByTermID(termID int) ([]course.Course, error) {
	panic("unimplemented")
}

// Delete implements course.Repository.
func (c *courseRepo) Delete(courseID int) error {
	panic("unimplemented")
}

// Save implements course.Repository.
func (c *courseRepo) Save(newTerm course.Course) (int, error) {
	panic("unimplemented")
}

// Update implements course.Repository.
func (c *courseRepo) Update(updated course.Course) error {
	panic("unimplemented")
}

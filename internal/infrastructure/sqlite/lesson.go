package sqlite

import (
	"gh_static_portfolio/internal/core/lesson"
	lessonFeature "gh_static_portfolio/internal/features/lesson"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
)

type lessonRepo struct {
	queries *database.Queries
}

func NewLessonRepo(queries *database.Queries) lessonFeature.Repository {
	return &lessonRepo{
		queries: queries,
	}

}

// ByID implements lesson.Repository.
func (l *lessonRepo) ByID(lessonID int) (lesson.Lesson, error) {
	panic("unimplemented")
}

// ByUnitID implements lesson.Repository.
func (l *lessonRepo) ByUnitID(unitID int) ([]lesson.Lesson, error) {
	panic("unimplemented")
}

// Delete implements lesson.Repository.
func (l *lessonRepo) Delete(lessonID int) error {
	panic("unimplemented")
}

// Save implements lesson.Repository.
func (l *lessonRepo) Save(newLesson lesson.Lesson) (int, error) {
	panic("unimplemented")
}

// Update implements lesson.Repository.
func (l *lessonRepo) Update(updated lesson.Lesson) error {
	panic("unimplemented")
}

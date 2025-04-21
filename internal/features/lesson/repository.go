package lesson

import "gh_static_portfolio/internal/core/lesson"

type Repository interface {
	ByUnitID(unitID int) ([]lesson.Lesson, error)
	ByID(lessonID int) (lesson.Lesson, error)
	Save(newLesson lesson.Lesson) (int, error)
	Update(updated lesson.Lesson) error
	Delete(lessonID int) error
}

package course

import "gh_static_portfolio/internal/core/course"

type Repository interface {
	ByTermID(termID int) ([]course.Course, error)
	ByID(courseID int) (course.Course, error)
	Save(newTerm course.Course) (int, error)
	Update(updated course.Course) error
	Delete(courseID int) error
}

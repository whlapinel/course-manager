package course

type Repository interface {
	ByTermID(termID int) ([]Course, error)
	ByID(courseID int) (Course, error)
	Save(course Course) (int, error)
	Update(updated Course) error
	Delete(courseID int) error
}

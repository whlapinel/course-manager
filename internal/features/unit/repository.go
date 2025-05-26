package unit

type Repository interface {
	ByCourseID(courseID int) ([]Unit, error)
	ByID(unitID int) (Unit, error)
	Save(newUnit Unit) (int, error)
	Update(updated Unit) error
	Delete(unitID int) error
}

package occasion

type Repository interface {
	Save(Occasion) (int, error)
	Update(Occasion) error
	Delete(int) error
	ByID(int) (Occasion, error)
	ByParentID(int) ([]Occasion, error)
}

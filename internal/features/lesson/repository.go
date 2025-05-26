package lesson

type Repository interface {
	ByUnitID(unitID int) ([]Lesson, error)
	ByID(lessonID int) (Lesson, error)
	Save(newLesson Lesson) (int, error)
	Update(updated Lesson) error
	Delete(lessonID int) error
}

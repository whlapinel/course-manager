package domain

func NewUnit(num int, sequence int, name string, descr string, lessons []Lesson) Unit {
	return Unit{Number: num, SequenceNum: sequence, Name: name, Description: descr, Lessons: lessons}

}

func (u Unit) GetTitle() string {
	return u.Name
}

// Always associated with a particular course
type Unit struct {
	ID          int
	CourseID    int
	Number      int
	SequenceNum int
	Name        string
	Description string
	Lessons     []Lesson
}

func (u *Unit) AddLesson(lesson Lesson) {
	u.Lessons = append(u.Lessons, lesson)
}

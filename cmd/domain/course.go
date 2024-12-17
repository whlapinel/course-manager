package domain

func NewCourseTemplate(title string, descr string, units []Unit) Course {
	return Course{Name: title, Description: descr, Units: units}
}

type CourseRepo interface {
	ReadFromCSV() ([]*Course, error)
	GetTemplates() ([]*Course, error)
	GetInstances() ([]*CourseInstance, error)
	SaveTemplate(*Course) (*Course, error)
	SaveInstance(*CourseInstance) error
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type Course struct {
	ID          int
	Name        string
	Description string
	Units       []Unit
	TermID      int
	TermName    string
}

type CourseType int

type CourseInstance struct {
	Course
	TemplateID int
}

func (c Course) CreateInstance() *CourseInstance {
	var units []Unit
	for _, unit := range c.Units {
		unit.TemplateID = unit.ID
		unit.ID = 0
		var lessons []Lesson
		for _, lesson := range unit.Lessons {
			lesson.TemplateID = lesson.ID
			lesson.ID = 0
			lessons = append(lessons, lesson)
		}
		unit.Lessons = lessons
		units = append(units, unit)
	}
	c.Units = units
	return &CourseInstance{
		Course:     c,
		TemplateID: c.ID,
	}

}

func (c Course) GetTitle() string {
	return c.Name
}

func (c *Course) AddUnit(unit Unit) {
	c.Units = append(c.Units, unit)
}

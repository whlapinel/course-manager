package domain

func NewCourseTemplate(title string, descr string, units []Unit) Course {
	return Course{Name: title, Description: descr, Units: units}
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type Course struct {
	ID          int
	Name        string
	Description string
	Units       []Unit
	Term
}

type CourseType int

type Courses []Course

func (i Courses) Courses() []Course {
	var courses []Course
	for _, course := range i {
		courses = append(courses, course)
	}
	return courses
}

func (c Course) GetTitle() string {
	return c.Name
}

func (c *Course) AddUnit(unit Unit) {
	c.Units = append(c.Units, unit)
}

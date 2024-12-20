package domain

func NewCourseTemplate(title string, descr string, units []Unit) CourseTemplate {
	return CourseTemplate{Name: title, Description: descr, Units: units}
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type CourseTemplate struct {
	ID          int
	Name        string
	Description string
	Units       []Unit
}

type CourseType int

type CourseInstance struct {
	CourseTemplate
	TemplateID int
	Term
}

func (c CourseTemplate) CreateInstance(term Term) CourseInstance {
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
	return CourseInstance{
		CourseTemplate: c,
		TemplateID:     c.ID,
		Term:           term,
	}

}

func (c CourseTemplate) GetTitle() string {
	return c.Name
}

func (c *CourseTemplate) AddUnit(unit Unit) {
	c.Units = append(c.Units, unit)
}

package domain

import "time"

type CurriculumLayer interface {
	GetImage() Image
	SetImage(Image)
}

func NewCourseTemplate(title string, descr string, units []*Unit) Course {
	return Course{Name: title, Description: descr, Units: units}
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type Course struct {
	ID          int
	Name        string
	Description string
	Units       []*Unit
	Image       Image
	Term
}

func (c Course) GetImage() Image {
	return c.Image
}

func (c *Course) SetImage(image Image) {
	c.Image = image
}

type CourseType int

type Courses []*Course

// takes a course and fits all lessons to dates in order, one lesson per day
func (c Course) FitToTerm(term Term) Course {
	c.Term = term
	currDay := 0
	for i, unit := range c.Units {
		for j, lesson := range unit.Lessons {
			currDate := term.InstructionalDays[currDay]
			lesson.Dates = []time.Time{currDate}
			if currDay < len(term.InstructionalDays)-1 {
				currDay++
			} else {
				lesson.Dates = []time.Time{}
			}
			unit.Lessons[j] = lesson
		}
		c.Units[i] = unit
	}
	return c
}

func (c Course) GetTitle() string {
	return c.Name
}

func (c *Course) AddUnit(unit Unit) {
	c.Units = append(c.Units, &unit)
}

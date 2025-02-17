package domain

import "time"

type NewCourseParams struct {
	TermID      int
	Name        string
	Description string
}

func NewCourse(params NewCourseParams) Course {
	return Course{
		Term: Term{
			ID: params.TermID,
		},
		Name:        params.Name,
		Description: params.Description,
	}
}

type Course struct {
	ID          int
	Name        string
	Description string
	Units       []Unit
	StandardSet StandardSet
	Image       Image
	Term
}

type CourseType int

type Courses []Course

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

func (c Course) GetName() string {
	return c.Name
}

func (c Course) GetDescription() string {
	return c.Description
}

func (c Course) GetID() interface{} {
	return c.ID
}

func (c Course) GetParentID() int {
	return c.Term.ID
}

func (c Course) Children() []CourseNode {
	var nodes []CourseNode
	for _, u := range c.Units {
		nodes = append(nodes, u)
	}
	return nodes
}

func (c Course) TypeName() string {
	return CourseTypeName.String()
}

func (c Course) ParentTypeName() string {
	return TermTypeName.String()
}

func (c Course) ChildTypeName() string {
	return UnitTypeName.String()
}

func (c *Course) AddUnit(unit Unit) {
	c.Units = append(c.Units, unit)
}

func (c Course) Designation() string {
	return ""
}

func (c Course) GetNumber() int {
	return -1
}

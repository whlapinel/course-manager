package dto

import (
	"gh_static_portfolio/internal/core/standard"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/ports"
	"time"
)

type Course struct {
	course.Course        `json:"course"`
	Units                []Unit `json:"units"`
	standard.StandardSet `json:"standardSet"`
}

type CourseType int

type Courses []Course

// takes a course and fits all lessons to dates in order, one lesson per day
func (c Course) FitToTerm(term Term) Course {
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

func (c Course) GetChildren() []ports.Node {
	var nodes []ports.Node
	for _, u := range c.Units {
		nodes = append(nodes, u)
	}
	return nodes
}

func (c Course) GetTypeName() string {
	return CourseTypeName.String()
}

func (c Course) GetParentTypeName() string {
	return TermTypeName.String()
}

func (c Course) GetChildTypeName() string {
	return UnitTypeName.String()
}

func (c *Course) AddUnit(unit Unit) {
	c.Units = append(c.Units, unit)
}

func (c Course) GetNumber() int {
	return -1
}

package service

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
)

func (svc CourseService) Nodes(params domain.NodePath) (domain.Nodes, error) {
	var nodes domain.Nodes
	if params.UserID != "" {
		user, err := svc.GetUser(params.UserID)
		if err != nil {
			return nodes, err
		}
		nodes.User = user
	} else {
		return nodes, nil
	}
	if params.TermID != 0 {
		term, err := svc.GetTerm(params.TermID)
		if err != nil {
			return nodes, err
		}
		nodes.Term = term
	} else {
		return nodes, nil
	}

	if params.CourseID != 0 {
		course, err := svc.GetCourse(params.CourseID)
		if err != nil {
			return nodes, err
		}
		nodes.Course = course
	} else {
		return nodes, nil
	}
	if params.UnitID != 0 {
		unit, err := svc.GetUnit(params.UnitID)
		if err != nil {
			return nodes, err
		}
		nodes.Unit = unit
	} else {
		return nodes, nil
	}
	if params.LessonID != 0 {
		lesson, err := svc.GetLesson(params.LessonID)
		if err != nil {
			return nodes, err
		}
		nodes.Lesson = lesson
	}
	return nodes, nil
}

type ParamsLength int

const (
	User ParamsLength = iota + 1
	Term
	Course
	Unit
	Lesson
)

// the last node will have children
func (svc CourseService) NodesWithChildren(params domain.NodePath) (domain.Nodes, error) {
	nodes, err := svc.Nodes(params)
	if err != nil {
		return nodes, err
	}
	switch ParamsLength(len(params.ToSlice())) {
	case User:
		terms, err := svc.GetTerms(params.UserID)
		if err != nil {
			return nodes, err
		}
		nodes.User.Terms = terms
	case Term:
		courses, err := svc.GetCourses(params.TermID)
		if err != nil {
			return nodes, err
		}
		nodes.Term.Courses = courses
	case Course:
		units, err := svc.GetUnits(params.CourseID)
		if err != nil {
			return nodes, err
		}
		nodes.Course.Units = units
	case Unit:
		lessons, err := svc.GetLessons(params.UnitID)
		if err != nil {
			return nodes, err
		}
		nodes.Unit.Lessons = lessons
	case Lesson:
		return nodes, fmt.Errorf(", NodesWithChildren called with params of length %d. Lesson does not have children", Lesson)
	}
	return nodes, nil

}

package data

import (
	"gh_static_portfolio/internal/domain"
	"log"
	"testing"
)

func TestNodeDir(t *testing.T) {
	term := domain.Term{
		ID: 1,
	}
	course := domain.Course{
		ID: 1,
	}
	unit := domain.Unit{
		ID: 1,
	}
	lesson := domain.Lesson{
		ID: 1,
	}
	unit.Lessons = append(unit.Lessons, lesson)
	course.Units = append(course.Units, unit)
	term.Courses = append(term.Courses, course)
	log.Println(NodeDirPath(term, course, unit))
}

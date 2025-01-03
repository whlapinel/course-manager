package data

import (
	"log"
	"testing"
)

func TestGetLessonFiles(t *testing.T) {
	terms, err := cr.GetTerms()
	if err != nil {
		t.Error(err)
	}
	term := terms[0]
	log.Println(term.Name)
	courses, err := cr.GetCourses(term.ID)
	if err != nil {
		t.Error(err)
	}
	course := courses[1]
	log.Println(course.Name)

	units, err := cr.GetUnits(course.ID)
	if err != nil {
		t.Error(err)
	}
	unit := units[7]
	log.Println(unit.Name)
	lessons, err := cr.GetLessons(unit.ID)
	if err != nil {
		t.Error(err)
	}
	for _, lesson := range lessons {
		log.Println(lesson.Files.ID)
	}

}

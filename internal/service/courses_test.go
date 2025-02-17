package service

import (
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCopyCourse(t *testing.T) {
	course, err := svc.CopyCourseToTerm(1, 2)
	if err != nil {
		t.Error(err)
	}
	log.Println(course)
	for _, date := range course.InstructionalDays {
		log.Println("Date:", date)
	}
	for _, unit := range course.Units {
		log.Println(unit.Name)
	}

}

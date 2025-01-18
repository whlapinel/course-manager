package service

import (
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetCourses(t *testing.T) {
	terms, err := cr.GetTerms()
	if err != nil {
		t.Error()
	}
	for i, term := range terms {
		log.Println("Term ", i, ":", term.Name, "id:", term.ID)
		courses, err := cr.GetCourses(term.ID)
		if err != nil {
			t.Error()
		}
		for j, course := range courses {
			log.Println("Course", j, "Name:", course.Name, "Dates:")
			for _, date := range course.InstructionalDays {
				log.Println("Date: ", date.Format(time.DateOnly))
			}
		}

	}

}

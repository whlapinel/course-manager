package data

import (
	"log"
	"testing"
	"time"
)

func TestGetCourses(t *testing.T) {
	term, err := cr.GetTerm(time.Now())
	if err != nil {
		t.Errorf("error fetching term: %s", err)
	}
	instances, err := cr.GetCourses(term.ID)
	if err != nil {
		t.Errorf("error geting instances: %s", err)
	}
	for _, instance := range instances {
		log.Println(instance.Name)
		for _, unit := range instance.Units {
			log.Println(unit.Name)
			for _, lesson := range unit.Lessons {
				log.Println(lesson.Name, lesson.Dates)
			}
		}
	}
}



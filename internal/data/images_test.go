package data

import (
	"log"
	"testing"
)

func TestGetImages(t *testing.T) {
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
	log.Println("image: ", course.Image.ID)

}

package domain

import (
	"log"
	"testing"
	"time"
)

func TestCreateInstance(t *testing.T) {
	lessons := []Lesson{
		{ID: 1, UnitID: 1, TemplateID: 0, Number: 1, Name: "Lesson 1.1", Description: "Lesson 1.1 Description"},
		{ID: 2, UnitID: 1, TemplateID: 0, Number: 2, Name: "Lesson 1.2", Description: "Lesson 1.2 Description"},
		{ID: 3, UnitID: 1, TemplateID: 0, Number: 3, Name: "Lesson 1.3", Description: "Lesson 1.3 Description"},
	}
	units := []Unit{
		{ID: 1,
			CourseID:    1,
			TemplateID:  0,
			Number:      1,
			SequenceNum: 1,
			Name:        "Unit 1",
			Description: "Unit 1 Description",
			Lessons:     lessons,
		},
	}
	template := NewCourseTemplate("Test Course Name", "Test Course Description", units)
	term, err := NewTerm(time.Now(), time.Now().AddDate(0, 3, 0), []time.Time{}, Semester, 1, "Fall 2024")
	if err != nil {
		t.Error(err)
	}
	instance := template.CreateInstance(term)
	log.Println("instance.TemplateID:", instance.TemplateID, "instance.ID: ", instance.CourseTemplate.ID)
	for _, unit := range instance.Units {
		log.Println("unit.TemplateID:", unit.TemplateID, "unit.ID: ", unit.ID)
		for _, lesson := range unit.Lessons {
			log.Println("lesson.TemplateID:", lesson.TemplateID, "lesson.ID: ", lesson.ID)

		}
	}

}

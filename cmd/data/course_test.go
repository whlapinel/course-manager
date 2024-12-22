package data

import (
	"log"
	"testing"
	"time"
)

func TestImportCoursesFromCSV(t *testing.T) {
	instances, err := importInstancesFromCSV()
	if err != nil {
		t.Errorf("error importing courses: %s", err)
	}
	for _, instance := range instances {
		log.Println(instance.CourseTemplate.Name)
		log.Println("num units: ", len(instance.Units))
		for _, unit := range instance.Units {
			log.Println(unit.Name)
			log.Println("num lessons: ", len(unit.Lessons))
			for _, lesson := range unit.Lessons {
				for _, date := range lesson.Dates {

					log.Println(lesson.Name, lesson.Description, date.Format(time.DateOnly))
				}
			}
		}
	}

}

func TestSaveInstance(t *testing.T) {
	terms, err := cr.ImportTermsFromCSV()
	if err != nil {
		t.Errorf("TermsLoader(): %s", err)
	}
	for i, term := range terms {
		terms[i].ID, err = cr.SaveTerm(term)
		if err != nil {
			t.Errorf("tr.Save(): %s", err)
		}
	}
	templates, err := importTemplatesFromCSV()
	if err != nil {
		t.Errorf("importCoursesFromCSV: %s", err)
	}
	for _, template := range templates {
		savedTemplate, err := cr.SaveTemplate(template)
		if err != nil {
			t.Errorf("cr.SaveTemplate(): %s", err)
		}
		for _, term := range terms {

			instance := savedTemplate.CreateInstance(term)
			err = cr.SaveInstance(instance)
			if err != nil {
				t.Errorf("cr.SaveInstance(): %s", err)
			}
		}

		log.Println("course ID: ", template.ID)

	}
}

func TestGetTemplates(t *testing.T) {
	templates, err := cr.GetTemplates()
	if err != nil {
		t.Errorf("error getting templates: %s", err)
	}
	for _, template := range templates {
		log.Println(template.Name)
		for _, unit := range template.Units {
			log.Println(unit.Name)
			log.Println(unit.Description)
			for _, lesson := range unit.Lessons {
				log.Println(lesson.Name)
			}
		}
	}

}

func TestGetInstances(t *testing.T) {
	term, err := cr.GetTerm(time.Now())
	if err != nil {
		t.Errorf("error fetching term: %s", err)
	}
	instances, err := cr.GetInstances(term.ID)
	if err != nil {
		t.Errorf("error geting instances: %s", err)
	}
	for _, instance := range instances {
		log.Println(instance.CourseTemplate.Name)
		for _, unit := range instance.Units {
			log.Println(unit.Name)
			for _, lesson := range unit.Lessons {
				log.Println(lesson.Name, lesson.Dates)
			}
		}
	}
}

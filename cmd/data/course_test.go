package data

import (
	"log"
	"testing"
	"time"
)

func TestImportCoursesFromCSV(t *testing.T) {
	courses, err := importCoursesFromCSV()
	if err != nil {
		t.Errorf("error importing courses: %s", err)
	}
	for _, course := range courses {
		log.Println(course.Name)
		log.Println("num units: ", len(course.Units))
		for _, unit := range course.Units {
			log.Println(unit.Name)
			log.Println("num lessons: ", len(unit.Lessons))
			for _, lesson := range unit.Lessons {
				log.Println(lesson.Name, lesson.Description, lesson.Date.Format(time.DateOnly))
			}
		}
	}

}

func TestGenerateInstances(t *testing.T) {
	date := time.Now().AddDate(0, 2, 0)
	courseInstances, err := GenerateCourseInstancesFromCSV2(date)
	if err != nil {
		t.Errorf("error generating instances from csv: %s", err)
	}
	for _, instance := range courseInstances {
		log.Println(instance.Name)
		log.Println("num units: ", len(instance.Units))
		for _, unit := range instance.Units {
			log.Println(unit.Name)
			log.Println("num lessons: ", len(unit.Lessons))
			for _, lesson := range unit.Lessons {
				log.Println(lesson.Name, lesson.Description, lesson.Date.Format(time.DateOnly))
			}
		}
	}
	err = WriteCourseInstancesToCSV(courseInstances)
	if err != nil {
		t.Errorf("error writing to csv: %s", err)
	}
}

func TestSaveInstance(t *testing.T) {
	terms, err := TermsLoader()
	if err != nil {
		t.Errorf("TermsLoader(): %s", err)
	}
	for _, term := range terms {
		_, err := tr.Save(term)
		if err != nil {
			t.Errorf("tr.Save(): %s", err)
		}
	}
	templates, err := importCoursesFromCSV()
	if err != nil {
		t.Errorf("importCoursesFromCSV: %s", err)
	}
	for _, template := range templates {
		savedTemplate, err := cr.SaveTemplate(template)
		if err != nil {
			t.Errorf("cr.SaveTemplate(): %s", err)
		}
		instance := savedTemplate.CreateInstance()
		err = cr.SaveInstance(instance)
		if err != nil {
			t.Errorf("cr.SaveInstance(): %s", err)
		}
		log.Println("course ID: ", template.ID)

	}
}

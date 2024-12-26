package data

import (
	"gh_static_portfolio/cmd/data/csv"
	"log"
	"testing"
	"time"
)

func TestImportCoursesFromCSV(t *testing.T) {
	instances, err := csv.ImportCoursesFromCSV()
	if err != nil {
		t.Errorf("error importing courses: %s", err)
	}
	for _, instance := range instances {
		log.Println(instance.Name)
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
	courses, err := csv.ImportCoursesFromCSV()
	if err != nil {
		t.Errorf("importCoursesFromCSV: %s", err)
	}
	for _, course := range courses {
		for _, term := range terms {
			course.Term = term
			course.ID, err = cr.SaveCourse(course)
			if err != nil {
				t.Errorf("cr.SaveInstance(): %s", err)
			}
		}

		log.Println("course ID: ", course.ID)

	}
}

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

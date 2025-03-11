package newgensite

import (
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/service"
	"os"
	"testing"
)

func TestGenerator(t *testing.T) {
	err := os.Chdir("../../")
	if err != nil {
		t.Error(err)
	}
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	repo := data.NewCourseRepo(queries)
	svc := service.NewCourseService(repo)
	user, err := svc.GetUser("101602110272674353046")
	if err != nil {
		t.Error(err)
	}
	term, err := svc.GetTerm(2)
	if err != nil {
		t.Error(err)
	}
	courses, err := svc.GetCourses(2)
	if err != nil {
		t.Error(err)
	}
	for i, course := range courses {
		units, err := svc.GetUnits(course.ID)
		if err != nil {
			t.Error(err)
		}
		for j, unit := range units {
			lessons, err := svc.GetLessons(unit.ID)
			if err != nil {
				t.Error(err)
			}
			units[j].Lessons = lessons
		}
		courses[i].Units = units
	}
	term.Courses = courses
	err = Generate(user, term)
	if err != nil {
		t.Error(err)
	}

}

package service

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	managertemplates "gh_static_portfolio/internal/templates/manager_templates"
	"log"
	"os"
	"testing"
)

var cr data.CourseRepo
var svc CourseService

func TestMain(m *testing.M) {
	// copy db to this dir

	queries, db, err := data.InitDB("course_manager.db")
	cr = data.NewCourseRepo(queries)
	svc = NewCourseService(cr)
	if err != nil {
		fmt.Println("failed to connect to db: ", err)
		os.Exit(1)
	}
	defer db.Close()
	code := m.Run()
	os.Exit(code)
}

func TestCalendarDates(t *testing.T) {

	course, err := svc.GetCourseForCalendar(5)
	if err != nil {
		t.Error(err)
	}
	courseCal := managertemplates.CourseCalendar{
		Course: course,
	}
	calData := courseCal.CalendarDates()
	for date, data := range calData {
		log.Println(date, data)
	}

}

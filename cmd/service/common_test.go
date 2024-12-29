package service

import (
	"fmt"
	"gh_static_portfolio/cmd/data"
	"os"
	"testing"
)

var cr data.CourseRepo
var svc CourseService

func TestMain(m *testing.M) {
	queries, db, err := data.InitDB("test_course_manager.db")
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

package data

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"os"
	"testing"
)

var tr TermRepo
var cr domain.CourseRepo

func TestMain(m *testing.M) {
	queries, db, err := InitDB("test_course_manager.db")
	tr = NewTermRepo(queries)
	cr = NewCourseRepo(queries)
	if err != nil {
		fmt.Println("failed to connect to db: ", err)
		os.Exit(1)
	}
	defer db.Close()
	code := m.Run()
	os.Exit(code)
}

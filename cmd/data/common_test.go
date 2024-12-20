package data

import (
	"fmt"
	"os"
	"testing"
)

var cr CourseRepo

func TestMain(m *testing.M) {
	queries, db, err := InitDB("test_course_manager.db")
	cr = NewCourseRepo(queries)
	if err != nil {
		fmt.Println("failed to connect to db: ", err)
		os.Exit(1)
	}
	defer db.Close()
	code := m.Run()
	os.Exit(code)
}

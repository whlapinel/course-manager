package data

import (
	"fmt"
	"os"
	"testing"
)

var tr TermRepo
var cr CourseRepo
var dsr DailyScheduleRepo
var ur UnitRepo
var lr LessonRepo

func TestMain(m *testing.M) {
	queries, db, err := InitDB("test_course_manager.db")
	tr = NewTermRepo(queries)
	cr = NewCourseRepo(queries)
	dsr = NewDailyScheduleRepo(queries)
	ur = NewUnitRepo(queries)
	lr = NewLessonRepo(queries)
	if err != nil {
		fmt.Println("failed to connect to db: ", err)
		os.Exit(1)
	}
	defer db.Close()
	code := m.Run()
	os.Exit(code)
}

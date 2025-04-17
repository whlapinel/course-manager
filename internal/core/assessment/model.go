package assessment

import "time"

type Assessment struct {
	ID           int
	CourseID     int
	UnitID       int
	LessonID     int
	Name         string
	Instructions string
	File         string
	Category     AssessmentCategory
	DateAssigned time.Time
	DateDue      time.Time
	Dropped      bool
}

type AssessmentCategory string

type AssessmentSortParam int

const (
	Assigned AssessmentSortParam = iota
	Due
	Category
)

var AssessmentSortParams = []AssessmentSortParam{Assigned, Due, Category}

func (param AssessmentSortParam) String() string {
	return []string{
		"assigned", "due", "category",
	}[param]
}

const (
	Prepare  AssessmentCategory = "prepare"
	Rehearse AssessmentCategory = "rehearse"
	Perform  AssessmentCategory = "perform"
	Midterm  AssessmentCategory = "midterm"
	Final    AssessmentCategory = "final"
)

type AssessmentCategories []AssessmentCategory

var Categories = []AssessmentCategory{
	Prepare,
	Rehearse,
	Perform,
	Midterm,
	Final,
}

type Assessments []Assessment

type AssessmentAnalysis struct {
	CategoryStats
	Active  int
	Dropped int
}
type CategoryStats map[AssessmentCategory]int

package assessment

import "time"

type Assessment struct {
	ID           int                `json:"id"`
	CourseID     int                `json:"courseID"`
	UnitID       int                `json:"unitID"`
	LessonID     int                `json:"lessonID"`
	Name         string             `json:"name"`
	Instructions string             `json:"instructions"`
	File         string             `json:"file"`
	Category     AssessmentCategory `json:"category"`
	DateAssigned time.Time          `json:"dateAssigned"`
	DateDue      time.Time          `json:"dateDue"`
	Dropped      bool               `json:"dropped"`
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

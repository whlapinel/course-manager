package domain

import (
	"slices"
	"strings"
	"time"
)

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

func (assms Assessments) Analysis() AssessmentAnalysis {
	var analysis AssessmentAnalysis
	categoryStats := make(CategoryStats)
	for _, assm := range assms {
		if count, ok := categoryStats[assm.Category]; ok {
			categoryStats[assm.Category] = count + 1
		} else {
			categoryStats[assm.Category] = 1
		}
		if assm.Dropped {
			analysis.Dropped += 1
		} else {
			analysis.Active += 1
		}
	}
	analysis.CategoryStats = categoryStats
	return analysis
}

func (assms Assessments) Sort(param AssessmentSortParam) {
	switch param {
	case Category:
		assms.sortByCategory()
	case Assigned:
		assms.sortByAssigned()
	case Due:
		assms.sortByDue()
	}
}

func (assms Assessments) sortByDue() {
	slices.SortFunc(assms, func(a, b Assessment) int {
		if a.DateDue.Before(b.DateDue) {
			return 1
		}
		return -1
	})
}
func (assms Assessments) sortByAssigned() {
	slices.SortFunc(assms, func(a, b Assessment) int {
		if a.DateAssigned.Before(b.DateAssigned) {
			return 1
		}
		return -1
	})
}

func (assms Assessments) sortByCategory() {
	slices.SortFunc(assms, func(a, b Assessment) int {
		if strings.ToLower(string(a.Category)) > strings.ToLower(string(b.Category)) {
			return 1
		}
		return -1
	})
}

func ActiveFilter(active bool) func(Assessment) bool {
	if active {
		return func(a Assessment) bool { return !a.Dropped }
	}
	return func(a Assessment) bool { return a.Dropped }
}

func CategoryFilter(category AssessmentCategory) func(Assessment) bool {
	if category == "" {
		return func(a Assessment) bool { return true }
	}
	return func(a Assessment) bool {
		return a.Category == category
	}
}

func StartFilter(startDate time.Time) func(Assessment) bool {
	if startDate.IsZero() {
		return func(a Assessment) bool {
			return true
		}
	}
	return func(a Assessment) bool {
		return !a.DateAssigned.Before(startDate)
	}
}

func EndFilter(endDate time.Time) func(Assessment) bool {
	if endDate.IsZero() {
		return func(a Assessment) bool {
			return true
		}
	}
	return func(a Assessment) bool {
		return !a.DateAssigned.After(endDate)
	}
}

type AssessmentFilter struct {
	Category  AssessmentCategory
	Start     time.Time
	End       time.Time
	Active    bool
	SortParam AssessmentSortParam
}

func (assms Assessments) FilterAndSort(filter AssessmentFilter) Assessments {
	assms = filterAssessments(
		assms,
		StartFilter(filter.Start),
		EndFilter(filter.End),
		CategoryFilter(filter.Category),
		ActiveFilter(filter.Active),
	)
	assms.Sort(filter.SortParam)
	return assms
}

func filterAssessments(assessments []Assessment, keepFuncs ...func(Assessment) bool) []Assessment {
	var filtered []Assessment
outer:
	for _, assm := range assessments {
		for _, fn := range keepFuncs {
			if !fn(assm) {
				continue outer
			}
		}
		filtered = append(filtered, assm)
	}
	return filtered
}

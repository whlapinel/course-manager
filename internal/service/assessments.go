package service

import (
	"gh_static_portfolio/internal/domain"
	"log"
	"time"
)

type SaveAssessmentParams struct {
	Assessment domain.Assessment
}

func (svc CourseService) SaveAssessment(params SaveAssessmentParams) (domain.Assessment, error) {
	return svc.repo.SaveAssessment(params.Assessment)
}

func (svc CourseService) GetLessonAssessments(lessonID int) ([]domain.Assessment, error) {
	return svc.repo.GetLessonAssessments(lessonID)
}

func (svc CourseService) GetAssessment(id int) (domain.Assessment, error) {
	return svc.repo.GetAssessment(id)
}

func (svc CourseService) FilterAssessmentsByCategoryAndDate(cat domain.AssessmentCategory, courseID int, start, end time.Time) ([]domain.Assessment, error) {
	assessments, err := svc.GetAssessmentsByCategory(cat, courseID)
	if err != nil {
		return nil, err
	}
	var filtered []domain.Assessment
	// if start.IsZero() {
	// 	start = start.AddDate(-10, 0, 0)
	// }
	// if end.IsZero() {
	// 	end = end.AddDate(10, 0, 0)
	// }
	for _, assessment := range assessments {
		if !(assessment.DateAssigned.After(start) || domain.IsSameDate(assessment.DateAssigned, start)) {
			continue
		}
		if !(assessment.DateAssigned.Before(end) || domain.IsSameDate(assessment.DateAssigned, end)) {
			continue
		}
		filtered = append(filtered, assessment)
	}
	return filtered, nil
}

func (svc CourseService) GetAssessmentsByCategory(category domain.AssessmentCategory, courseID int) ([]domain.Assessment, error) {
	if category == "" {
		return svc.GetAllCourseAssessments(courseID)
	}
	return svc.repo.GetAssessmentsByCategory(category, courseID)
}

func (svc CourseService) GetAllCourseAssessments(courseID int) ([]domain.Assessment, error) {
	return svc.repo.GetAllCourseAssessments(courseID)
}

type UpdateAssessmentParams struct {
	domain.Assessment
}

func (svc CourseService) UpdateAssessment(params UpdateAssessmentParams) error {
	log.Println("Service: Assessment", params.Assessment)
	return svc.repo.UpdateAssessment(params.Assessment)
}

func (svc CourseService) DeleteAssessment(assessmentID int) error {
	return svc.repo.DeleteAssessment(assessmentID)
}

// func CategoryFilter(category domain.AssessmentCategory) func(domain.Assessment) bool {
// 	if category == "" {
// 		return func(a domain.Assessment) bool { return true }
// 	}
// 	return func(a domain.Assessment) bool {
// 		return a.Category == category
// 	}
// }

// func StartFilter(startDate time.Time) func(domain.Assessment) bool {
// 	if startDate.IsZero() {
// 		return func(a domain.Assessment) bool {
// 			return true
// 		}
// 	}
// 	return func(a domain.Assessment) bool {
// 		return !a.DateAssigned.Before(startDate)
// 	}
// }

// func EndFilter(endDate time.Time) func(domain.Assessment) bool {
// 	if endDate.IsZero() {
// 		return func(a domain.Assessment) bool {
// 			return true
// 		}
// 	}
// 	return func(a domain.Assessment) bool {
// 		return !a.DateAssigned.After(endDate)
// 	}
// }

// type AssessmentFilter struct {
// 	Category  domain.AssessmentCategory
// 	Start     time.Time
// 	End       time.Time
// 	SortParam domain.AssessmentSortParam
// }

// func (h CourseService) FilterAssessments(assessments []domain.Assessment, filter AssessmentFilter) []domain.Assessment {
// 	return h.filterAssessments(
// 		assessments,
// 		StartFilter(filter.Start),
// 		EndFilter(filter.End),
// 		CategoryFilter(filter.Category),
// 	)
// }

// func (h CourseService) filterAssessments(assessments []domain.Assessment, keepFuncs ...func(domain.Assessment) bool) []domain.Assessment {
// 	var filtered []domain.Assessment
// outer:
// 	for _, assm := range assessments {
// 		for _, fn := range keepFuncs {
// 			if !fn(assm) {
// 				continue outer
// 			}
// 		}
// 		filtered = append(filtered, assm)
// 	}
// 	return filtered
// }

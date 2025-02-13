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
	for _, assessment := range assessments {
		if (assessment.DateAssigned.After(start) || domain.IsSameDate(assessment.DateAssigned, start)) &&
			(assessment.DateAssigned.Before(end) || domain.IsSameDate(assessment.DateAssigned, end)) {
			filtered = append(filtered, assessment)
		}
	}
	return filtered, nil
}

func (svc CourseService) GetAssessmentsByCategory(category domain.AssessmentCategory, courseID int) ([]domain.Assessment, error) {
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

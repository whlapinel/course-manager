package data

import (
	"context"
	"database/sql"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"time"
)

func (cr CourseRepo) SaveAssessment(assessment domain.Assessment) (domain.Assessment, error) {
	var dropped int64
	if assessment.Dropped {
		dropped = 1
	} else {
		dropped = 0
	}
	dbAssessment, err := cr.queries.SaveAssessment(context.Background(), database.SaveAssessmentParams{
		LessonID:     int64(assessment.LessonID),
		Name:         assessment.Name,
		Instructions: assessment.Instructions,
		File: sql.NullString{
			Valid:  assessment.File != "",
			String: assessment.File,
		},
		Category:     string(assessment.Category),
		DateAssigned: assessment.DateAssigned.Format(time.DateOnly),
		DateDue:      assessment.DateDue.Format(time.DateOnly),
		Dropped:      dropped,
	})
	if err != nil {
		return domain.Assessment{}, err
	}
	return cr.fromDBAssessment(dbAssessment)
}

func (cr CourseRepo) GetAssessment(id int) (domain.Assessment, error) {
	dbAssessment, err := cr.queries.GetAssessmentByID(context.Background(), int64(id))
	if err != nil {
		return domain.Assessment{}, err
	}
	assessment, err := cr.fromDBAssessment(dbAssessment)
	if err != nil {
		return domain.Assessment{}, err
	}
	return assessment, nil
}

func (cr CourseRepo) GetLessonAssessments(lessonID int) ([]domain.Assessment, error) {
	dbAssessments, err := cr.queries.GetAssessmentsByLessonID(context.Background(), int64(lessonID))
	if err != nil {
		return nil, err
	}
	var assessments []domain.Assessment
	for _, dbAssessment := range dbAssessments {
		assessment, err := cr.fromDBAssessment(dbAssessment)
		if err != nil {
			return nil, err
		}
		assessments = append(assessments, assessment)
	}
	return assessments, nil
}

func (cr CourseRepo) fromDBAssessment(dbAssmt database.Assessment) (domain.Assessment, error) {
	assigned, err := time.Parse(time.DateOnly, dbAssmt.DateAssigned)
	if err != nil {
		return domain.Assessment{}, err
	}
	due, err := time.Parse(time.DateOnly, dbAssmt.DateDue)
	if err != nil {
		return domain.Assessment{}, err
	}
	var dropped bool
	if dbAssmt.Dropped == 0 {
		dropped = false
	} else {
		dropped = true
	}
	return domain.Assessment{
		ID:           int(dbAssmt.ID),
		LessonID:     int(dbAssmt.LessonID),
		Name:         dbAssmt.Name,
		Instructions: dbAssmt.Instructions,
		File:         dbAssmt.File.String,
		Category:     domain.AssessmentCategory(dbAssmt.Category),
		DateAssigned: assigned,
		DateDue:      due,
		Dropped:      dropped,
	}, nil

}

func (cr CourseRepo) GetAssessmentsByCategory(category domain.AssessmentCategory, courseID int) ([]domain.Assessment, error) {
	dbAssessments, err := cr.queries.GetAssessmentsByCategory(context.Background(), database.GetAssessmentsByCategoryParams{
		ID:       int64(courseID),
		Category: string(category),
	})
	if err != nil {
		return nil, err
	}
	var assessments []domain.Assessment
	for _, dbAssessment := range dbAssessments {
		assessment, err := cr.fromDBAssessment(dbAssessment)
		if err != nil {
			return nil, err
		}
		assessments = append(assessments, assessment)
	}
	return assessments, nil
}

func (cr CourseRepo) GetAllCourseAssessments(courseID int) ([]domain.Assessment, error) {
	dbAssessments, err := cr.queries.GetCourseAssessments(context.Background(), int64(courseID))
	if err != nil {
		return nil, err
	}
	var assessments []domain.Assessment
	for _, dbassessment := range dbAssessments {
		assessment, err := cr.fromDBAssessment(dbassessment)
		if err != nil {
			return nil, err
		}
		assessments = append(assessments, assessment)
	}
	return assessments, nil
}

func (cr CourseRepo) DeleteAssessment(assessmentID int) error {
	return cr.queries.DeleteAssessment(context.Background(), int64(assessmentID))
}

func (cr CourseRepo) UpdateAssessment(assessment domain.Assessment) error {
	var dropped int
	if assessment.Dropped {
		dropped = 1
	} else {
		dropped = 0
	}
	err := cr.queries.UpdateAssessment(context.Background(), database.UpdateAssessmentParams{
		ID:           int64(assessment.ID),
		LessonID:     int64(assessment.LessonID),
		Name:         assessment.Name,
		Instructions: assessment.Instructions,
		File: sql.NullString{
			Valid:  assessment.File != "",
			String: assessment.File,
		},
		Category:     string(assessment.Category),
		DateAssigned: assessment.DateAssigned.Format(time.DateOnly),
		DateDue:      assessment.DateDue.Format(time.DateOnly),
		Dropped:      int64(dropped),
	})
	if err != nil {
		return err
	}
	return nil
}

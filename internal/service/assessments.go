package service

import "gh_static_portfolio/internal/domain"

type SaveAssessmentParams struct {
	Assessment domain.Assessment
}

func (svc CourseService) SaveAssessment(params SaveAssessmentParams) (domain.Assessment, error) {
	return svc.repo.SaveAssessment(params.Assessment)
}

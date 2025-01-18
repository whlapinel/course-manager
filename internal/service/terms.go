package service

import (
	"gh_static_portfolio/internal/domain"
)

func (svc CourseService) SaveTerm(term domain.Term) (int, error) {
	id, err := svc.repo.SaveTerm(term)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc CourseService) UpdateTerm(term domain.Term) error {
	return svc.repo.UpdateTerm(term)
}

func (svc CourseService) DeleteTerm(termID int) error {
	return svc.repo.DeleteTerm(termID)
}

func (svc CourseService) GetTerm(termID int) (domain.Term, error) {
	term, err := svc.repo.GetTermWithDates(termID)
	if err != nil {
		return term, err
	}
	return term, nil
}

func (svc CourseService) GetTerms() ([]domain.Term, error) {
	terms, err := svc.repo.GetTerms()
	if err != nil {
		return nil, err
	}
	return terms, nil
}

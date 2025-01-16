package service

import (
	"gh_static_portfolio/cmd/domain"
	"log"
)

func (svc CourseService) GetTerm(termID int) (domain.Term, error) {
	term, err := svc.repo.GetTermByID(termID)
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
func (svc CourseService) GetTermDates(termID int) (domain.Term, error) {
	var term domain.Term
	term, err := svc.repo.GetTermDates(termID)
	if err != nil {
		return term, err
	}
	if len(term.InstructionalDays) == 0 {
		log.Println("warning: term instructional days was 0. fyne_app courseService GetTermDates")
	}
	return term, nil
}

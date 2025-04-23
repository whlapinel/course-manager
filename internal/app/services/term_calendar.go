package services

import (
	"gh_static_portfolio/internal/app/dto"
)

type TermCalendarService struct {
	term func(termID int) (dto.Term, error)
}

func NewTermCalendarService(term func(termID int) (dto.Term, error)) *TermCalendarService {
	return &TermCalendarService{
		term: term,
	}
}

func (svc *TermCalendarService) Term(termID int) (dto.Term, error) {
	termDTO, err := svc.term(termID)
	if err != nil {
		return dto.Term{}, err
	}
	return termDTO, nil
}

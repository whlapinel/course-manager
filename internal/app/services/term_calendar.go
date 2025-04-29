package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/termoccasion"
)

type TermCalendarService struct {
	getTerm         func(termID int) (dto.Term, error)
	occasionService *termoccasion.Service
}

func NewTermCalendarService(
	getTerm func(termID int) (dto.Term, error),
	occasionService *termoccasion.Service,
) *TermCalendarService {
	return &TermCalendarService{
		getTerm:         getTerm,
		occasionService: occasionService,
	}
}

func (svc *TermCalendarService) Term(termID int) (dto.Term, error) {
	termDTO, err := svc.getTerm(termID)
	if err != nil {
		return dto.Term{}, err
	}
	return termDTO, nil
}

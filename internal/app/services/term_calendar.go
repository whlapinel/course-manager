package services

import (
	"gh_static_portfolio/internal/app/dto"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
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

func (svc *TermCalendarService) CalendarDates(termID int) (calendarviews.CalendarDates, error) {
	dates := make(calendarviews.CalendarDates)
	occasions, err := svc.occasionService.ByParentID(termID)
	if err != nil {
		return nil, err
	}
	for _, occ := range occasions {
		dates[occ.Date] = calendarviews.CalendarDate{
			TermOccasions: append(dates[occ.Date].TermOccasions, occ),
		}
	}
	return dates, err
}

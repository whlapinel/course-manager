package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/features/termoccasion"
	"gh_static_portfolio/internal/app/views/calendar"
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

func (svc *TermCalendarService) Occasion(occasionID int) (occasion.Occasion, error) {
	return svc.occasionService.ByID(occasionID)
}

func (svc *TermCalendarService) CalendarDates(termID int) (calendarviews.CalendarDates, error) {
	dates := make(calendarviews.CalendarDates)
	occasions, err := svc.occasionService.ByTermID(termID)
	if err != nil {
		return nil, err
	}
	for _, occ := range occasions {
		dates[occ.Date] = calendarviews.CalendarDate{
			Occasions: append(dates[occ.Date].Occasions, occ),
		}
	}
	return dates, err
}

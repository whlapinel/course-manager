package service

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"time"
)

type SaveTermParams struct {
	Name        string
	Description string
	Start       time.Time
	End         time.Time
}

func (svc CourseService) SaveTerm(svcTerm SaveTermParams) (int, error) {
	term, err := domain.NewTerm(domain.NewTermParams{
		Name:        svcTerm.Name,
		Description: svcTerm.Description,
		Start:       svcTerm.Start,
		End:         svcTerm.End,
	})
	if err != nil {
		return 0, err
	}
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

// since the non-instruct days are extrapolated rather than explicitly defined, we can simply delete
// the date from the column. lesson dates will be deleted using on cascade
func (svc CourseService) AddNonInstructDay(termID int, date time.Time) error {
	// shift all lessons for courses in term right or left
	err := svc.ShiftLessonsOnDate(termID, date)
	if err != nil {
		return err
	}
	err = svc.repo.DeleteDate(termID, date)
	if err != nil {
		return err
	}
	return nil
}

// this will move all dates to the left for a given date in preparation for deleting the date
func (svc CourseService) ShiftLessonsOnDate(termID int, date time.Time) error {
	term, err := svc.repo.GetTermWithDates(termID)
	if err != nil {
		return err
	}
	var newDate time.Time
	instructDays := term.InstructionalDays
	if len(instructDays) == 0 {
		return fmt.Errorf("length of term instructional days is 0")
	}
	for i, instructDate := range instructDays {
		if instructDate == date {
			if i != len(instructDays)-1 {
				newDate = instructDays[i+1]
			} else {
				newDate = instructDays[i-1]
			}
		}
	}
	lessons, err := svc.repo.GetLessonsOnDate(date, termID)
	if err != nil {
		return err
	}
	// important! dates were not fetched so the dates field should be empty.
	// Thus we do not need to remove the dates before saving. Any lessons added
	// to that junction table should get deleted using on cascade anyway.
	for _, lesson := range lessons {
		lesson.Dates = append(lesson.Dates, newDate)
		svc.repo.SaveLessonDate(lesson)
	}
	return nil
}

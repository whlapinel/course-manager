package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"log"
	"time"
)

func (cr CourseRepo) GetTerm(date time.Time, userID string) (*domain.Term, error) {
	dbTerms, err := cr.queries.GetTerms(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	for _, dbTerm := range dbTerms {
		start, err := time.Parse(time.DateOnly, dbTerm.Start)
		if err != nil {
			return nil, err
		}
		end, err := time.Parse(time.DateOnly, dbTerm.End)
		if err != nil {
			return nil, err
		}
		currentTerm := &domain.Term{
			ID:          int(dbTerm.ID),
			Name:        dbTerm.Name,
			Description: dbTerm.Description.String,
			Start:       start,
			End:         end,
		}
		if currentTerm.Start.Before(date) && currentTerm.End.After(date) {
			return currentTerm, nil
		}
	}
	return nil, nil
}

func (cr CourseRepo) GetTermByID(termID int) (domain.Term, error) {
	dbTerm, err := cr.queries.GetTermByID(context.Background(), int64(termID))
	if err != nil {
		return domain.Term{}, err
	}
	dates, err := parseDates([]string{dbTerm.Start, dbTerm.End})
	if err != nil {
		return domain.Term{}, err
	}
	return domain.Term{
		ID:          int(dbTerm.ID),
		Name:        dbTerm.Name,
		Description: dbTerm.Description.String,
		Start:       dates[0],
		End:         dates[1],
	}, nil

}

func (cr CourseRepo) GetTerms(userID string) ([]domain.Term, error) {
	dbGetTermRows, err := cr.queries.GetTerms(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	var terms []domain.Term
	var term domain.Term
	for i, dbGetTermRow := range dbGetTermRows {
		parsedStart, err := time.Parse(time.DateOnly, dbGetTermRow.Start)
		if err != nil {
			return nil, err
		}
		parsedEnd, err := time.Parse(time.DateOnly, dbGetTermRow.End)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			term = domain.Term{
				ID:          int(dbGetTermRow.ID),
				Start:       parsedStart,
				End:         parsedEnd,
				Name:        dbGetTermRow.Name,
				Description: dbGetTermRow.Description.String,
			}
		}
		// if we've hit a new term, append the current term and create a new one
		if dbGetTermRow.ID != int64(term.ID) {
			log.Println("new term encountered: appending current term and creating new.", dbGetTermRow.Name, term.Name)
			terms = append(terms, term)
			term = domain.Term{
				ID:          int(dbGetTermRow.ID),
				Start:       parsedStart,
				End:         parsedEnd,
				Name:        dbGetTermRow.Name,
				Description: dbGetTermRow.Description.String,
			}
		}
		if dbGetTermRow.Date.Valid {
			parsedInstructDate, err := time.Parse(time.DateOnly, dbGetTermRow.Date.String)
			if err != nil {
				return nil, err
			}
			term.InstructionalDays = append(term.InstructionalDays, parsedInstructDate)
		}
	}
	terms = append(terms, term)
	return terms, nil

}

func (cr CourseRepo) GetTermWithDates(termID int) (domain.Term, error) {
	var term domain.Term
	term, err := cr.GetTermByID(termID)
	if err != nil {
		return term, err
	}
	dbDates, err := cr.queries.GetTermDates(context.Background(), int64(termID))
	if len(dbDates) == 0 {
		log.Println("dates returned: 0. CourseRepo.GetTermDates")
	}
	if err != nil {
		return term, nil
	}
	currDate := term.Start
	currInstructIndex := 0
	for !currDate.After(term.End) {
		if currDate.Weekday() != time.Saturday && currDate.Weekday() != time.Sunday &&
			!(currInstructIndex > len(dbDates)-1) {
			dbDate := dbDates[currInstructIndex]
			date, err := time.Parse(time.DateOnly, dbDate.Date)
			if err != nil {
				return term, err
			}
			if currDate == date {
				term.InstructionalDays = append(term.InstructionalDays, currDate)
				currInstructIndex++
			} else {
				term.NonInstructionalDays = append(term.NonInstructionalDays, currDate)
			}
		}
		currDate = currDate.AddDate(0, 0, 1)
	}

	return term, nil
}

func (cr CourseRepo) GetTermOccasions(termID int) ([]domain.Occasion, error) {
	dbOccasions, err := cr.queries.GetTermOccasions(context.Background(), int64(termID))
	if err != nil {
		return nil, err
	}
	var occasions []domain.Occasion
	for _, dbOccasion := range dbOccasions {
		date, err := time.Parse(time.DateOnly, dbOccasion.Date)
		if err != nil {
			return nil, err
		}
		occasion := domain.Occasion{
			ID:     int(dbOccasion.ID),
			TermID: termID,
			Date:   date,
			Name:   dbOccasion.Name,
		}
		occasions = append(occasions, occasion)

	}
	return occasions, nil

}

func (cr CourseRepo) GetOccasionByID(occasionID int) (domain.Occasion, error) {
	dbOccasion, err := cr.queries.GetOccasionByID(context.Background(), int64(occasionID))
	if err != nil {
		return domain.Occasion{}, err
	}
	date, err := time.Parse(time.DateOnly, dbOccasion.Date)
	if err != nil {
		return domain.Occasion{}, err
	}
	return domain.Occasion{
		ID:     int(dbOccasion.ID),
		TermID: int(dbOccasion.TermID),
		Date:   date,
		Name:   dbOccasion.Name,
	}, nil
}

func (cr CourseRepo) UpdateOccasion(occasion domain.Occasion) error {
	return cr.queries.UpdateOccasion(context.Background(), database.UpdateOccasionParams{
		ID:   int64(occasion.ID),
		Name: occasion.Name,
	})

}

func (cr CourseRepo) DeleteOccasion(occasion domain.Occasion) error {
	return cr.queries.DeleteOccasion(context.Background(), int64(occasion.ID))
}

func (t CourseRepo) SaveTerm(term domain.Term) (int, error) {
	termParams := database.SaveTermParams{
		Name: term.Name,
		Description: sql.NullString{
			Valid:  term.Description != "",
			String: term.Description,
		},
		Start: term.Start.Format(time.DateOnly),
		End:   term.End.Format(time.DateOnly),
	}
	dbTerm, err := t.queries.SaveTerm(context.Background(), termParams)
	if err != nil {
		return 0, fmt.Errorf("CourseRepo.SaveTerm: %s", err)
	}
	for _, date := range term.InstructionalDays {
		dateParams := database.SaveDateParams{
			TermID: dbTerm.ID,
			Date:   date.Format(time.DateOnly),
		}
		t.queries.SaveDate(context.Background(), dateParams)
	}
	return int(dbTerm.ID), nil
}

func (r CourseRepo) UpdateTerm(term domain.Term) error {
	err := r.queries.UpdateTerm(context.Background(), database.UpdateTermParams{
		ID:   int64(term.ID),
		Name: term.Name,
		Description: sql.NullString{
			Valid:  term.Description != "",
			String: term.Description,
		},
		Start: term.Start.Format(time.DateOnly),
		End:   term.End.Format(time.DateOnly),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c CourseRepo) DeleteTerm(termID int) error {
	_, err := c.queries.DeleteTerm(context.Background(), int64(termID))
	if err != nil {
		return err
	}
	return nil
}

// func SaveTerm(ctx context.Context, term domain.Term, queries *database.Queries) (database.Term, error) {
// 	params := database.SaveTermParams{}
// 	params.Start = term.Start.Format(time.DateOnly)
// 	params.End = term.End.Format(time.DateOnly)
// 	dbTerm, err := queries.SaveTerm(ctx, params)
// 	if err != nil {
// 		return database.Term{}, err
// 	}
// 	return dbTerm, nil
// }

func parseDates(dateStrings []string) ([]time.Time, error) {
	var dates []time.Time
	for _, dateString := range dateStrings {
		date, err := time.Parse(time.DateOnly, dateString)
		if err != nil {
			return nil, err
		}
		dates = append(dates, date)

	}
	return dates, nil
}

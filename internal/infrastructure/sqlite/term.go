package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	core "gh_static_portfolio/internal/core/term"
	termfeature "gh_static_portfolio/internal/features/term"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"time"
)

type termRepo struct {
	queries *database.Queries
}

func NewTermRepo(queries *database.Queries) termfeature.Repository {
	return &termRepo{
		queries: queries,
	}

}

func (repo *termRepo) ByID(termID int) (core.Term, error) {
	var term core.Term
	dbTerm, err := repo.queries.GetTermByID(context.Background(), int64(termID))
	if err != nil {
		return core.Term{}, err
	}
	dates, err := parseDates([]string{dbTerm.Start, dbTerm.End})
	if err != nil {
		return core.Term{}, err
	}
	term = core.Term{
		ID:          int(dbTerm.ID),
		Name:        dbTerm.Name,
		Description: dbTerm.Description.String,
		Start:       dates[0],
		End:         dates[1],
	}
	dbDates, err := repo.queries.GetTermDates(context.Background(), int64(termID))
	if err != nil {
		return core.Term{}, err
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

func (repo *termRepo) ByUserID(userID string) ([]core.Term, error) {
	dbGetTermRows, err := repo.queries.GetTerms(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	var terms []core.Term
	var currTerm core.Term
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
			currTerm = core.Term{
				ID:          int(dbGetTermRow.ID),
				Start:       parsedStart,
				End:         parsedEnd,
				Name:        dbGetTermRow.Name,
				Description: dbGetTermRow.Description.String,
			}
		}
		// if we've hit a new term, append the current term and create a new one
		if dbGetTermRow.ID != int64(currTerm.ID) {
			terms = append(terms, currTerm)
			currTerm = core.Term{
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
			currTerm.InstructionalDays = append(currTerm.InstructionalDays, parsedInstructDate)
		}
	}
	if currTerm.ID != 0 {
		terms = append(terms, currTerm)
	}
	return terms, nil

}

func (repo *termRepo) Save(term core.Term) (int, error) {
	termParams := database.SaveTermParams{
		UserID: term.UserID,
		Name:   term.Name,

		Description: sql.NullString{
			Valid:  term.Description != "",
			String: term.Description,
		},
		Start: term.Start.Format(time.DateOnly),
		End:   term.End.Format(time.DateOnly),
	}
	dbTerm, err := repo.queries.SaveTerm(context.Background(), termParams)
	if err != nil {
		
		return 0, fmt.Errorf("termRepo.SaveTerm: %s", err)
	}
	for _, date := range term.InstructionalDays {
		dateParams := database.SaveDateParams{
			TermID: dbTerm.ID,
			Date:   date.Format(time.DateOnly),
		}
		repo.queries.SaveDate(context.Background(), dateParams)
	}
	return int(dbTerm.ID), nil
}

func (repo *termRepo) Update(term core.Term) error {
	err := repo.queries.UpdateTerm(context.Background(), database.UpdateTermParams{
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

func (repo *termRepo) Delete(termID int) error {
	_, err := repo.queries.DeleteTerm(context.Background(), int64(termID))
	if err != nil {
		return err
	}
	return nil
}

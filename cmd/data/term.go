package data

import (
	"context"
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"os"
	"strconv"
	"time"
)

type TermRepo interface {
	Save(term *domain.Term) (int, error)
	ReadFromCSV() ([]*domain.Term, error)
}

type termRepo struct {
	queries *database.Queries
}

// All implements TermRepo.
func (t termRepo) All() ([]*domain.Term, error) {
	panic("unimplemented")
}

// ReadFromCSV implements TermRepo.
func (t termRepo) ReadFromCSV() ([]*domain.Term, error) {
	terms, err := TermsLoader()
	if err != nil {
		return nil, err
	}
	return terms, nil
}

// Save implements TermRepo.
func (t termRepo) Save(term *domain.Term) (int, error) {
	termParams := database.SaveTermParams{
		Name:  term.Name,
		Start: term.Start.Format(time.DateOnly),
		End:   term.End.Format(time.DateOnly),
	}
	dbTerm, err := t.queries.SaveTerm(context.Background(), termParams)
	if err != nil {
		return 0, fmt.Errorf("termRepo.Save(): %s", err)
	}
	for i, date := range term.InstructionalDays {
		dateParams := database.SaveDateParams{
			TermID:    dbTerm.ID,
			DayNumber: int64(i) + 1,
			Date:      date.Format(time.DateOnly),
		}
		t.queries.SaveDate(context.Background(), dateParams)
	}
	return int(dbTerm.ID), nil
}

// WriteToCSV implements TermRepo.
func (t termRepo) WriteToCSV(*domain.Term) error {
	panic("unimplemented")
}

func NewTermRepo(queries *database.Queries) TermRepo {
	return termRepo{queries: queries}
}

const termsPath = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/terms.csv"
const nonIDaysPath = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/non_instruct_days.csv"

const (
	termIDCol = iota
	termStartCol
	termEndCol
	termTypeCol
	termNameCol
)

const (
	nonInstructionalDateCol = 1
)

func nonInstructionalDaysLoader() (*domain.NonInstructionalDays, error) {
	dates := domain.NonInstructionalDays{}
	file, err := os.Open(nonIDaysPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for i, record := range records {
		if i == 0 {
			continue
		}
		termID, err := strconv.Atoi(record[termIDCol])
		if err != nil {
			return nil, err
		}
		dates.TermID = append(dates.TermID, termID)
		date, err := time.Parse(time.DateOnly, record[nonInstructionalDateCol])
		if err != nil {
			return nil, err
		}
		dates.Dates = append(dates.Dates, date)
	}
	return &dates, nil
}

func filterNonInstructionalDates(termID int, dates *domain.NonInstructionalDays) []time.Time {
	filtered := []time.Time{}
	for i, date := range dates.Dates {
		if dates.TermID[i] == termID {
			filtered = append(filtered, date)
		}
	}
	return filtered
}

func TermsLoader() ([]*domain.Term, error) {
	file, err := os.Open(termsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	dates, err := nonInstructionalDaysLoader()
	if err != nil {
		return nil, err
	}
	terms := []*domain.Term{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		startDate, err := time.Parse(time.DateOnly, record[termStartCol])
		if err != nil {
			return nil, err
		}
		endDate, err := time.Parse(time.DateOnly, record[termEndCol])
		if err != nil {
			return nil, err
		}
		termID, err := strconv.Atoi(record[termIDCol])
		if err != nil {
			return nil, err
		}
		termDates := filterNonInstructionalDates(termID, dates)
		termType := record[termTypeCol]

		termName := record[termNameCol]
		term, err := domain.NewTerm(startDate, endDate, termDates, domain.TermType(termType), termID, termName)
		if err != nil {
			return nil, err
		}
		terms = append(terms, term)

	}
	return terms, nil

}

func SaveTerm(ctx context.Context, term domain.Term, queries *database.Queries) (database.Term, error) {
	params := database.SaveTermParams{}
	params.Start = term.Start.Format(time.DateOnly)
	params.End = term.End.Format(time.DateOnly)
	dbTerm, err := queries.SaveTerm(ctx, params)
	if err != nil {
		return database.Term{}, err
	}
	return dbTerm, nil
}

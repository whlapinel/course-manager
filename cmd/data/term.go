package data

import (
	"context"
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"strconv"
	"time"
)

// GetTerm implements TermRepo.
func (cr CourseRepo) GetTerm(date time.Time) (*domain.Term, error) {
	dbTerm, err := cr.queries.GetTerm(context.Background(), date.Format(time.DateOnly))
	if err != nil {
		return nil, err
	}
	start, err := time.Parse(time.DateOnly, dbTerm.Start)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.DateOnly, dbTerm.End)
	if err != nil {
		return nil, err
	}
	term := &domain.Term{
		ID:    int(dbTerm.ID),
		Name:  dbTerm.Name,
		Start: start,
		End:   end,
	}
	if term.Start.IsZero() {
		return nil, fmt.Errorf("term.Start not initialized")
	}
	return term, nil
}

func (cr CourseRepo) GetTermByID(termID int) (domain.Term, error) {
	dbTerm, err := cr.queries.GetTermByID(context.Background(), int64(termID))
	if err != nil {
		return domain.Term{}, nil
	}
	dates, err := parseDates([]string{dbTerm.Start, dbTerm.End})
	if err != nil {
		return domain.Term{}, nil
	}
	return domain.Term{
		ID:    int(dbTerm.ID),
		Name:  dbTerm.Name,
		Start: dates[0],
		End:   dates[1],
	}, nil

}

func (cr CourseRepo) GetTerms() ([]domain.Term, error) {
	dbTerms, err := cr.queries.GetTerms(context.Background())
	if err != nil {
		return nil, err
	}
	var terms []domain.Term
	for _, dbTerm := range dbTerms {
		parsedStart, err := time.Parse(time.DateOnly, dbTerm.Start)
		if err != nil {
			return nil, err
		}
		parsedEnd, err := time.Parse(time.DateOnly, dbTerm.End)
		if err != nil {
			return nil, err
		}
		term := domain.Term{
			ID:    int(dbTerm.ID),
			Start: parsedStart,
			End:   parsedEnd,
			Name:  dbTerm.Name,
		}
		terms = append(terms, term)

	}
	return terms, nil

}

func (cr CourseRepo) GetTermDates(termID int) (domain.Term, error) {
	var term domain.Term
	dbDates, err := cr.queries.GetTermDates(context.Background(), int64(termID))
	if len(dbDates) == 0 {
		log.Println("dates returned: 0. CourseRepo.GetTermDates")
	}
	if err != nil {
		return term, nil
	}
	for _, dbDate := range dbDates {
		date, err := time.Parse(time.DateOnly, dbDate.Date)
		if err != nil {
			return term, err
		}
		term.InstructionalDays = append(term.InstructionalDays, date)
	}
	return term, nil
}

// ReadFromCSV implements TermRepo.
func (t CourseRepo) ReadFromCSV() ([]domain.Term, error) {
	terms, err := t.ImportTermsFromCSV()
	if err != nil {
		return nil, err
	}
	return terms, nil
}

func (t CourseRepo) SaveTerm(term domain.Term) (int, error) {
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

const termsPath = "/home/whlapinel/personal_projects/course_manager/cmd/data/csv_files/terms.csv"
const nonIDaysPath = "/home/whlapinel/personal_projects/course_manager/cmd/data/csv_files/non_instruct_days.csv"

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

func (t CourseRepo) ImportTermsFromCSV() ([]domain.Term, error) {
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
	terms := []domain.Term{}
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

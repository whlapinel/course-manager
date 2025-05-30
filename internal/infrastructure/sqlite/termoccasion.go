package sqlite

import (
	"context"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/features/termoccasion"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"time"
)

type termOccasionRepo struct {
	queries *database.Queries
}

func NewTermOccasionRepo(queries *database.Queries) termoccasion.Repository {
	return &termOccasionRepo{
		queries: queries,
	}
}

// ByID implements termoccasion.Repository.
func (t *termOccasionRepo) ByID(id int) (occasion.Occasion, error) {
	var occasion occasion.Occasion
	dbOccasion, err := t.queries.GetTermOccasionByID(context.Background(), int64(id))
	if err != nil {
		return occasion, err
	}
	date, err := time.Parse(time.DateOnly, dbOccasion.Date)
	if err != nil {
		return occasion, err
	}
	occasion.Name = dbOccasion.Name
	occasion.Date = date
	occasion.ID = int(dbOccasion.ID)
	occasion.ParentID = int(dbOccasion.TermID)
	return occasion, nil
}

// ByParentID implements termoccasion.Repository.
func (t *termOccasionRepo) ByParentID(termID int) ([]occasion.Occasion, error) {
	dbOccasions, err := t.queries.GetTermOccasions(context.Background(), int64(termID))
	if err != nil {
		return nil, err
	}
	var occasions []occasion.Occasion
	for _, dbOccasion := range dbOccasions {
		date, err := time.Parse(time.DateOnly, dbOccasion.Date)
		if err != nil {
			return nil, err
		}
		occasion := occasion.Occasion{
			ID:       int(dbOccasion.ID),
			ParentID: int(dbOccasion.TermID),
			Date:     date,
			Name:     dbOccasion.Name,
		}
		occasions = append(occasions, occasion)
	}
	return occasions, nil

}

// Delete implements termoccasion.Repository.
func (t *termOccasionRepo) Delete(id int) error {
	return t.queries.DeleteTermOccasion(context.Background(), int64(id))
}

// Save implements termoccasion.Repository.
func (t *termOccasionRepo) Save(occ occasion.Occasion) (int, error) {
	dbTO, err := t.queries.SaveTermOccasion(context.Background(), database.SaveTermOccasionParams{
		TermID: int64(occ.ParentID),
		Date:   occ.Date.Format(time.DateOnly),
		Name:   occ.Name,
	})
	if err != nil {
		return 0, err
	}
	return int(dbTO.ID), nil
}

// Update implements termoccasion.Repository.
func (t *termOccasionRepo) Update(occ occasion.Occasion) error {
	return t.queries.UpdateTermOccasion(context.Background(), database.UpdateTermOccasionParams{
		ID:   int64(occ.ID),
		Name: occ.Name,
	})
}

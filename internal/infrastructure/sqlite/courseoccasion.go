package sqlite

import (
	"context"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/features/termoccasion"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"time"
)

type courseOccasionRepo struct {
	queries *database.Queries
}

// ByID implements termoccasion.Repository.
func (t *courseOccasionRepo) ByID(id int) (occasion.Occasion, error) {
	panic("unimplemented")
}

// ByTermID implements termoccasion.Repository.
func (t *courseOccasionRepo) ByTermID(termID int) ([]occasion.Occasion, error) {
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
func (t *courseOccasionRepo) Delete(id occasion.Occasion) error {
	panic("unimplemented")
}

// Save implements termoccasion.Repository.
func (t *courseOccasionRepo) Save(occasion.Occasion) (int, error) {
	panic("unimplemented")
}

// Update implements termoccasion.Repository.
func (t *courseOccasionRepo) Update(occasion.Occasion) error {
	panic("unimplemented")
}

func NewCourseOccasionRepo(queries *database.Queries) termoccasion.Repository {
	return &courseOccasionRepo{
		queries: queries,
	}
}

package sqlite

import (
	"context"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/features/courseoccasion"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"time"
)

type courseOccasionRepo struct {
	queries *database.Queries
}

func NewCourseOccasionRepo(queries *database.Queries) courseoccasion.Repository {
	return &courseOccasionRepo{
		queries: queries,
	}
}

func (t *courseOccasionRepo) Delete(id int) error {
	return t.queries.DeleteCourseOccasion(context.Background(), int64(id))
}

func (t *courseOccasionRepo) ByID(id int) (occasion.Occasion, error) {
	var occasion occasion.Occasion
	dbOccasion, err := t.queries.GetCourseOccasionByID(context.Background(), int64(id))
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
	occasion.ParentID = int(dbOccasion.CourseID)
	return occasion, nil
}

func (t *courseOccasionRepo) ByParentID(courseID int) ([]occasion.Occasion, error) {
	dbOccasions, err := t.queries.GetCourseOccasions(context.Background(), int64(courseID))
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
			ParentID: int(dbOccasion.CourseID),
			Date:     date,
			Name:     dbOccasion.Name,
		}
		occasions = append(occasions, occasion)
	}
	return occasions, nil

}

func (t *courseOccasionRepo) Save(occ occasion.Occasion) (int, error) {
	dbTO, err := t.queries.SaveCourseOccasion(context.Background(), database.SaveCourseOccasionParams{
		CourseID: int64(occ.ParentID),
		Date:     occ.Date.Format(time.DateOnly),
		Name:     occ.Name,
	})
	if err != nil {
		return 0, err
	}
	return int(dbTO.ID), nil
}

func (t *courseOccasionRepo) Update(occ occasion.Occasion) error {
	return t.queries.UpdateCourseOccasion(context.Background(), database.UpdateCourseOccasionParams{
		ID:   int64(occ.ID),
		Name: occ.Name,
	})
}

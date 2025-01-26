package data

import (
	"context"
	"gh_static_portfolio/internal/domain"
)

func (cr CourseRepo) SaveStandardSet(set domain.StandardSet) (domain.StandardSet, error) {
	dbSet, err := cr.queries.SaveStandardSet(context.Background(), set.CourseName)
	if err != nil {
		return domain.StandardSet{}, err
	}
	return domain.StandardSet{
		ID:         int(dbSet.ID),
		CourseName: dbSet.CourseName,
	}, nil
}

func (cr CourseRepo) GetStandardSets() ([]domain.StandardSet, error) {
	dbSets, err := cr.queries.GetAllStandardSets(context.Background())
	if err != nil {
		return nil, err
	}
	var standardSets []domain.StandardSet
	for _, dbSet := range dbSets {
		set := domain.StandardSet{
			ID:         int(dbSet.ID),
			CourseName: dbSet.CourseName,
		}
		standardSets = append(standardSets, set)
	}
	return standardSets, nil
}

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

func (cr CourseRepo) GetStandardSetByID(id int) (domain.StandardSet, error) {
	dbSet, err := cr.queries.GetStdSetByID(context.Background(), int64(id))
	if err != nil {
		return domain.StandardSet{}, err
	}
	set := domain.StandardSet{
		ID:         int(dbSet.ID),
		CourseName: dbSet.CourseName,
	}
	standards, err := cr.GetCourseStandards(set)
	if err != nil {
		return domain.StandardSet{}, err
	}
	for i, standard := range standards {
		objectives, err := cr.GetStandardObjectives(standard, set)
		if err != nil {
			return domain.StandardSet{}, err
		}
		standards[i].Children = objectives
	}
	set.Standards = standards
	return set, nil
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

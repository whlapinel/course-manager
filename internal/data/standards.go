package data

import (
	"context"
	"database/sql"
	"encoding/csv"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"os"
	"strconv"
)

const (
	StdNo = iota
	StdDescr
	StdCourse // name
)

const (
	ObjNo = iota
	ObjWeight
	ObjDescr
	ObjCourse
)

const (
	Python1 = "Python Programming I Honors"
	Python2 = "Python Programming II Honors"
)

func (cr CourseRepo) SaveStandard(std domain.Standard) (domain.Standard, error) {
	dbStandard, err := cr.queries.SaveStandard(context.Background(), database.SaveStandardParams{
		ParentID: sql.NullInt64{
			Valid: std.ParentID != 0,
			Int64: int64(std.ParentID),
		},
		SetID:  int64(std.StdSet.ID),
		Number: int64(std.Number),
		Name:   std.Name,
		Description: sql.NullString{
			Valid:  std.Description != "",
			String: std.Description,
		},
	})
	if err != nil {
		return domain.Standard{}, err
	}
	std.ID = int(dbStandard.ID)
	return std, nil
}
func (cr CourseRepo) ImportStandards(filename string, setID int) ([]domain.Standard, error) {
	set, err := cr.queries.GetStdSetByID(context.Background(), int64(setID))
	if err != nil {
		return nil, err
	}
	courseName := set.CourseName
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var standards []domain.Standard
	for i, record := range records {
		if i == 0 {
			continue
		}
		stdNum, err := strconv.Atoi(record[StdNo])
		if err != nil {
			return nil, err
		}
		stdDesc := record[StdDescr]
		stdCourse := record[StdCourse]
		standard := domain.Standard{
			StdSet: domain.StandardSet{
				ID:         setID,
				CourseName: courseName,
			},
			Number: stdNum,
			Name:   stdDesc,
		}
		if stdCourse == courseName {
			standards = append(standards, standard)
		}
	}
	return standards, nil
}

func (cr CourseRepo) ImportObjectives(filename string, stdID int) ([]domain.Standard, error) {
	return nil, nil
}

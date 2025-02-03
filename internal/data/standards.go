package data

import (
	"context"
	"database/sql"
	"encoding/csv"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"log"
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

func (cr CourseRepo) GetObjectiveByID(objectiveID int) (domain.Objective, error) {
	standard, err := cr.GetStandardByID(objectiveID)
	if err != nil {
		return domain.Objective{}, err
	}
	return standard.Objective, nil
}

func (cr CourseRepo) GetStandardByID(standardID int) (domain.Standard, error) {
	dbStandard, err := cr.queries.GetStandardByID(context.Background(), int64(standardID))
	if err != nil {
		return domain.Standard{}, err
	}
	standard := domain.Standard{
		StdSet: domain.StandardSet{
			ID: int(dbStandard.ID),
		},
		Objective: domain.Objective{
			ID:          int(dbStandard.ID),
			ParentID:    int(dbStandard.ParentID.Int64),
			Number:      int(dbStandard.Number),
			ParentNum:   int(dbStandard.ParentNum),
			Name:        dbStandard.Name,
			Description: dbStandard.Description.String,
		},
	}
	return standard, nil
}
func (cr CourseRepo) GetStandardObjectives(standard domain.Standard, set domain.StandardSet) ([]domain.Objective, error) {
	dbObjectives, err := cr.queries.GetStandardObjectives(context.Background(), sql.NullInt64{
		Valid: standard.ID != 0,
		Int64: int64(standard.ID),
	})
	if err != nil {
		return nil, err
	}
	var objectives []domain.Objective
	for _, dbObjective := range dbObjectives {
		objective := domain.Objective{
			ID:          int(dbObjective.ID),
			ParentID:    int(dbObjective.ParentID.Int64),
			Number:      int(dbObjective.Number),
			ParentNum:   int(dbObjective.ParentNum),
			Name:        dbObjective.Name,
			Description: dbObjective.Description.String,
		}
		objectives = append(objectives, objective)
	}
	return objectives, nil
}

func (cr CourseRepo) GetCourseStandards(set domain.StandardSet) ([]domain.Standard, error) {
	dbStandards, err := cr.queries.GetCourseStandards(context.Background(), int64(set.ID))
	if err != nil {
		return nil, err
	}
	var standards []domain.Standard
	for _, dbStandard := range dbStandards {
		standard := domain.Standard{
			StdSet: set,
			Objective: domain.Objective{
				ID:          int(dbStandard.ID),
				Number:      int(dbStandard.Number),
				Name:        dbStandard.Name,
				Description: dbStandard.Description.String,
			},
		}
		standards = append(standards, standard)
	}
	return standards, nil
}

func (cr CourseRepo) SaveObjective(obj domain.Objective) (domain.Objective, error) {
	standard := domain.Standard{
		Objective: obj,
	}
	standard, err := cr.SaveStandard(standard)
	if err != nil {
		return domain.Objective{}, err
	}
	return standard.Objective, nil
}

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

func (cr CourseRepo) ImportStandards(filename string, set domain.StandardSet) ([]domain.Standard, error) {
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
		stdCourse := record[StdCourse]
		if stdCourse != courseName {
			continue
		}
		stdNum, err := strconv.Atoi(record[StdNo])
		if err != nil {
			return nil, err
		}
		stdDesc := record[StdDescr]
		standard := domain.Standard{
			StdSet: domain.StandardSet{
				ID:         set.ID,
				CourseName: courseName,
			},
			Objective: domain.Objective{
				Number: stdNum,
				Name:   stdDesc,
			},
		}
		standards = append(standards, standard)
	}
	return standards, nil
}

func (cr CourseRepo) ImportObjectives(filename string, set domain.StandardSet, standard domain.Standard) ([]domain.Objective, error) {
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
	var objectives []domain.Objective
	for i, record := range records {
		if i == 0 {
			continue
		}
		objCourse := record[ObjCourse]
		if objCourse != courseName {
			continue
		}
		stdNum, err := strconv.Atoi(record[ObjNo][:1])
		if err != nil {
			return nil, err
		}
		if stdNum != standard.Number {
			continue
		}
		objNum, err := strconv.Atoi(record[ObjNo][2:])
		if err != nil {
			return nil, err
		}
		if objNum == 0 {
			continue
		}
		objDescr := record[ObjDescr]
		objective := domain.Objective{

			ParentID: standard.ID,
			Number:   objNum,
			Name:     objDescr,
		}
		objectives = append(objectives, objective)
	}
	log.Println(objectives)
	return objectives, nil

}

func (cr CourseRepo) CreateStandardSets(filename string) ([]domain.StandardSet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var courseNames = make(map[string]struct{})
	var sets []domain.StandardSet
	for i, record := range records {
		if i == 0 {
			continue
		}
		courseName := record[StdCourse]
		courseNames[courseName] = struct{}{}
	}
	for name := range courseNames {
		set, err := cr.SaveStandardSet(domain.StandardSet{
			CourseName: name,
		})
		if err != nil {
			return nil, err
		}
		sets = append(sets, set)
	}
	return sets, nil

}

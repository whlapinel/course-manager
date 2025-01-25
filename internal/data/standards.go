package data

import (
	"encoding/csv"
	"fmt"
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

func (cr CourseRepo) ImportStandards(filename string, courseName string, setID int) ([]domain.Standard, error) {
	if courseName != Python1 && courseName != Python2 {
		return nil, fmt.Errorf("courseName does not match list")
	}
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

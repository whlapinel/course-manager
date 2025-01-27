package data

import (
	"context"
	"database/sql"
	"log"
	"testing"
)

// Warning: this changes the database!!
func TestImportStandards(t *testing.T) {
	sets, err := cr.CreateStandardSets("./csv_files/standards.csv")
	if err != nil {
		t.Error(err)
	}
	for _, set := range sets {
		standards, err := cr.ImportStandards("./csv_files/standards.csv", set)
		if err != nil {
			t.Error(err)
		}
		for _, standard := range standards {
			standard, err := cr.SaveStandard(standard)
			if err != nil {
				t.Error(err)
			}
			objectives, err := cr.ImportObjectives("./csv_files/objectives.csv", set, standard)
			if err != nil {
				t.Error(err)
			}
			if len(objectives) == 0 {
				t.Error("objectives is empty")
			}
			for _, obj := range objectives {
				log.Println(obj)
				_, err := cr.SaveStandard(obj)
				if err != nil {
					t.Error(err)
				}
			}
		}
	}

}

func TestGetStandards(t *testing.T) {
	sets, err := cr.GetStandardSets()
	if err != nil {
		t.Error(err)
	}
	for _, set := range sets {
		log.Println(set.CourseName)
		standards, err := cr.queries.GetCourseStandards(context.Background(), int64(set.ID))
		if err != nil {
			t.Error(err)
		}
		log.Println("STANDARDS for", set.CourseName)
		for _, standard := range standards {
			objectives, err := cr.queries.GetStandardObjectives(context.Background(), sql.NullInt64{
				Valid: standard.ID != 0,
				Int64: standard.ID,
			})
			if err != nil {
				t.Error(err)
			}
			log.Println("OBJECTIVES for", standard.Number, standard.Name)
			for _, objective := range objectives {
				log.Println(objective.Number, objective.Name, objective.ParentID.Int64)
			}
		}
	}

}

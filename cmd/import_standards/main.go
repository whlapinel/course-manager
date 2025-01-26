package main

import (
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"log"
)

func main() {
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := data.NewCourseRepo(queries)
	sets, err := CreateStandardSets(repo)
	if err != nil {
		log.Fatal(err)
	}
	for i, set := range sets {
		standards, err := repo.ImportStandards("internal/data/csv_files/standards.csv", set.ID)
		if err != nil {
			log.Fatal(err)
		}
		for j, standard := range standards {
			standard, err = repo.SaveStandard(standard)
			if err != nil {
				log.Fatal(err)
			}
			objectives, err := repo.ImportObjectives("internal/data/csv_files/objectives.csv", standard.ID)
			if err != nil {
				log.Fatal(err)
			}
			for k, objective := range objectives {
				objective, err = repo.SaveStandard(objective)
				if err != nil {
					log.Fatal(err)
				}
				objectives[k] = objective
			}
			standards[j].Children = objectives
		}
		sets[i].Standards = standards
	}
	for _, set := range sets {
		log.Println(set.ID, set.CourseName)
		for _, standard := range set.Standards {
			log.Println(standard.ID, standard.Number, standard.Name)
			for _, obj := range standard.Children {
				log.Println(obj.ID, obj.Number, obj.Name)
			}
		}
	}
}

func CreateStandardSets(repo data.CourseRepo) ([]domain.StandardSet, error) {
	python1, err := repo.SaveStandardSet(domain.StandardSet{
		CourseName: data.Python1,
	})
	if err != nil {
		return nil, err
	}
	python2, err := repo.SaveStandardSet(domain.StandardSet{
		CourseName: data.Python2,
	})
	if err != nil {
		return nil, err
	}
	return []domain.StandardSet{
		python1,
		python2,
	}, nil

}

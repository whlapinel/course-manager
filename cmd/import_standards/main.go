package main

import (
	"gh_static_portfolio/internal/data"
	"log"
)

func main() {
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := data.NewCourseRepo(queries)
	repo.ImportStandards("internal/data/csv_files/standards.csv", data.Python1)
}

func CreateStandardSets(repo data.CourseRepo) error {
	
}

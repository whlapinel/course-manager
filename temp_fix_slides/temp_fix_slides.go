package main

import (
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
)

func main() {
	queries, db, err := data.InitDB("../course_manager.db")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	terms, err := courseRepo.GetTerms()
	for _, term := range terms {
		courses, err := courseRepo.GetCourses(term.ID)
		if err != nil {
			log.Fatal(err)
		}

		for _, course := range courses {

			units, err := courseRepo.GetUnits(course.ID)
			if err != nil {
				log.Fatal(err)
			}

			for _, unit := range units {

				lessons, err := courseRepo.GetLessons(unit.ID)
				if err != nil {
					log.Fatal(err)
				}

				for _, lesson := range lessons {
					srcPath := data.OldSlidesMarkdownFilePath(lesson)
					slides := domain.NewSlides(lesson.Name, lesson.Description, srcPath)
					_, err := courseRepo.SaveSlides(slides, lesson)
					if err != nil {
						log.Fatal(err)
					}
					err = os.Remove(srcPath)
					if err != nil {
						log.Fatal(err)
					}

				}
			}
		}

	}

}

package main

import (
	"errors"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"io/fs"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	queries, db, err := data.InitDB("course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	terms, err := courseRepo.GetTerms()
	if err != nil {
		log.Fatal(err)
	}
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
					_, err := os.Stat(srcPath)
					if err != nil {
						if errors.Is(err, fs.ErrNotExist) {
							continue
						}
					}
					slides := domain.NewSlides(lesson.Name, lesson.Description, srcPath)
					_, err = courseRepo.SaveSlides(slides, lesson)
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

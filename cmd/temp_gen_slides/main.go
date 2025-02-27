// this is solely a one-time program for writing generated slides html files in the source files directories
package main

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const production = "production"
const development = "development"

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(logger)
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	startMarp := exec.Command("marp", "-s", "internal/data")
	err = startMarp.Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 30)
	defer func() {
		startMarp.Process.Signal(os.Interrupt)
	}()
	defer db.Close()
	repo := data.NewCourseRepo(queries)
	svc := service.NewCourseService(repo)
	user, err := repo.GetUser("101602110272674353046")
	if err != nil {
		log.Fatal("error getting user", err)
	}
	terms, err := repo.GetTerms(user.ID)
	if err != nil {
		log.Fatal("error getting terms", err)
	}
	for _, term := range terms {
		courses, err := repo.GetCourses(term.ID)
		if err != nil {
			log.Fatal("error getting courses", err)
		}
		for _, course := range courses {
			units, err := repo.GetUnits(course.ID)
			if err != nil {
				log.Fatal("error getting units", err)
			}
			for _, unit := range units {
				lessons, err := repo.GetLessons(unit.ID)
				if err != nil {
					log.Fatal("error getting units", err)
				}
				for _, lesson := range lessons {
					path := domain.NodePath{
						UserID:   user.ID,
						TermID:   term.ID,
						CourseID: course.ID,
						UnitID:   unit.ID,
						LessonID: lesson.ID,
					}
					nodes := domain.Nodes{
						User:   user,
						Term:   term,
						Course: course,
						Unit:   unit,
						Lesson: lesson,
					}
					markdownPath := data.SlidesMarkdownFilePath(nodes.ToSlice()...)
					_, err := os.Stat(markdownPath)
					if err != nil {
						if os.IsNotExist(err) {
							continue
						}
					}
					slides, err := svc.GetSlides(path)
					if err != nil {
						log.Println(err)
					}
					htmlPath := data.SlidesHTMLFilePath(nodes.ToSlice()...)
					html, err := os.Create(htmlPath)
					if err != nil {
						log.Println(err)
					}
					written, err := html.Write([]byte(slides))
					if err != nil {
						log.Println(err)
					}
					log.Printf("%d written to file %s", written, htmlPath)
					html.Close()

				}
			}
		}
	}

}

func LoadEnvironment() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("ENV") == development {
		godotenv.Load(".env.development")
		return development, nil
	}
	if os.Getenv("ENV") == production {
		godotenv.Load(".env.production")
		return production, nil
	}
	return "", fmt.Errorf("environment not expected:'%s'", os.Getenv("ENV"))
}

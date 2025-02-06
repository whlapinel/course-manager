package sitegenerator

import (
	"context"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates"
	sst "gh_static_portfolio/internal/templates/statictemplates"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	_ "github.com/mattn/go-sqlite3"
)

// Generate creates the static site using concurrent operations.
func Generate(courseRepo data.CourseRepo) error {
	log.Println("sitegenerator.Generate(): generating site")

	// Run build commands concurrently.
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		GenerateTempl()
		return nil
	})
	g.Go(func() error {
		BuildTailwind()
		return nil
	})
	g.Go(func() error {
		BuildTypeScript()
		return nil
	})
	if err := g.Wait(); err != nil {
		return err
	}

	// Clear HTML files concurrently in each directory.
	clearDirs := templates.DirectoriesClearList()
	g, ctx = errgroup.WithContext(ctx)
	for _, dir := range clearDirs {
		directory := dir // capture loop variable
		g.Go(func() error {
			return ClearHTMLFiles(directory)
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	// Delete directories concurrently.
	deleteDirs := templates.DeleteDirList()
	g, ctx = errgroup.WithContext(ctx)
	for _, dir := range deleteDirs {
		directory := dir
		g.Go(func() error {
			err := os.RemoveAll(directory)
			if err != nil {
				return fmt.Errorf("failed to delete directory %s: %v", directory, err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	// Render the home and contact pages concurrently.
	homePage := sst.NewHomePage()
	contactPage := sst.NewContactPage()
	g, ctx = errgroup.WithContext(ctx)
	g.Go(func() error {
		return RenderPage(homePage)
	})
	g.Go(func() error {
		return RenderPage(contactPage)
	})
	if err := g.Wait(); err != nil {
		return err
	}

	// Get the terms from the course repository and select the current term.
	terms, err := courseRepo.GetTerms()
	if err != nil {
		return fmt.Errorf("error fetching term: %s", err)
	}
	var currentTerm domain.Term
	for _, term := range terms {
		if term.Start.Before(time.Now().AddDate(0, 0, 7)) && term.End.After(time.Now()) {
			currentTerm = term
		}
	}
	if currentTerm.Start.IsZero() {
		return errors.New("main(): term not initialized")
	}

	// Get courses for the current term.
	courses, err := courseRepo.GetCourses(currentTerm.ID)
	if err != nil {
		return fmt.Errorf("error getting courses: %v", err)
	}

	// Render the courses list page concurrently.
	coursesPage := sst.NewCoursesListPage(courses)
	g, ctx = errgroup.WithContext(ctx)
	g.Go(func() error {
		return RenderPage(coursesPage)
	})

	// Use sync.Once to ensure the node directory copy happens only once.
	var copyOnce sync.Once
	var copyErr error

	// Process each course concurrently.
	for _, course := range courses {
		course := course // capture loop variable
		g.Go(func() error {
			// Copy the node directory only once.
			copyOnce.Do(func() {
				copyErr = data.CopyNodeDir(data.NodeDirPath(currentTerm), templates.StaticSiteRootDir)
			})
			if copyErr != nil {
				return copyErr
			}
			if course.Term.Start.IsZero() {
				return errors.New("main(): instance.Term.Start is zero")
			}
			if err := RenderPage(sst.NewCourseCalendarPage(course)); err != nil {
				return fmt.Errorf("failed to render calendar page: %v", err)
			}
			if err := RenderPage(sst.NewCoursePage(course)); err != nil {
				return fmt.Errorf("failed to render course page: %v", err)
			}

			// Process each unit in the course concurrently.
			var unitGroup errgroup.Group
			for _, unit := range course.Units {
				unit := unit // capture loop variable
				unitGroup.Go(func() error {
					if err := RenderPage(sst.NewUnitPage(unit, course)); err != nil {
						return fmt.Errorf("failed to render unit page: %v", err)
					}
					lessons, err := courseRepo.GetLessons(unit.ID)
					if err != nil {
						return fmt.Errorf("failed to get lessons for unit %v: %w", unit, err)
					}
					// Process each lesson concurrently.
					var lessonGroup errgroup.Group
					for _, lesson := range lessons {
						lesson := lesson // capture loop variable
						lessonGroup.Go(func() error {
							lessonDates, err := courseRepo.GetLessonDates(lesson.ID)
							if err != nil {
								return fmt.Errorf("failed to get lesson dates: %v", err)
							}
							lesson.Dates = lessonDates

							objectives, err := courseRepo.GetLessonObjectives(lesson)
							if err != nil {
								return fmt.Errorf("failed to get lesson objectives: %v", err)
							}
							lesson.Standards = objectives

							// Generate slides for the lesson (this function logs internally).
							GenerateSlides(currentTerm, course, unit, lesson)

							if err := RenderPage(sst.NewLessonPage(lesson, unit, course, true, true)); err != nil {
								return fmt.Errorf("failed to render lesson page: %v", err)
							}
							return nil
						})
					}
					if err := lessonGroup.Wait(); err != nil {
						return err
					}
					return nil
				})
			}
			if err := unitGroup.Wait(); err != nil {
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// ClearHTMLFiles removes all .html files from the given directory.
func ClearHTMLFiles(directory string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			// Use filepath.Join to build the path safely.
			err := os.Remove(filepath.Join(directory, file.Name()))
			if err != nil {
				return fmt.Errorf("failed to delete file %s: %v", file.Name(), err)
			}
		}
	}
	return nil
}

// BuildTailwind runs the Tailwind build command.
func BuildTailwind() {
	err := exec.Command("npx", "tailwindcss", "-i", "./input.css", "-o", "./python/docs/styles/styles.css").Run()
	if err != nil {
		log.Println(err)
	}
}

// GenerateTempl runs the templating command.
func GenerateTempl() {
	err := exec.Command("templ", "generate").Run()
	if err != nil {
		log.Println(err)
	}
}

// BuildTypeScript runs the TypeScript build command.
func BuildTypeScript() {
	err := exec.Command("tsc", "--build").Run()
	if err != nil {
		log.Println(err)
	}
}

// RenderPage creates the directory for the page, creates the file, and renders the page component.
func RenderPage(page sst.Page) error {
	err := os.MkdirAll(filepath.Dir(page.Path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	f, err := os.Create(page.Path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer f.Close()
	err = page.Component.Render(context.Background(), f)
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}
	return nil
}

// GenerateSlides regenerates the slides HTML if the Markdown file has been updated.
func GenerateSlides(term, course, unit, lesson domain.CourseNode) {
	log.Println("GenerateSlides(): generating slides")
	// File paths.
	markdownPath := data.SlidesMarkdownFilePath(term, course, unit, lesson)
	log.Println("markdownPath:", markdownPath)
	htmlPath := data.SlidesHTMLFilePath(term, course, unit, lesson)
	log.Println("htmlPath:", htmlPath)

	// Get file information.
	mdInfo, err := os.Stat(markdownPath)
	if err != nil {
		log.Println("file not exists:", os.IsNotExist(err))
		log.Println(err)
		if !errors.Is(err, fs.ErrNotExist) {
			log.Printf("Error: Could not access %s: %v\n", markdownPath, err)
		}
		return
	}

	htmlInfo, err := os.Stat(htmlPath)
	log.Println("file not exists:", os.IsNotExist(err))
	log.Println(err)
	if os.IsNotExist(err) {
		// If HTML file doesn't exist, regenerate it.
		regenerateHTML(markdownPath, htmlPath)
		return
	} else if err != nil {
		log.Printf("Error: Could not access %s: %v\n", htmlPath, err)
		return
	}

	// Compare modification times.
	mdModTime := mdInfo.ModTime()
	htmlModTime := htmlInfo.ModTime()
	if mdModTime.After(htmlModTime) {
		regenerateHTML(markdownPath, htmlPath)
	}
	log.Println("reached end of GenerateSlides()")
}

// regenerateHTML runs the command to regenerate HTML from a Markdown file.
func regenerateHTML(markdownFile, htmlFile string) {
	cmd := exec.Command("marp", markdownFile, "-o", htmlFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: Failed to regenerate HTML: %v\n", err)
		return
	}
	fmt.Printf("HTML file %s successfully regenerated from %s\n", htmlFile, markdownFile)
}

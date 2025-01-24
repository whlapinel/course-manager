package sitegenerator

import (
	"context"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates"
	sst "gh_static_portfolio/internal/templates/static_site_templates"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Generate(courseRepo data.CourseRepo) error {
	log.Println("sitegenerator.Generate(): generating site")
	GenerateTempl()
	BuildTailwind()
	BuildTypeScript()
	directories := templates.DirectoriesClearList()
	for _, directory := range directories {
		ClearHTMLFiles(directory)
	}
	// delete ./python/docs/courses so there are no orphan files
	for _, dir := range templates.DeleteDirList() {
		err := os.RemoveAll(dir)
		if err != nil {
			return fmt.Errorf("failed to delete directory: %s: %v", dir, err)
		}
	}

	// Generate home page
	homePage := sst.NewHomePage()
	err := RenderPage(homePage)
	if err != nil {
		return fmt.Errorf("failed to render pages: %v", err)
	}
	// Generate contact page
	contactPage := sst.NewContactPage()
	err = RenderPage(contactPage)
	if err != nil {
		return fmt.Errorf("failed to render pages: %v", err)
	}
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
	log.Println()
	if currentTerm.Start.IsZero() {
		return fmt.Errorf("main(): term not initialized")
	}
	// Generate "courses I teach" list page
	courses, err := courseRepo.GetCourses(currentTerm.ID)
	if err != nil {
		return fmt.Errorf("error getting instances: %v", err)
	}
	coursesPage := sst.NewCoursesListPage(courses)
	err = RenderPage(coursesPage)
	if err != nil {
		return fmt.Errorf("failed to render pages: %v", err)
	}
	for _, course := range courses {
		err = data.CopyNodeDir(data.NodeDirPath(currentTerm), templates.StaticSiteRootDir)
		if err != nil {
			return err
		}
		if course.Term.Start.IsZero() {
			return fmt.Errorf("main(): instance.Term.Start is zero")
		}
		// Generate calendar page for each course
		calendarPage := sst.NewCourseCalendarPage(*course)
		err = RenderPage(calendarPage)
		if err != nil {
			return fmt.Errorf("failed to render pages: %v", err)
		}
		coursePage := sst.NewCoursePage(*course)
		err = RenderPage(coursePage)
		if err != nil {
			return fmt.Errorf("failed to render pages: %v", err)
		}
		for _, unit := range course.Units {
			unitPage := sst.NewUnitPage(*unit, *course)
			err = RenderPage(unitPage)
			if err != nil {
				return fmt.Errorf("failed to render pages: %v", err)
			}
			lessons, err := courseRepo.GetLessons(unit.ID)
			if err != nil {
				return fmt.Errorf("failed to get lessons: %s", err)
			}
			// Generate page for each lesson
			for _, lesson := range lessons {
				lessonPage := sst.NewLessonPage(*lesson, *unit, *course, true, true)
				err = RenderPage(lessonPage)
				if err != nil {
					return fmt.Errorf("failed to render pages: %v", err)
				}
			}
		}
	}
	return nil
}

func ClearHTMLFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			err := os.Remove(directory + file.Name())
			if err != nil {
				log.Fatalf("failed to delete file: %v", err)
			}
		}
	}
}

func BuildTailwind() {
	err := exec.Command("npx", "tailwindcss", "-i", "./input.css", "-o", "./python/docs/styles/styles.css").Run()
	if err != nil {
		log.Println(err)
	}
}

func GenerateTempl() {
	err := exec.Command("templ", "generate").Run()
	if err != nil {
		log.Println(err)
	}
}

func BuildTypeScript() {
	err := exec.Command("tsc", "--build").Run()
	if err != nil {
		log.Println(err)
	}
}

func RenderPage(page sst.Page) error {
	err := os.MkdirAll(filepath.Dir(page.Path), os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}
	f, err := os.Create(page.Path)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	err = page.Component.Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}

// compares the last time the files were modified. if markdown was modified after html, regenerates html
func GenerateSlides(term, course, unit, lesson domain.CourseNode) {
	log.Println("GenerateSlides(): generating slides")
	// File paths
	markdownPath := data.SlidesMarkdownFilePath(term, course, unit, lesson)
	log.Println("markdownPath:", markdownPath)
	htmlPath := data.SlidesHTMLFilePath(term, course, unit, lesson)
	log.Println("htmlPath:", htmlPath)

	// Get file information
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
		// If HTML file doesn't exist, regenerate it
		regenerateHTML(markdownPath, htmlPath)
		return
	} else if err != nil {
		log.Printf("Error: Could not access %s: %v\n", htmlPath, err)
		return
	}

	// Compare modification times
	mdModTime := mdInfo.ModTime()
	htmlModTime := htmlInfo.ModTime()
	if mdModTime.After(htmlModTime) {
		regenerateHTML(markdownPath, htmlPath)
	}
	log.Println("reached end of GenerateSlides()")
}

// Regenerate HTML from the Markdown file
func regenerateHTML(markdownFile, htmlFile string) {
	cmd := exec.Command("marp", markdownFile, "-o", htmlFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: Failed to regenerate HTML: %v\n", err)
		return
	}
	fmt.Printf("HTML file %s successfully regenerated from %s\n", htmlFile, markdownFile)
}

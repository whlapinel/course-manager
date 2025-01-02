package sitegenerator

import (
	"context"
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/templates"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Generate() {
	GenerateTempl()
	BuildTailwind()
	BuildTypeScript()
	directories := templates.DirectoriesClearList()
	for _, directory := range directories {
		ClearHTMLFiles(directory)
	}
	// Generate home page
	homePage := templates.NewHomePage()
	err := RenderPage(homePage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	// Generate contact page
	contactPage := templates.NewContactPage()
	err = RenderPage(contactPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	// Database
	queries, db, err := data.InitDB("course_manager.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	courseRepo := data.NewCourseRepo(queries)
	terms, err := courseRepo.GetTerms()
	if err != nil {
		log.Fatalf("error fetching term: %s", err)
	}
	var currentTerm domain.Term
	for _, term := range terms {
		if term.Start.Before(time.Now()) && term.End.After(time.Now()) {
			currentTerm = term
		}
	}
	log.Println()
	if currentTerm.Start.IsZero() {
		log.Fatal("main(): term not initialized")
	}
	// Generate "courses I teach" list page
	courses, err := courseRepo.GetCourses(currentTerm.ID)
	if err != nil {
		log.Fatalf("error getting instances: %v", err)
	}
	coursesPage := templates.NewCoursesListPage(courses)
	err = RenderPage(coursesPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	for _, course := range courses {
		if course.Term.Start.IsZero() {
			log.Fatal("main(): instance.Term.Start is zero")
		}
		// Generate calendar page for each course
		calendarPage := templates.NewCourseCalendarPage(*course)
		err = RenderPage(calendarPage)
		if err != nil {
			log.Fatalf("failed to render pages: %v", err)
		}
		log.Println("Site generator main() course: Name: ", course.Name)
		coursePage := templates.NewCoursePage(*course)
		err = RenderPage(coursePage)
		if err != nil {
			log.Fatalf("failed to render pages: %v", err)
		}
		// Generate page for each unit
		for _, unit := range course.Units {
			log.Println("looping through units in main():", unit.Name)
			unitPage := templates.NewUnitPage(*unit, *course)
			err = RenderPage(unitPage)
			if err != nil {
				log.Fatalf("failed to render pages: %v", err)
			}
			// Generate page for each lesson
			for _, lesson := range unit.Lessons {
				lessonPage := templates.NewLessonPage(*lesson, *unit, *course)
				GenerateSlides(filepath.Dir(lessonPage.Path))
				err = RenderPage(lessonPage)
				if err != nil {
					log.Fatalf("failed to render pages: %v", err)
				}
			}
		}
	}
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

func RenderPage(page templates.Page) error {
	log.Println("RenderPage: ", page.Title)
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

func GenerateSlides(dir string) {
	// File paths
	markdownFile := "slides.md"
	htmlFile := "slides.html"
	markdownPath := path.Join(dir, markdownFile)
	htmlPath := path.Join(dir, htmlFile)

	// Get file information
	mdInfo, err := os.Stat(markdownPath)
	if err != nil {
		log.Printf("Error: Could not access %s: %v\n", markdownPath, err)
		return
	}

	htmlInfo, err := os.Stat(htmlPath)
	if os.IsNotExist(err) {
		// If HTML file doesn't exist, regenerate it
		log.Println("HTML file does not exist. Generating...")
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
		log.Println("Markdown file is newer. Regenerating HTML...")
		regenerateHTML(markdownPath, htmlPath)
	} else {
		log.Println("No need to regenerate. HTML is up-to-date.")
	}
}

// this is a temporary function to save all slides in the main directory, this function should be deleted
func SaveSlides(lesson domain.Lesson, unit domain.Unit, course domain.Course) error {
	srcPath := templates.SlidesMarkdownPath(lesson, unit, course)
	dstPath := data.SlidesMarkdownFilePath(lesson)
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() error {
		err := srcFile.Close()
		if err != nil {
			return err
		}
		return nil
	}()
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	bytes, err := io.Copy(srcFile, dstFile)
	if err != nil {
		return err
	}
	log.Println("Copied ", bytes, " from ", srcPath, " to", dstPath)
	return nil
}

func SaveFiles(lesson domain.Lesson, unit domain.Unit, course domain.Course, cr data.CourseRepo) error {
	filesDir := templates.LessonFilesPath(lesson, unit, course)
	files, err := os.ReadDir(filesDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		srcPath := filepath.Join(filesDir, file.Name())
		newFile := domain.NewFile(file.Name(), "transferred from generator", srcPath)
		newFile.ID, err = cr.SaveFile(newFile)
		if err != nil {
			return err
		}
		err := cr.AddFileToLesson(newFile, lesson)
		if err != nil {
			return err
		}
	}
	return nil
}

// This will copy the html files from the ./cmd/data/slides directory.
// It should replace the generate slides function.
// Slides will be generated by Fyne app rather than in Generator
func CopySlides(lesson domain.Lesson, unit domain.Unit, course domain.Course) error {
	srcPath := data.SlidesHTMLFilePath(lesson)
	dstPath := templates.SlidesPath(lesson, unit, course)
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() error {
		err := srcFile.Close()
		if err != nil {
			return err
		}
		return nil
	}()
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	bytes, err := io.Copy(srcFile, dstFile)
	if err != nil {
		return err
	}
	log.Println("Copied ", bytes, " from ", srcPath, " to", dstPath)
	return nil
}

func CopyFiles(lesson domain.Lesson, unit domain.Unit, course domain.Course, cr data.CourseRepo) error {
	files, err := cr.GetLessonFiles(lesson)
	if err != nil {
		return err
	}
	for _, file := range files {
		srcPath := data.LessonFilePath(file)
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer func() error {
			err := srcFile.Close()
			if err != nil {
				return err
			}
			return nil
		}()
		dstPath := templates.LessonFilesURL(lesson, unit, course)
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		bytes, err := io.Copy(srcFile, dstFile)
		if err != nil {
			return err
		}
		log.Println("Copied ", bytes, " from ", srcPath, " to", dstPath)
	}
	return nil
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

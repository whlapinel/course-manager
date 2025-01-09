package sitegenerator

import (
	"context"
	"errors"
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/templates"
	"io"
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
	homePage := templates.NewHomePage()
	err := RenderPage(homePage)
	if err != nil {
		return fmt.Errorf("failed to render pages: %v", err)
	}
	// Generate contact page
	contactPage := templates.NewContactPage()
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
		if term.Start.Before(time.Now()) && term.End.After(time.Now()) {
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
	coursesPage := templates.NewCoursesListPage(courses)
	err = RenderPage(coursesPage)
	if err != nil {
		return fmt.Errorf("failed to render pages: %v", err)
	}
	for _, course := range courses {
		if course.Term.Start.IsZero() {
			return fmt.Errorf("main(): instance.Term.Start is zero")
		}
		// Generate calendar page for each course
		calendarPage := templates.NewCourseCalendarPage(*course)
		err = RenderPage(calendarPage)
		if err != nil {
			return fmt.Errorf("failed to render pages: %v", err)
		}
		coursePage := templates.NewCoursePage(*course)
		err = RenderPage(coursePage)
		if err != nil {
			return fmt.Errorf("failed to render pages: %v", err)
		}
		err = CopyCourseImage(*course)
		if err != nil {
			return err
		}
		// Generate page for each unit
		for _, unit := range course.Units {
			unitPage := templates.NewUnitPage(*unit, *course)
			err = RenderPage(unitPage)
			if err != nil {
				return fmt.Errorf("failed to render pages: %v", err)
			}
			err = CopyUnitImage(*unit, *course)
			if err != nil {
				return fmt.Errorf("failed to copy unit image: %s", err)
			}
			lessons, err := courseRepo.GetLessons(unit.ID)
			if err != nil {
				return fmt.Errorf("failed to get lessons: %s", err)
			}
			// Generate page for each lesson
			for _, lesson := range lessons {
				_, err := os.Stat(data.NewSlidesMarkdownFilePath(lesson.ID))
				hasSlides := err != nil
				_, err = os.Stat(data.NewLessonFilesDirPath(lesson.ID))
				hasFiles := err != nil
				lessonPage := templates.NewLessonPage(*lesson, *unit, *course, hasSlides, hasFiles)
				err = RenderPage(lessonPage)
				if err != nil {
					return fmt.Errorf("failed to render pages: %v", err)
				}
				if hasSlides {
					GenerateSlides(*lesson)
					err = CopySlides(*lesson, *unit, *course)
					if err != nil {
						log.Println("failed to copy slides: ", err)
					}
				}
				if hasFiles {
					err = CopyFiles(*lesson, *unit, *course, courseRepo)
					if err != nil {
						log.Println("failed to copy files: ", err)
					}
				}
				err = CopyLessonImage(*lesson, *unit, *course)
				if err != nil {
					log.Println("failed to copy image: ", err)
				}
			}
		}
	}
	return nil
}

func CopyFiles(lesson domain.Lesson, unit domain.Unit, course domain.Course, cr data.CourseRepo) error {
	srcRoot := data.NewLessonFilesDirPath(lesson.ID)
	// if directory doesn't exist, early return
	_, err := os.Stat(srcRoot)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		} else {
			return err
		}
	}
	files, err := os.ReadDir(srcRoot)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}
	log.Println("Source ", srcRoot)
	destRoot := templates.LessonFilesPath(lesson, unit, course)
	log.Println("Dest ", destRoot)
	return filepath.WalkDir(srcRoot, func(srcPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destRoot, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}
		return copyFile(srcPath, destPath)
	})
}

func copyFile(srcPath, destPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

func CopyCourseImage(course domain.Course) error {
	srcPath := data.OldImagesPath(course.Image)
	if !FileExists(srcPath) {
		return nil
	}
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	dstPath := templates.CourseImagePath(course)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}
func CopyUnitImage(unit domain.Unit, course domain.Course) error {
	if unit.Image.ID == 0 {
		return nil
	}
	srcPath := data.OldImagesPath(unit.Image)
	if !FileExists(srcPath) {
		return fmt.Errorf("file not found: %s", srcPath)
	}
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	dstPath := templates.UnitImagePath(unit, course)
	err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
	if err != nil {
		return err
	}
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}
func CopyLessonImage(lesson domain.Lesson, unit domain.Unit, course domain.Course) error {
	srcPath := data.OldImagesPath(lesson.Image)
	if !FileExists(srcPath) {
		return nil
	}
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	dstPath := templates.LessonImagePath(lesson, unit, course)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
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

func RenderPage(page templates.Page) error {
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

func GenerateSlides(lesson domain.Lesson) {
	// File paths
	markdownPath := data.NewSlidesMarkdownFilePath(lesson.ID)
	htmlPath := data.NewSlidesHTMLFilePath(lesson.ID)

	// Get file information
	mdInfo, err := os.Stat(markdownPath)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			log.Printf("Error: Could not access %s: %v\n", markdownPath, err)
		}
		return
	}

	htmlInfo, err := os.Stat(htmlPath)
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
}

// // this is a temporary function to save all slides in the main directory, this function should be deleted once that process is complete
// func SaveSlides(lesson domain.Lesson, unit domain.Unit, course domain.Course) error {
// 	srcPath := templates.SlidesMarkdownPath(lesson, unit, course)
// 	dstPath := data.SlidesMarkdownFilePath(lesson)
// 	err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
// 	errInfo := fmt.Sprintf("srcPath: %s\ndstPath: %s", srcPath, dstPath)
// 	if err != nil {
// 		return nil
// 	}
// 	srcFile, err := os.Open(srcPath)
// 	if err != nil {
// 		return fmt.Errorf(errInfo, err)
// 	}
// 	defer func() error {
// 		err := srcFile.Close()
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}()
// 	dstFile, err := os.Create(dstPath)
// 	if err != nil {
// 		return fmt.Errorf(errInfo, err)
// 	}
// 	bytes, err := io.Copy(dstFile, srcFile)
// 	if err != nil {
// 		return fmt.Errorf(errInfo, err)
// 	}
// 	log.Println("Copied ", bytes, " from ", srcPath, " to", dstPath)
// 	return nil
// }

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
		panic("this should not happen")
	}
	return true
}

// This will copy the html files from the ./cmd/data/slides directory.
// It should replace the generate slides function.
// Slides will be generated by Fyne app rather than in Generator
func CopySlides(lesson domain.Lesson, unit domain.Unit, course domain.Course) error {
	srcPath := data.NewSlidesHTMLFilePath(lesson.ID)
	dstPath := templates.SlidesPath(lesson, unit, course)
	srcFile, err := os.Open(srcPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
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

package sitegenerator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates"
	managertemplates "gh_static_portfolio/internal/templates/manager_templates"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Generate creates the static site using concurrent operations.
func Generate(courseRepo data.CourseRepo, userID string) error {
	user := domain.User{
		ID: userID,
	}
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
	homePage := mt.StaticNewHomePage()
	contactPage := mt.StaticNewContactPage()
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
	terms, err := courseRepo.GetTerms(userID)
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
	occasions, err := courseRepo.GetTermOccasions(currentTerm.ID)
	if err != nil {
		return err
	}
	currentTerm.Occasions = occasions
	// Get courses for the current term.
	courses, err := courseRepo.GetCourses(currentTerm.ID)
	if err != nil {
		return fmt.Errorf("error getting courses: %v", err)
	}

	// Render the courses list page concurrently.
	coursesPage := mt.StaticNewCoursesListPage(courses)
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
		course.Term = currentTerm
		g.Go(func() error {
			// Copy the node directory only once.
			copyOnce.Do(func() {
				copyErr = data.CopyNodeDir(data.NodeDirPath(user, currentTerm), templates.StaticSiteRootDir)
			})
			if copyErr != nil {
				return copyErr
			}
			if course.Term.Start.IsZero() {
				return errors.New("main(): instance.Term.Start is zero")
			}
			if err := RenderPage(mt.StaticNewCourseCalendarPage(course)); err != nil {
				return fmt.Errorf("failed to render calendar page: %v", err)
			}
			if err := RenderPage(mt.StaticNewCoursePage(course)); err != nil {
				return fmt.Errorf("failed to render course page: %v", err)
			}

			// Process each unit in the course concurrently.
			var unitGroup errgroup.Group
			for _, unit := range course.Units {
				unit := unit // capture loop variable
				unitGroup.Go(func() error {
					if err := RenderPage(mt.StaticNewUnitPage(unit, course)); err != nil {
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

							assessments, err := courseRepo.GetLessonAssessments(lesson.ID)
							if err != nil {
								return fmt.Errorf("failed to get lesson assessments: %v", err)
							}
							lesson.Assessments = assessments

							// Generate slides for the lesson (this function logs internally).
							err = GenerateSlides(user, currentTerm, course, unit, lesson)
							if err != nil {
								if !os.IsNotExist(err) {
									return err
								}
							}
							page := mt.LessonPage{
								Lesson:    lesson,
								Unit:      unit,
								Course:    course,
								HasSlides: true,
								HasFiles:  true,
							}

							filesDir := templates.LessonFilesPath(lesson, unit, course)
							RenderMarkdownFiles(lesson.Name, filesDir)

							if err := RenderPage(mt.StaticNewLessonPage(page)); err != nil {
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
func RenderPage(page mt.StaticPage) error {
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

// this will generate html from all markdown files within the files directory
func RenderMarkdownFiles(title, filesPath string) error {
	dirEntries, err := os.ReadDir(filesPath)
	if os.IsNotExist(err) {
		return err
	}
	if err != nil {
		log.Fatal(err)
	}

	// Filter out non-Markdown files
	var mdFiles []os.DirEntry
	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".md" {
			mdFiles = append(mdFiles, entry)
		}
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
			),
		)),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var markdownGroup errgroup.Group
	for _, entry := range mdFiles {
		log.Println("rendering markdown for:", entry.Name())

		inputPath := filepath.Join(filesPath, entry.Name())
		markdownGroup.Go(func() error {
			// Read entire file at once
			contents, err := os.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("Failed to read file %s: %v", inputPath, err)
			}

			var buf bytes.Buffer
			if err := md.Convert(contents, &buf); err != nil {
				log.Fatalf("Failed to convert Markdown: %v", err)
			}

			outputPath := filepath.Join(filesPath, strings.TrimSuffix(entry.Name(), ".md")+".html")
			output, err := os.Create(outputPath)
			if err != nil {
				log.Fatalf("Failed to create output file %s: %v", outputPath, err)
			}
			defer output.Close()

			log.Println("Writing:", outputPath)
			data := mt.MarkdownDocument{
				Title:   title,
				Content: buf.String(),
				Static:  true,
			}
			err = mt.DocLayout(data).Render(context.Background(), output)
			if err != nil {
				return err
			}
			return nil
		})

	}
	if err := markdownGroup.Wait(); err != nil {
		return err
	}
	return nil
}

// GenerateSlides regenerates the slides HTML if the Markdown file has been updated.
func GenerateSlides(nodes ...domain.CourseNode) error {
	// File paths.
	markdownPath := data.SlidesMarkdownFilePath(nodes...)
	lesson, ok := nodes[len(nodes)-1].(domain.Lesson)
	if !ok {
		return fmt.Errorf("node is not a lesson: %v", nodes[len(nodes)-1])
	}
	unit, ok := nodes[len(nodes)-2].(domain.Unit)
	if !ok {
		return fmt.Errorf("node is not a unit: %v", nodes[len(nodes)-2])
	}
	course, ok := nodes[len(nodes)-3].(domain.Course)
	if !ok {
		return fmt.Errorf("node is not a course: %v", nodes[len(nodes)-3])
	}
	if lesson.ID == 377 {
		log.Println("generating slides for lesson ID:", lesson.ID)
	}
	htmlPath := templates.SlidesPath(lesson, unit, course)
	log.Println("htmlPath:", htmlPath)

	// Get file information.
	mdInfo, err := os.Stat(markdownPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Error: Could not access %s: %v\n", markdownPath, err)
		}
		return err
	}
	htmlInfo, err := os.Stat(htmlPath)
	if os.IsNotExist(err) {
		// If HTML file doesn't exist, regenerate it.
		// regenerateHTML(markdownPath, htmlPath)
		if lesson.ID == 377 {
			log.Println("regenerating HTML for", lesson.ID, "at path:", htmlPath)
		}
		err := newRegenerateHTML(htmlPath, nodes...)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		log.Printf("Error: Could not access %s: %v\n", htmlPath, err)
		return fs.ErrNotExist
	}

	// Compare modification times.
	mdModTime := mdInfo.ModTime()
	htmlModTime := htmlInfo.ModTime()
	if lesson.ID == 377 {
		log.Println("markdown modtime", lesson.ID, mdModTime)
		log.Println("html modtime", lesson.ID, htmlModTime)
	}
	if mdModTime.After(htmlModTime) {
		// regenerateHTML(markdownPath, htmlPath)
		err := newRegenerateHTML(htmlPath, nodes...)
		if err != nil {
			return err
		}
	}
	return nil
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

func newRegenerateHTML(htmlFile string, nodes ...domain.CourseNode) error {
	log.Println("regenerating html for:", htmlFile)
	var params mt.NodePath
	for i, node := range nodes {
		if i == 0 {
			params.UserID.Value = node.GetID()
		} else if i == 1 {
			params.TermID.Value = node.GetID()
		} else if i == 2 {
			params.CourseID.Value = node.GetID()
		} else if i == 3 {
			params.UnitID.Value = node.GetID()
		} else if i == 4 {
			params.LessonID.Value = node.GetID()
		}
	}
	slidesContent, err := GetSlides(params)
	if err != nil {
		return err
	}
	file, err := os.Create(htmlFile)
	if err != nil {
		return err
	}
	written, err := file.Write([]byte(slidesContent))
	if err != nil {
		return err
	}
	log.Println(written, "bytes written to ", htmlFile)
	return nil
}

func GetSlides(params managertemplates.NodePath) (string, error) {
	path, err := marpSlidesPath(params)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func marpSlidesPath(params managertemplates.NodePath) (string, error) {
	baseURL := "http://localhost:8080"
	userParam := fmt.Sprintf("user_%s", params.UserID.Value)
	termParam := fmt.Sprintf("term_%d", params.TermID.Value)
	courseParam := fmt.Sprintf("course_%d", params.CourseID.Value)
	unitParam := fmt.Sprintf("unit_%d", params.UnitID.Value)
	lessonParam := fmt.Sprintf("lesson_%d", params.LessonID.Value)
	return url.JoinPath(baseURL, "users", userParam, "terms", termParam, "courses", courseParam, "units", unitParam, "lessons", lessonParam, "slides.md")
}

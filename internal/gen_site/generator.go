package sitegenerator

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/app"
	templates "gh_static_portfolio/internal/templates/shared"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

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

func GetSlides(params mt.NodePath) (string, error) {
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

func marpSlidesPath(params mt.NodePath) (string, error) {
	baseURL := "http://localhost:8080"
	userParam := fmt.Sprintf("user_%s", params.UserID.Value)
	termParam := fmt.Sprintf("term_%d", params.TermID.Value)
	courseParam := fmt.Sprintf("course_%d", params.CourseID.Value)
	unitParam := fmt.Sprintf("unit_%d", params.UnitID.Value)
	lessonParam := fmt.Sprintf("lesson_%d", params.LessonID.Value)
	return url.JoinPath(baseURL, "users", userParam, "terms", termParam, "courses", courseParam, "units", unitParam, "lessons", lessonParam, "slides.md")
}

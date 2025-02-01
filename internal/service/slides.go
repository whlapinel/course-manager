package service

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	sitegenerator "gh_static_portfolio/internal/gen_site"
	managertemplates "gh_static_portfolio/internal/templates/manager_templates"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func marpSlidesPath(params managertemplates.CourseIDParams) (string, error) {
	baseURL := "http://localhost:8080"
	termParam := fmt.Sprintf("term_%d", params.TermID.Value)
	courseParam := fmt.Sprintf("course_%d", params.CourseID.Value)
	unitParam := fmt.Sprintf("unit_%d", params.UnitID.Value)
	lessonParam := fmt.Sprintf("lesson_%d", params.LessonID.Value)
	return url.JoinPath(baseURL, termParam, "courses", courseParam, "units", unitParam, "lessons", lessonParam, "slides.md")
}

// This check to see if the file already exists. if not, should create a new markdown file, write the template to it,
// and generate the html file. Returns the file path
func (svc CourseService) CreateSlidesIfNotExist(nodes ...domain.CourseNode) (string, error) {
	markdownPath := data.SlidesMarkdownFilePath(nodes...)
	_, err := os.Stat(markdownPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(markdownPath), os.ModePerm)
		if err != nil {
			return "", err
		}
		file, err := os.Create(markdownPath)
		if err != nil {
			return "", err
		}
		// write template to file
		templateFileContents, err := svc.SlidesTemplate(nodes[len(nodes)-1])
		if err != nil {
			return "", err
		}
		_, err = file.Write(templateFileContents)
		if err != nil {
			return "", err
		}
	}
	sitegenerator.GenerateSlides(nodes[0], nodes[1], nodes[2], nodes[3])
	return markdownPath, nil
}

func (svc CourseService) SlidesTemplate(lesson domain.CourseNode) ([]byte, error) {
	templateFileContents, err := os.ReadFile("./cmd/web_app/slide_template.md")
	if err != nil {
		return nil, err
	}
	templateOther := fmt.Sprintf(
		`
# %s

# **Warmup**
		
# **Agenda**
		
# **Looking ahead**`,

		lesson.GetName())
	templateFileContents = append(templateFileContents, []byte(templateOther)...)
	return templateFileContents, nil
}

func (svc CourseService) GetSlides(params managertemplates.CourseIDParams) (string, error) {
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

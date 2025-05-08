package templates

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/core/assessment"
	"gh_static_portfolio/internal/shared/node"
	components "gh_static_portfolio/internal/newtemplates/components/base"
	"path/filepath"

	"github.com/a-h/templ"
)

type Link struct {
	Text string
	URL  string
}
type StaticLessonDetailsPage struct {
	node.Nodes
	StaticNodeDetailsPage
	LessonSlidesURL string
	AssessmentsURL  string
	ViewMarkdownURL func(relPath string, nodes ...node.Node) string
	FilesURLFunc    func(relPath string, nodes ...node.Node) string
}

func (page StaticLessonDetailsPage) Component() templ.Component {
	return Layout{
		PageTitle: page.Node.GetName(),
		Head: Head{
			User:      page.User.(dto.User),
			AssetsURL: page.AssetsURL,
		}.Component(),
		Page: StaticLessonDetailsComponent(page),
	}.Component()
}

func (page StaticLessonDetailsPage) Info() templ.Component {
	return page.Objectives()
}

func (page StaticLessonDetailsPage) Objectives() templ.Component {
	lesson := page.Lesson.(dto.Lesson)
	if len(lesson.Standards) == 0 {
		return nil
	}
	var table components.Table
	table.Title = "Objectives"
	var headers = []string{"Designation", "Name"}
	var rows [][]templ.Component
	for _, obj := range lesson.Standards {
		row := []templ.Component{
			components.TableTextCell{Text: obj.Designation()}.Component(),
			components.TableTextCell{Text: obj.Name}.Component(),
		}
		rows = append(rows, row)
	}
	table.Headers = headers
	table.Rows = rows
	return table.Component()

}

func (page StaticLessonDetailsPage) Tabs() templ.Component {
	assetsURLFunc := func(relPath string) string {
		path := page.AssetsURL("js")
		path = filepath.Join(path, relPath)
		return path
	}
	data := components.TabSet{
		AssetsURLFunc: assetsURLFunc,
		Tabs: []components.TabLink{
			{Name: "Slides", URL: page.LessonSlidesURL},
			{Name: "Files", URL: page.StaticNodeDetailsPage.FilesPageURL},
			{Name: "Assessments", URL: page.AssessmentsURL},
		},
	}
	return data.Component()
}

type AssessmentsFragment struct {
	StaticLessonDetailsPage
	Assessments []assessment.Assessment
	Path        string
}

func (data AssessmentsFragment) Filepath() string {
	return data.Path
}

func (data AssessmentsFragment) Component() templ.Component {
	return AssessmentsComponent(data)
}

type SlidesFragment struct {
	Path            string
	LessonSlidesURL string
}

func (data SlidesFragment) Component() templ.Component {
	return SlidesComponent(data.LessonSlidesURL)
}

func (data SlidesFragment) Filepath() string {
	return data.Path
}

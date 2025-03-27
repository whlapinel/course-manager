package templates

import (
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"
	"path/filepath"

	"github.com/a-h/templ"
)

type Link struct {
	Text string
	URL  string
}
type StaticLessonDetailsPage struct {
	domain.Nodes
	StaticNodeDetailsPage
	LessonSlidesURL string
	AssessmentsURL  string
	ViewMarkdownURL func(relPath string, nodes ...domain.CourseNode) string
	FilesURLFunc    func(relPath string, nodes ...domain.CourseNode) string
}

func (page StaticLessonDetailsPage) Component() templ.Component {
	return Layout{
		PageTitle: page.Node.GetName(),
		Head: Head{
			User:      page.User,
			AssetsURL: page.AssetsURL,
		}.Component(),
		Page: StaticLessonDetailsComponent(page),
	}.Component()
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
	Assessments []domain.Assessment
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

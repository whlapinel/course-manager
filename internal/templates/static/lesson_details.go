package templates

import (
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"

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
	ViewMarkdownURL func(relPath string, nodes ...domain.CourseNode) string
	FilesURLFunc    func(relPath string, nodes ...domain.CourseNode) string
}

func (page StaticLessonDetailsPage) Component() templ.Component {
	return components.Layout{
		PageTitle: page.Node.GetName(),
		Head: Head{
			User:      page.User,
			AssetsURL: page.AssetsURL,
		}.Component(),
		Page: StaticLessonDetailsComponent(page),
	}.Component()
}

func (page StaticLessonDetailsPage) Slides() templ.Component {
	return SlidesComponent(page)
}

func (page StaticLessonDetailsPage) Assessments() templ.Component {
	return AssessmentsComponent(page.Lesson.Assessments, page)
}

type NewStaticLessonDetailsPageParams struct {
	StaticLessonDetailsPage
}

func NewStaticLessonDetailsPage(params NewStaticLessonDetailsPageParams) StaticLessonDetailsPage {
	return params.StaticLessonDetailsPage
}

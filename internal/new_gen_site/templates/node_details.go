package templates

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates/components"

	"github.com/a-h/templ"
)

type StaticNodeDetailsPage struct {
	StaticLayout
	Path string
	Node domain.CourseNode
}

func (page StaticNodeDetailsPage) BreadCrumbs() templ.Component {
	return page.StaticLayout.BreadCrumbs.Component()
}
func (page StaticNodeDetailsPage) Component() templ.Component {
	return components.Layout{
		PageTitle: page.Node.GetName(),
		Head: Head{
			User:      page.User,
			AssetsURL: page.AssetsURL,
		}.Component(),
		Page: StaticNodeDetailsComponent(page),
	}.Component()

}

func (page StaticNodeDetailsPage) Layout() StaticLayout {
	return page.StaticLayout
}

func (page StaticNodeDetailsPage) Filepath() string {
	return page.Path
}

type StaticNodeDetailsParams struct {
	StaticLayout
	Node domain.CourseNode
	Path string
}

func NewStaticNodeDetailsPage(params StaticNodeDetailsParams) StaticNodeDetailsPage {
	return StaticNodeDetailsPage{
		StaticLayout: params.StaticLayout,
		Path:         params.Path,
		Node:         params.Node,
	}

}

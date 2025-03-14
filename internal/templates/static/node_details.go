package templates

import (
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type StaticNodeDetailsPage struct {
	PageData
	Path string
	Node domain.CourseNode
}

func (page StaticNodeDetailsPage) BreadCrumbs() templ.Component {
	return page.PageData.BreadCrumbs.Component()
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

func (page StaticNodeDetailsPage) Layout() PageData {
	return page.PageData
}

func (page StaticNodeDetailsPage) Filepath() string {
	return page.Path
}

type StaticNodeDetailsParams struct {
	PageData
	Node domain.CourseNode
	Path string
}

func NewStaticNodeDetailsPage(params StaticNodeDetailsParams) StaticNodeDetailsPage {
	return StaticNodeDetailsPage{
		PageData: params.PageData,
		Path:     params.Path,
		Node:     params.Node,
	}

}

package templates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
)

type StaticNodeDetailsPage struct {
	Path string
	Node domain.CourseNode
}

func (page StaticNodeDetailsPage) Component() templ.Component {
	return StaticNodeDetailsComponent(page)
}

func (page StaticNodeDetailsPage) Filepath() string {
	return page.Path
}

type StaticNodeDetailsParams struct {
	Node domain.CourseNode
	Path string
}

func NewStaticNodeDetailsPage(params StaticNodeDetailsParams) StaticNodeDetailsPage {
	return StaticNodeDetailsPage{
		Path: params.Path,
		Node: params.Node,
	}

}

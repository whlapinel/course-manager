package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	tpl "gh_static_portfolio/internal/templates"

	"github.com/a-h/templ"
)

type NodeCreatePage struct {
	ParentNode        domain.CourseNode
	NodeType          domain.NodeTypeName
	Params            CourseIDParams
	PostCreateNodeURL string
	CancelURL         string
	BreadCrumbsData   BreadCrumbs
}

func (page NodeCreatePage) Component() templ.Component {
	return NodeCreateComponent(page)
}

func (page NodeCreatePage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: fmt.Sprintf("New %s for %s", page.NodeType.String(), page.ParentNode.GetName()),
		UpNav: UpNav{
			URL:  page.CancelURL,
			Text: "Cancel",
		},
	}
}

func (page NodeCreatePage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData

}

func NodeCreateFormID(nodeType domain.NodeTypeName) string {
	return tpl.KebabCase(nodeType.String()) + "-form"
}

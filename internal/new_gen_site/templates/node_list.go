package templates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates/components"
	"log"

	"github.com/a-h/templ"
)

type StaticLayout struct {
	User      domain.User
	AssetsURL func(relPath string) string
	NavItems  []components.NavItem
}

type StaticPage interface {
	Component() templ.Component
	Filepath() string
}

type StaticNodeListPage struct {
	Path       string
	Parent     domain.CourseNode
	ChildLinks []ChildNodeLink
}

func (page StaticNodeListPage) Filepath() string {
	return page.Path
}

type StaticNodeListParams struct {
	Nodes        domain.Nodes
	Path         string
	ChildUrlFunc func(nodes ...domain.CourseNode) string
}

func NewStaticNodeListPage(params StaticNodeListParams) (StaticNodeListPage, error) {
	var page StaticNodeListPage
	currentNode := params.Nodes.CurrentNode()
	children := currentNode.Children()
	log.Println("current node:", currentNode.TypeName())
	log.Println("children:", children[0].TypeName())
	if len(children) == 0 {
		return StaticNodeListPage{}, fmt.Errorf("no children for %s", currentNode.GetName())
	}
	for _, child := range children {
		link := ChildNodeLink{
			Text: child.GetName(),
			URL:  params.ChildUrlFunc(params.Nodes.ToSlice(child)...),
		}
		log.Println("link URL", link.URL)
		page.Parent = params.Nodes.CurrentNode()
		page.Path = params.Path
		page.ChildLinks = append(page.ChildLinks, link)

	}
	return page, nil

}

func (page StaticNodeListPage) Component() templ.Component {
	return StaticNodeListComponent(page)
}

type ChildNodeLink struct {
	URL  string
	Text string
}

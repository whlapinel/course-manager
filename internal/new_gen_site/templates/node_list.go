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
	components.BreadCrumbs
}

type StaticPage interface {
	Component() templ.Component
	Filepath() string
	Layout() StaticLayout
}

type StaticNodeListPage struct {
	StaticLayout
	Path   string
	Parent domain.CourseNode
	components.Table
	ChildLinks []ChildNodeLink
}

func (page StaticNodeListPage) BreadCrumbs() templ.Component {
	return page.StaticLayout.BreadCrumbs.Component()
}
func (page StaticNodeListPage) PageHeading() templ.Component {
	return components.PageHeading{
		Title: page.Parent.GetName(),
	}.Component()
}

func (page StaticNodeListPage) Layout() StaticLayout {
	return page.StaticLayout
}

func (page StaticNodeListPage) Filepath() string {
	return page.Path
}

type StaticNodeListParams struct {
	StaticLayout
	Nodes                    domain.Nodes
	Path                     string
	ListChildChildrenURLFunc func(nodes ...domain.CourseNode) string
	ChildDetailsURLFunc      func(nodes ...domain.CourseNode) string
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
	childHasChildren := children[0].ChildTypeName() != ""
	var tableHeaders []string
	if childHasChildren {
		tableHeaders = []string{
			"Details", "Children", "Designation", "Name",
		}
	} else {
		tableHeaders = []string{
			"Details", "Designation", "Name",
		}
	}
	var rows [][]templ.Component
	for _, child := range children {
		var row []templ.Component
		details := components.TableLinkCell{
			Text: "Details",
			URL:  params.ChildDetailsURLFunc(params.Nodes.ToSlice(child)...),
		}.Component()
		row = append(row, details)
		if childHasChildren {
			childChildren := components.TableLinkCell{
				Text: fmt.Sprintf("%ss", child.ChildTypeName()),
				URL:  params.ListChildChildrenURLFunc(params.Nodes.ToSlice(child)...),
			}.Component()
			row = append(row, childChildren)
		}
		desig := components.TableTextCell{
			Text: child.Designation(),
		}.Component()
		row = append(row, desig)
		nameText := components.TableTextCell{
			Text: child.GetName(),
		}.Component()
		row = append(row, nameText)
		rows = append(rows, row)
	}
	page.Parent = currentNode
	page.Path = params.Path
	page.StaticLayout = params.StaticLayout
	page.Table = components.Table{
		Title:   fmt.Sprintf("%ss in %s", currentNode.ChildTypeName(), currentNode.GetName()),
		Rows:    rows,
		Headers: tableHeaders,
	}

	// for _, child := range children {
	// 	link := ChildNodeLink{
	// 		Text: child.GetName(),
	// 		URL:  params.ChildUrlFunc(params.Nodes.ToSlice(child)...),
	// 	}
	// 	log.Println("link URL", link.URL)
	// 	page.ChildLinks = append(page.ChildLinks, link)

	// }
	return page, nil

}

func (page StaticNodeListPage) TableComponent() templ.Component {
	return page.Table.Component()
}

func (page StaticNodeListPage) Component() templ.Component {
	return components.Layout{
		PageTitle: page.Parent.GetName(),
		Page:      StaticNodeListComponent(page),
		Head: Head{
			User:      page.User,
			AssetsURL: page.AssetsURL,
		}.Component(),
	}.Component()
}

type ChildNodeLink struct {
	URL  string
	Text string
}

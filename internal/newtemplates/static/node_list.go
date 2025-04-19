package templates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type PageData struct {
	User      domain.User
	AssetsURL func(relPath string) string
	NavItems  []components.NavItem
	components.BreadCrumbs
}

type StaticPage interface {
	Component() templ.Component
	Filepath() string
	// Layout() PageData
}

type StaticNodeListPage struct {
	PageData
	Path   string
	Parent domain.CourseNode
	components.Table
}

func (page StaticNodeListPage) BreadCrumbs() templ.Component {
	return page.PageData.BreadCrumbs.Component()
}
func (page StaticNodeListPage) PageHeading() templ.Component {
	return components.PageHeading{
		Title: page.Parent.GetName(),
	}.Component()
}

func (page StaticNodeListPage) Layout() PageData {
	return page.PageData
}

func (page StaticNodeListPage) Filepath() string {
	return page.Path
}

type StaticNodeListParams struct {
	PageData
	Nodes                    domain.Nodes
	Path                     string
	ListChildChildrenURLFunc func(nodes ...domain.CourseNode) string
	ChildDetailsURLFunc      func(nodes ...domain.CourseNode) string
	CourseCalendarURL        func(nodes ...domain.CourseNode) string
}

func NewStaticNodeListPage(params StaticNodeListParams) (StaticNodeListPage, error) {
	var page StaticNodeListPage
	table, err := TableData(params)
	if err != nil {
		return StaticNodeListPage{}, err
	}
	page.Table = table
	page.Parent = params.Nodes.CurrentNode()
	page.Path = params.Path
	page.PageData = params.PageData
	return page, nil

}

func TableData(params StaticNodeListParams) (components.Table, error) {
	currentNode := params.Nodes.CurrentNode()
	children := currentNode.Children()
	if len(children) == 0 {
		return components.Table{}, fmt.Errorf("no children for %s", currentNode.GetName())
	}
	childHasChildren := children[0].ChildTypeName() != ""
	var tableHeaders []string
	childIsCourse := params.Nodes.CurrentNode().TypeName() == domain.TermTypeName.String()
	if childHasChildren {
		tableHeaders = []string{
			"Details", "Children", "Designation", "Name",
		}
		if childIsCourse {
			tableHeaders = []string{
				"Details", "Children", "Calendar", "Designation", "Name",
			}
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
			if childIsCourse {
				courseCal := components.TableLinkCell{
					Text: "Calendar",
					URL:  params.CourseCalendarURL(params.Nodes.ToSlice(child)...),
				}.Component()
				row = append(row, courseCal)
			}
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
	return components.Table{
		Title:   fmt.Sprintf("%ss in %s", currentNode.ChildTypeName(), currentNode.GetName()),
		Rows:    rows,
		Headers: tableHeaders,
	}, nil
}

func (page StaticNodeListPage) TableComponent() templ.Component {
	return page.Table.Component()
}

func (page StaticNodeListPage) Component() templ.Component {
	return Layout{
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

package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	tpl "gh_static_portfolio/internal/templates"
)

type NodeListPage struct {
	Params           CourseIDParams
	ParentNode       domain.CourseNode
	Children         []domain.CourseNode
	ChildUI          [][]ComponentData
	ChildDetailsRHN  string // details for child e.g. if listing units, this would be the route handler name to show unit details
	ChildChildrenRHN string // children of child e.g. if listing units, this would be the route handler name to list unit lessons
	ShowNewChildURL  string // e.g. if listing units, this would be the route handler name to show new unit form
	DeleteChildRHN   string
	UpNavURL         string
	E                *echo.Echo // for generating URLs from route handler name
	BreadCrumbsData  BreadCrumbs
}

type NodeDeleteButton struct {
	DeleteChildRHN string
}

func (page NodeListPage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData
}

func (page NodeListPage) DeleteNodeButton(node domain.CourseNode) templ.Component {
	return HXButton{
		Text:      "Delete",
		HxConfirm: fmt.Sprintf("Are you sure you want to delete %s '%s'", node.TypeName(), node.GetName()),
		Method:    HxDelete,
		URL:       page.E.Reverse(page.DeleteChildRHN, AddNodeChildIDToParams(page.Params, node.GetID()).ToSlice()...),
		HxTarget:  page.ListItemElementID(node).Selector(),
	}.Component()
}

func (page NodeListPage) NodeChildrenButton(node domain.CourseNode) templ.Component {
	return HXButton{
		Text:     node.ChildTypeName() + "s",
		Method:   HxGet,
		URL:      page.E.Reverse(page.ChildChildrenRHN, AddNodeChildIDToParams(page.Params, node.GetID()).ToSlice()...),
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeListPage) NodeDetailsButton(node domain.CourseNode) templ.Component {
	return HXButton{
		Text:     "Details",
		Method:   HxGet,
		URL:      page.E.Reverse(page.ChildDetailsRHN, AddNodeChildIDToParams(page.Params, node.GetID()).ToSlice()...),
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeListPage) NodeCreateButton() templ.Component {
	return HXButton{
		Text:     fmt.Sprintf("Add %s", page.ParentNode.ChildTypeName()),
		Method:   HxGet,
		URL:      page.ShowNewChildURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeListPage) Component() templ.Component {
	return NodeListComponent(page)
}

func (list NodeListPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: list.PageTitle(),
		UpNav: UpNav{
			URL:  list.UpNavURL,
			Text: list.upNavText(),
		},
	}

}

func (list NodeListPage) upNavText() string {
	if list.UpNavURL != "" && list.ParentNode != nil {
		return fmt.Sprintf("Up to %s", upNavText(list.ParentNode.TypeName()))
	} else {
		return ""
	}
}

func (list NodeListPage) PageTitle() string {
	if list.ParentNode.TypeName() != domain.RootTypeName.String() {
		return fmt.Sprintf("%ss for %s: %s", list.ParentNode.ChildTypeName(), list.ParentNode.TypeName(), list.ParentNode.GetName())
	} else {
		return fmt.Sprintf("%ss", list.ParentNode.ChildTypeName())
	}
}

func (list NodeListPage) ListItemElementID(node domain.CourseNode) ElementID {
	return ElementID(fmt.Sprintf("%s-%d", tpl.KebabCase(node.TypeName()), node.GetID()))
}

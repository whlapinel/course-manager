package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	tpl "gh_static_portfolio/internal/templates/shared"
)

type NodeListPage struct {
	Params           domain.NodePath
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
	return cmp.Button{
		HxConfirm: fmt.Sprintf("Are you sure you want to delete %s '%s'", node.TypeName(), node.GetName()),
		Method:    cmp.HxDelete,
		URL:       page.E.Reverse(page.DeleteChildRHN, AddParams(page.Params, node.GetID())...),
		HxTarget:  page.ListItemElementID(node).Selector(),
		Image:     cmp.DeleteImage(),
	}.Component()
}

func (page NodeListPage) NodeChildrenButton(node domain.CourseNode) templ.Component {
	return cmp.Button{
		Text:     node.ChildTypeName() + "s",
		Method:   cmp.HxGet,
		URL:      page.E.Reverse(page.ChildChildrenRHN, AddParams(page.Params, node.GetID())...),
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeListPage) NodeDetailsButton(node domain.CourseNode) templ.Component {
	return cmp.Button{
		Method:   cmp.HxGet,
		URL:      page.E.Reverse(page.ChildDetailsRHN, AddParams(page.Params, node.GetID())...),
		HxTarget: "#page",
		PushURL:  true,
		Image:    cmp.InfoIcon(),
	}.Component()
}

func (page NodeListPage) NodeCreateButton() templ.Component {
	return cmp.Button{
		Text:     page.ParentNode.ChildTypeName(),
		Method:   cmp.HxGet,
		URL:      page.ShowNewChildURL,
		HxTarget: "#page",
		PushURL:  true,
		Image:    cmp.AddImage(),
	}.Component()
}

func (page NodeListPage) Component() templ.Component {
	return NodeListComponent(page)
}

func (list NodeListPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: list.PageTitle(),
		UpNav: cmp.UpNav{
			URL:  list.UpNavURL,
			Text: list.upNavText(),
		},
		Crumbs: list.BreadCrumbs().BreadCrumbs(),
	}

}

func (list NodeListPage) upNavText() string {
	if list.UpNavURL != "" && list.ParentNode != nil {
		return fmt.Sprintf(" to %s", upNavText(list.ParentNode.TypeName()))
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

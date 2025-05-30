package appcomponents

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/util"
	"gh_static_portfolio/internal/shared/web"

	"github.com/a-h/templ"
)

type NodeListPage struct {
	ParentNode      ports.Node
	Children        []ports.Node
	ChildUI         [][]ComponentData
	ChildDetailsURL web.AddParams // details for child e.g. if listing units, this would be the route handler name to show unit details
	DeleteChildURL  web.AddParams
	ShowNewChildURL string // e.g. if listing units, this would be the route handler name to show new unit form
	UpNavURL        string
	BreadCrumbsData BreadCrumbs
	CourseManagerLayout
}

func (p NodeListPage) HTMXResponse() templ.Component {
	return p.Component()
}
func (p NodeListPage) NonHTMXResponse() templ.Component {
	return p.CourseManagerLayout.WithPage(p.Component())
}

type NodeDeleteButton struct {
	DeleteChildRHN string
}

func (page NodeListPage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData
}

func (page NodeListPage) DeleteNodeButton(node ports.Node) templ.Component {
	return cmp.Button{
		HxConfirm: fmt.Sprintf("Are you sure you want to delete %s '%s'", node.GetTypeName(), node.GetName()),
		Method:    cmp.HxDelete,
		URL:       page.DeleteChildURL(node.GetID()),
		HxTarget:  page.ListItemElementID(node).Selector(),
		Image:     cmp.DeleteImage(),
	}.Component()
}

func (page NodeListPage) NodeDetailsButton(node ports.Node) templ.Component {
	return cmp.Button{
		Method:   cmp.HxGet,
		URL:      page.ChildDetailsURL(node.GetID()),
		HxTarget: "#page",
		PushURL:  true,
		Image:    cmp.InfoIcon(),
	}.Component()
}

func (page NodeListPage) NodeCreateButton() templ.Component {
	return cmp.Button{
		Text:     page.ParentNode.GetChildTypeName(),
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
		return fmt.Sprintf("Up to %s", upNavText(list.ParentNode.GetName()))
	} else {
		return ""
	}
}

func (list NodeListPage) PageTitle() string {
	if list.ParentNode.GetTypeName() != dto.RootTypeName.String() {
		return fmt.Sprintf("%ss for %s: %s", list.ParentNode.GetChildTypeName(), list.ParentNode.GetTypeName(), list.ParentNode.GetName())
	} else {
		return fmt.Sprintf("%ss", list.ParentNode.GetChildTypeName())
	}
}

func (list NodeListPage) ListItemElementID(node ports.Node) ElementID {
	return ElementID(fmt.Sprintf("%s-%d", util.KebabCase(node.GetTypeName()), node.GetID()))
}

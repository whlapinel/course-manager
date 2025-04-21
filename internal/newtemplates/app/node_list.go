package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"gh_static_portfolio/internal/shared/web"
	cmp "gh_static_portfolio/internal/templates/components/base"
	"log"

	"github.com/a-h/templ"

	tpl "gh_static_portfolio/internal/templates/shared"
)

type NodeListPage struct {
	ParentNode       templates.Node
	Children         []templates.Node
	ChildUI          [][]ComponentData
	ChildDetailsURL  web.AddParams // details for child e.g. if listing units, this would be the route handler name to show unit details
	ChildChildrenURL web.AddParams // children of child e.g. if listing units, this would be the route handler name to list unit lessons
	DeleteChildURL   web.AddParams
	ShowNewChildURL  string // e.g. if listing units, this would be the route handler name to show new unit form
	UpNavURL         string
	BreadCrumbsData  BreadCrumbs
}

type NodeDeleteButton struct {
	DeleteChildRHN string
}

func (page NodeListPage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData
}

func (page NodeListPage) DeleteNodeButton(node templates.Node) templ.Component {
	return cmp.Button{
		HxConfirm: fmt.Sprintf("Are you sure you want to delete %s '%s'", node.TypeName(), node.GetName()),
		Method:    cmp.HxDelete,
		URL:       page.DeleteChildURL(node.GetID()),
		HxTarget:  page.ListItemElementID(node).Selector(),
		Image:     cmp.DeleteImage(),
	}.Component()
}

func (page NodeListPage) NodeChildrenButton(node templates.Node) templ.Component {
	log.Println("In NodeChildrenButton: node:", node.GetID(), node.GetName())
	log.Println("Child children URL without params:", page.ChildChildrenURL())
	log.Println("Child children URL without params:", page.ChildChildrenURL(node.GetID()))
	return cmp.Button{
		Text:     node.ChildTypeName() + "s",
		Method:   cmp.HxGet,
		URL:      page.ChildChildrenURL(node.GetID()),
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeListPage) NodeDetailsButton(node templates.Node) templ.Component {
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

func (list NodeListPage) ListItemElementID(node templates.Node) ElementID {
	return ElementID(fmt.Sprintf("%s-%d", tpl.KebabCase(node.TypeName()), node.GetID()))
}

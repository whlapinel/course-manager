package components

import "github.com/a-h/templ"

type PageLayout struct {
	PageTitle string
	Crumbs    BreadCrumbs
	UpNav     UpNav
}

type UpNav struct {
	URL  string
	Text string
}

type BreadCrumbs struct {
	Items []BreadCrumbsItem
}

func (data PageLayout) BreadCrumbs() templ.Component {
	return data.Crumbs.Component()
}

func (data BreadCrumbs) Component() templ.Component {
	return BreadCrumbsComponent(data)
}

type BreadCrumbsItem struct {
	NavItem
}

func (data BreadCrumbsItem) Component() templ.Component {
	data.Class = "text-sm font-medium text-gray-400 hover:text-gray-200"
	return data.NavItem.Component()
}

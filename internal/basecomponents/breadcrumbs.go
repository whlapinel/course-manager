package basecomponents

import "github.com/a-h/templ"

type BreadCrumbs struct {
	Items []BreadCrumbsItem
}

func (data BreadCrumbs) Component() templ.Component {
	return BreadCrumbsComponent(data)
}

type BreadCrumbsItem struct {
	NavItem
}

func (data BreadCrumbsItem) Component() templ.Component {
	data.Class = "text-sm font-medium text-gray-400 hover:text-gray-200 cursor-pointer"
	return data.NavItem.Component()
}

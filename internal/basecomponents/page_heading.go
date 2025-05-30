package basecomponents

import "github.com/a-h/templ"

type PageHeading struct {
	Title string
}

func (data PageHeading) Component() templ.Component {
	return PageHeadingComponent(data)
}

type PageLayout struct {
	PageTitle string
	Crumbs    BreadCrumbs
	UpNav     UpNav
}

type UpNav struct {
	URL  string
	Text string
}

func (data PageLayout) BreadCrumbs() templ.Component {
	return data.Crumbs.Component()
}

func (data PageLayout) PageHeading() templ.Component {
	return PageHeading{
		Title: data.PageTitle,
	}.Component()
}

package templates

import (
	"gh_static_portfolio/internal/shared/node"
	components "gh_static_portfolio/internal/newtemplates/components/base"
	"log"

	"github.com/a-h/templ"
)

type StaticNodeDetailsPage struct {
	PageData
	Path         string
	Node         node.Node
	FilesPageURL string
}

func (page StaticNodeDetailsPage) GetNode() node.Node {
	return page.Node
}

func (page StaticNodeDetailsPage) BreadCrumbs() templ.Component {
	return page.PageData.BreadCrumbs.Component()
}
func (page StaticNodeDetailsPage) Component() templ.Component {
	log.Println("Node:", page.Node.Designation())
	log.Println("FILESPAGE URL", page.FilesPageURL)
	return Layout{
		PageTitle: page.Node.GetName(),
		Head: Head{
			User:      page.User,
			AssetsURL: page.AssetsURL,
		}.Component(),
		Page: StaticNodeDetailsComponent(page),
	}.Component()
}

func (page StaticNodeDetailsPage) Info() templ.Component {
	return nil
}

func (page StaticNodeDetailsPage) Tabs() templ.Component {
	log.Println("FILESPAGE URL", page.FilesPageURL)

	data := components.TabSet{
		Tabs: []components.TabLink{
			{Name: "Files", URL: page.FilesPageURL},
		},
	}
	return data.Component()
}

func (page StaticNodeDetailsPage) Layout() PageData {
	return page.PageData
}

func (page StaticNodeDetailsPage) Filepath() string {
	return page.Path
}

type DetailsPage interface {
	GetNode() node.Node
	Info() templ.Component // for other info that might not apply to all nodes
	Tabs() templ.Component
	BreadCrumbs() templ.Component
}

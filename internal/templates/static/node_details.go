package templates

import (
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"
	"log"

	"github.com/a-h/templ"
)

type StaticNodeDetailsPage struct {
	PageData
	Path         string
	Node         domain.CourseNode
	FilesPageURL string
}

func (page StaticNodeDetailsPage) GetNode() domain.CourseNode {
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
	GetNode() domain.CourseNode
	Tabs() templ.Component
	BreadCrumbs() templ.Component
}

package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
)

type NodeDetailsPage struct {
	Params            CourseIDParams
	ParentNode        domain.CourseNode
	Node              domain.CourseNode
	GetEditNodeURL    string
	PostEditNodeURL   string
	ListChildrenURL   string
	UpNavURL          string
	IsEdit            bool
	Slides            string
	NodeImageURL      func() string
	BreadCrumbsData   BreadCrumbs
	CourseCalendarURL string
}

func (page NodeDetailsPage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData
}

func (page NodeDetailsPage) Component() templ.Component {
	return NodeDetailsComponent(page)
}

func (page NodeDetailsPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: page.PageTitle(),
		UpNav: UpNav{
			URL:  page.UpNavURL,
			Text: page.upNavText(),
		},
	}
}

func (page NodeDetailsPage) CalendarButton() templ.Component {
	return HXButton{
		Text:     "Calendar",
		Method:   HxGet,
		URL:      page.CourseCalendarURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeDetailsPage) PageTitle() string {
	desig := page.Node.Designation()
	if desig != "" {
		return fmt.Sprintf("%s (%s)", desig, page.Node.GetName())
	}
	return page.Node.GetName()

}

func (page NodeDetailsPage) upNavText() string {
	parentPageText := page.Node.TypeName()
	if page.Node.ParentTypeName() == string(domain.RootTypeName) {
		parentPageText = "Home"
		return fmt.Sprintf("Up to %s", parentPageText)
	}
	return fmt.Sprintf("Up to %ss", parentPageText)
}

func (page NodeDetailsPage) ListChildrenButton() templ.Component {
	return HXButton{
		Text:     fmt.Sprintf("%ss", page.Node.ChildTypeName()),
		Method:   HxGet,
		URL:      page.ListChildrenURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

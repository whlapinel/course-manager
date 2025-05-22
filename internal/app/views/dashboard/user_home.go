package dashboardviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/base"

	"github.com/a-h/templ"
)

type UserHomePage struct {
	User            dto.User
	ListTermsURL    string
	GenerateSiteURL string
	SyncSiteURL     string
	StaticSiteURL   string
	ac.BreadCrumbs
	ac.CourseManagerLayout
}

func (page UserHomePage) GenerateSiteButton() templ.Component {
	return cmp.Button{
		Text:     "Generate Site",
		Method:   cmp.HxPost,
		URL:      page.GenerateSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}

func (page UserHomePage) ViewTermsButton() templ.Component {
	return cmp.Button{
		Text:     "View Terms",
		Method:   cmp.HxGet,
		URL:      page.ListTermsURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page UserHomePage) ViewSiteLink() templ.Component {
	return cmp.Link{
		Text:   "Static Site",
		URL:    page.StaticSiteURL,
		Target: cmp.NewTab,
	}.Component()
}

func (page UserHomePage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Home",
		Crumbs:    page.BreadCrumbs.BreadCrumbs(),
	}
}

func (page UserHomePage) Component() templ.Component {
	return UserHomePageComponent(page)
}

func (page UserHomePage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page UserHomePage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.Component2(page.Component())
}

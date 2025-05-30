package dashboardviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/basecomponents"
	"strconv"

	"github.com/a-h/templ"
)

type UserHomePage struct {
	User            dto.User
	Terms           []dto.Term
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
		Attributes: templ.Attributes{
			"hx-include": "#term",
		},
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

func (page UserHomePage) TermSelect() templ.Component {
	var options []cmp.Option
	for _, term := range page.Terms {
		option := cmp.Option{
			Content: term.Name,
			Value:   strconv.Itoa(term.ID),
		}
		options = append(options, option)
	}
	return cmp.NewSelectWithLabel("Term", options).Component()
}

func (page UserHomePage) Component() templ.Component {
	return UserHomePageComponent(page)
}

func (page UserHomePage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page UserHomePage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.WithPage(page.Component())
}

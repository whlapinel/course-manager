package managertemplates

import (
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type UserHomePage struct {
	User            dto.User
	ListTermsURL    string
	GenerateSiteURL string
	SyncSiteURL     string
	BreadCrumbs
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

func (page UserHomePage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Home",
		Crumbs:    page.BreadCrumbs.BreadCrumbs(),
	}

}

func (page UserHomePage) Component() templ.Component {
	return UserHomePageComponent(page)
}

package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type UserHomePage struct {
	User            domain.User
	ListTermsURL    string
	GenerateSiteURL string
	SyncSiteURL     string
}

func (page UserHomePage) GenerateSiteButton() templ.Component {
	return cmp.HXButton{
		Text:     "Generate Site",
		Method:   cmp.HxPost,
		URL:      page.GenerateSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}
func (page UserHomePage) SyncButton() templ.Component {
	return cmp.HXButton{
		Text:     "Sync Site",
		Method:   cmp.HxPost,
		URL:      page.SyncSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}
func (page UserHomePage) ViewTermsButton() templ.Component {
	return cmp.HXButton{
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
		Crumbs:    page.BreadCrumbs().BreadCrumbs(),
	}

}

func (page UserHomePage) Component() templ.Component {
	return UserHomePageComponent(page)
}

func (page UserHomePage) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{
		User: page.User,
	}
}

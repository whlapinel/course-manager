package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
)

type UserHomePage struct {
	User            domain.User
	ListTermsURL    string
	GenerateSiteURL string
	SyncSiteURL     string
}

func (page UserHomePage) GenerateSiteButton() templ.Component {
	return HXButton{
		Text:     "Generate Site",
		Method:   HxPost,
		URL:      page.GenerateSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}
func (page UserHomePage) SyncButton() templ.Component {
	return HXButton{
		Text:     "Sync Site",
		Method:   HxPost,
		URL:      page.SyncSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}
func (page UserHomePage) ViewTermsButton() templ.Component {
	return HXButton{
		Text:     "View Terms",
		Method:   HxGet,
		URL:      page.ListTermsURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page UserHomePage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Home",
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

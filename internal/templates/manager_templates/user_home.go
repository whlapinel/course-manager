package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates/components"

	"github.com/a-h/templ"
)

type UserHomePage struct {
	User            domain.User
	ListTermsURL    string
	GenerateSiteURL string
	SyncSiteURL     string
}

func (page UserHomePage) GenerateSiteButton() templ.Component {
	return components.HXButton{
		Text:     "Generate Site",
		Method:   components.HxPost,
		URL:      page.GenerateSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}
func (page UserHomePage) SyncButton() templ.Component {
	return components.HXButton{
		Text:     "Sync Site",
		Method:   components.HxPost,
		URL:      page.SyncSiteURL,
		HxTarget: "#confirmation",
	}.Component()
}
func (page UserHomePage) ViewTermsButton() templ.Component {
	return components.HXButton{
		Text:     "View Terms",
		Method:   components.HxGet,
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

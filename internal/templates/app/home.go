package managertemplates

import (
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type HomePage struct {
	UsersURL   string
	SigninURL  string
	SignupURL  string
	SignoutURL string
}

func (page HomePage) UsersAuthButton() templ.Component {
	return cmp.HXButton{
		Text:     "Dashboard",
		Method:   cmp.HxGet,
		URL:      page.UsersURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page HomePage) SignupButton() templ.Component {
	return cmp.HXButton{
		Text:     "Sign Up",
		Method:   cmp.HxGet,
		URL:      page.SignupURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}
func (page HomePage) SigninButton() templ.Component {
	return cmp.HXButton{
		Text:     "Sign In",
		Method:   cmp.HxGet,
		URL:      page.SigninURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}
func (page HomePage) SignoutButton() templ.Component {
	return cmp.HXButton{
		Text:     "Sign Out 🛑",
		Method:   cmp.HxPost,
		URL:      page.SignoutURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page HomePage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Home",
	}
}

func (page HomePage) Component() templ.Component {
	return HomePageComponent(page)
}

func (page HomePage) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{}
}

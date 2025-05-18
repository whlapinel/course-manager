package home

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	cmp "gh_static_portfolio/internal/base"

	"github.com/a-h/templ"
)

type HomePage struct {
	UsersURL   string
	SigninURL  string
	SignupURL  string
	SignoutURL string
}

func (page HomePage) UsersAuthButton() templ.Component {
	return cmp.Button{
		Text:     "Dashboard",
		Method:   cmp.HxGet,
		URL:      page.UsersURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page HomePage) SignupButton() templ.Component {
	return cmp.Button{
		Text:     "Sign Up",
		Method:   cmp.HxGet,
		URL:      page.SignupURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}
func (page HomePage) SigninButton() templ.Component {
	return cmp.Button{
		Text:     "Sign In",
		Method:   cmp.HxGet,
		URL:      page.SigninURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}
func (page HomePage) SignoutButton() templ.Component {
	return cmp.Button{
		Text:     "Sign Out",
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

func (page HomePage) BreadCrumbs() appcomponents.BreadCrumbs {
	return appcomponents.BreadCrumbs{}
}
package managertemplates

import "github.com/a-h/templ"

type HomePage struct {
	UsersURL   string
	SigninURL  string
	SignupURL  string
	SignoutURL string
}

func (page HomePage) UsersAuthButton() templ.Component {
	return HXButton{
		Text:     "Dashboard",
		Method:   HxGet,
		URL:      page.UsersURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page HomePage) SignupButton() templ.Component {
	return HXButton{
		Text:     "Sign Up",
		Method:   HxGet,
		URL:      page.SignupURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}
func (page HomePage) SigninButton() templ.Component {
	return HXButton{
		Text:     "Sign In",
		Method:   HxGet,
		URL:      page.SigninURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}
func (page HomePage) SignoutButton() templ.Component {
	return HXButton{
		Text:     "Sign Out 🛑",
		Method:   HxPost,
		URL:      page.SignoutURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page HomePage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Home",
	}
}

func (page HomePage) Component() templ.Component {
	return HomePageComponent(page)
}

func (page HomePage) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{}
}

package appcomponents

import (
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type CourseManagerLayout struct {
	AssetsURLFunc func(...string) string
	HomeURL       string
	SigninURL     string
	SignupURL     string
	SignoutURL    string
	User          dto.User
	Page          templ.Component
}

func (cml CourseManagerLayout) Component2(page templ.Component) templ.Component {
	navItems := []cmp.NavItem{
		cml.NewSignupButton(),
		cml.NewSigninButton(),
		cml.NewSignoutButton(),
	}
	head := HeadComponent()
	layout := cmp.Layout{
		HomeURL:   cml.HomeURL,
		UserImage: cml.User.Picture,
		NavItems:  navItems,
		Head:      head,
		Page:      page,
		UserMenu: cmp.UserMenu{
			Image: cml.User.Picture,
			Links: []cmp.Link{
				{
					Text:   "Sign Out",
					URL:    cml.SignoutURL,
					Target: "#page",
					HTMX:   true,
					Attributes: templ.Attributes{
						string(cmp.HxPost): cml.SignoutURL,
					},
				},
			},
		},
	}
	return cmp.LayoutComponent(layout)
}
func (cml CourseManagerLayout) Component() templ.Component {
	navItems := []cmp.NavItem{
		cml.NewSignupButton(),
		cml.NewSigninButton(),
		cml.NewSignoutButton(),
	}
	head := HeadComponent()
	layout := cmp.Layout{
		HomeURL:   cml.HomeURL,
		UserImage: cml.User.Picture,
		NavItems:  navItems,
		Head:      head,
		Page:      cml.Page,
		UserMenu: cmp.UserMenu{
			Image: cml.User.Picture,
			Links: []cmp.Link{
				{
					Text:   "Sign Out",
					URL:    cml.SignoutURL,
					Target: "#page",
					HTMX:   true,
					Attributes: templ.Attributes{
						string(cmp.HxPost): cml.SignoutURL,
					},
				},
			},
		},
	}
	return cmp.LayoutComponent(layout)
}

func (data CourseManagerLayout) NewSigninButton() cmp.NavItem {
	return cmp.NavItem{
		Text:   "Sign In",
		Method: cmp.HxGet,
		URL:    data.SigninURL,
	}
}

func (data CourseManagerLayout) NewSignupButton() cmp.NavItem {
	return cmp.NavItem{
		Text:   "Sign Up",
		URL:    data.SignupURL,
		Method: cmp.HxGet,
	}
}

func (data CourseManagerLayout) NewSignoutButton() cmp.NavItem {
	return cmp.NavItem{
		Text:      "Sign Out",
		HxConfirm: "Are you sure you want to sign out?",
		Method:    cmp.HxPost,
		URL:       data.SignoutURL,
	}
}

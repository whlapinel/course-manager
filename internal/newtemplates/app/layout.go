package managertemplates

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type CourseManagerLayout struct {
	PageTitle     string
	AssetsURLFunc func(...string) string
	HomeURL       string
	SigninURL     string
	SignupURL     string
	SignoutURL    string
	User          dto.User
	Page          templ.Component
}

func BaseLayout(reverse web.Reverse, page templ.Component, user dto.User) templ.Component {
	cml := CourseManagerLayout{
		HomeURL:    "/",
		Page:       page,
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
	return cml.Component()

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

type BreadCrumbs struct {
	node.Nodes
	UserDetailsURL   string
	TermDetailsURL   string
	CourseDetailsURL string
	UnitDetailsURL   string
	LessonDetailsURL string
}

func (data BreadCrumbs) Component() templ.Component {
	return NewBreadCrumbsComponent(data)
}

func (data BreadCrumbs) BreadCrumbs() cmp.BreadCrumbs {
	var items []cmp.BreadCrumbsItem
	if data.User == nil {
		return cmp.BreadCrumbs{}
	}
	if data.User.GetID() != "" {
		item := cmp.BreadCrumbsItem{
			NavItem: cmp.NavItem{
				Text:   data.User.GetName(),
				URL:    data.UserDetailsURL,
				Method: cmp.HxGet,
			},
		}
		items = append(items, item)

		if data.Term != nil && data.Term.GetID() != 0 {
			item := cmp.BreadCrumbsItem{
				NavItem: cmp.NavItem{
					Text:   data.Term.GetName(),
					URL:    data.TermDetailsURL,
					Method: cmp.HxGet,
				},
			}
			items = append(items, item)

			if data.Course != nil && data.Course.GetID() != 0 {
				item := cmp.BreadCrumbsItem{
					NavItem: cmp.NavItem{
						Text:   data.Course.GetName(),
						URL:    data.CourseDetailsURL,
						Method: cmp.HxGet,
					},
				}
				items = append(items, item)

				if data.Unit != nil && data.Unit.GetID() != 0 {
					item := cmp.BreadCrumbsItem{
						NavItem: cmp.NavItem{
							Text:   data.Unit.Designation(),
							URL:    data.UnitDetailsURL,
							Method: cmp.HxGet,
						},
					}
					items = append(items, item)

					if data.Lesson != nil && data.Lesson.GetID() != 0 {
						item := cmp.BreadCrumbsItem{
							NavItem: cmp.NavItem{
								Text:   data.Lesson.Designation(),
								URL:    data.LessonDetailsURL,
								Method: cmp.HxGet,
							},
						}
						items = append(items, item)

					}
				}
			}
		}
	}
	bc := cmp.BreadCrumbs{
		Items: items,
	}
	return bc
}

func NewBreadCrumbsComponent(data BreadCrumbs) templ.Component {
	bc := data.BreadCrumbs()
	return bc.Component()
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

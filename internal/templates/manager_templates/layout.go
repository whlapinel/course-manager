package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates/components"
	cmp "gh_static_portfolio/internal/templates/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseManagerLayout struct {
	PageTitle  string
	SigninURL  string
	SignupURL  string
	SignoutURL string
	User       domain.User
	Page       templ.Component
	E          *echo.Echo
}

func (cml CourseManagerLayout) Component() templ.Component {
	navItems := []templ.Component{
		cml.NewSignupButton(),
		cml.NewSigninButton(),
		cml.NewSignoutButton(),
	}
	head := HeadComponent()
	layout := components.Layout{
		NavItems: navItems,
		Head:     head,
		Page:     cml.Page,
	}
	return cmp.LayoutComponent(layout)
}

type BreadCrumbs struct {
	User             domain.User
	Term             domain.Term
	Course           domain.Course
	Unit             domain.Unit
	Lesson           domain.Lesson
	UserDetailsURL   string
	TermDetailsURL   string
	CourseDetailsURL string
	UnitDetailsURL   string
	LessonDetailsURL string
}

func (data BreadCrumbs) Component() templ.Component {
	return NewBreadCrumbsComponent(data)
}

func (data BreadCrumbs) BreadCrumbs() components.BreadCrumbs {
	var items []cmp.BreadCrumbsItem
	if data.User.ID != "" {
		item := cmp.BreadCrumbsItem{
			NavItem: cmp.NavItem{
				Text: data.User.GetName(),
				URL:  data.UserDetailsURL,
			},
		}
		items = append(items, item)
		if data.Term.ID != 0 {
			item := cmp.BreadCrumbsItem{
				NavItem: cmp.NavItem{
					Text: data.Term.GetName(),
					URL:  data.TermDetailsURL,
				},
			}
			items = append(items, item)

			if data.Course.ID != 0 {
				item := cmp.BreadCrumbsItem{
					NavItem: cmp.NavItem{
						Text: data.Course.GetName(),
						URL:  data.CourseDetailsURL,
					},
				}
				items = append(items, item)

				if data.Unit.ID != 0 {
					item := cmp.BreadCrumbsItem{
						NavItem: cmp.NavItem{
							Text: data.Unit.Designation(),
							URL:  data.UnitDetailsURL,
						},
					}
					items = append(items, item)

					if data.Lesson.ID != 0 {
						item := cmp.BreadCrumbsItem{
							NavItem: cmp.NavItem{
								Text: data.Lesson.Designation(),
								URL:  data.LessonDetailsURL,
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

func (data CourseManagerLayout) NewSigninButton() templ.Component {
	return cmp.NavItem{
		Text:   "Sign In",
		Method: cmp.HxGet,
		URL:    data.SigninURL,
	}.Component()
}

func (data CourseManagerLayout) NewSignupButton() templ.Component {
	return cmp.NavItem{
		Text:   "Sign Up",
		URL:    data.SignupURL,
		Method: cmp.HxGet,
	}.Component()
}

func (data CourseManagerLayout) NewSignoutButton() templ.Component {
	return cmp.NavItem{
		Text:      "Sign Out",
		HxConfirm: "Are you sure you want to sign out?",
		Method:    cmp.HxPost,
		URL:       data.SignoutURL,
	}.Component()
}

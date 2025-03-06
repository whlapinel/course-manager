package managertemplates

import (
	"gh_static_portfolio/internal/domain"
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
	return CourseManagerLayoutComponent(cml)
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
	return BreadCrumbsComponent(data)
}

func (data CourseManagerLayout) SigninButton() templ.Component {
	return cmp.HXButton{
		Text:     "Sign In",
		Method:   cmp.HxGet,
		URL:      data.SigninURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (data CourseManagerLayout) SignupButton() templ.Component {
	return cmp.HXButton{
		Text:     "Sign Up",
		Method:   cmp.HxGet,
		URL:      data.SignupURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (data CourseManagerLayout) SignoutButton() templ.Component {
	return cmp.HXButton{
		Text:      "Sign Out",
		HxConfirm: "Are you sure you want to sign out?",
		Method:    cmp.HxPost,
		URL:       data.SignoutURL,
		HxTarget:  "#page",
		PushURL:   true,
	}.Component()
}

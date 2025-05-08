package authviews

import (
	appcomponents "gh_static_portfolio/internal/app/components"

	"github.com/a-h/templ"
)

type SigninPage struct {
	GoogleClientID  string
	GoogleSigninURL string
	appcomponents.CourseManagerLayout
}

type SignUpPage struct {
	GoogleClientID  string
	GoogleSignupURL string
}

type SignoutPage struct {
}

func (page SignoutPage) Component() templ.Component {
	return SignoutPageComponent(page)
}

func (page SigninPage) Component() templ.Component {
	return SignInPageComponent(page)
}

func (page SignUpPage) Component() templ.Component {
	return SignupPageComponent(page)
}

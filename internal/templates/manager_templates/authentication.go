package managertemplates

import "github.com/a-h/templ"

type SigninPage struct {
	GoogleClientID  string
	GoogleSigninURL string
}

type SignUpPage struct {
	GoogleClientID  string
	GoogleSignupURL string
}

func (page SigninPage) Component() templ.Component {
	return SignInPageComponent(page)
}

func (page SignUpPage) Component() templ.Component {
	return SignupPageComponent(page)
}

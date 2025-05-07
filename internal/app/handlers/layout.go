package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/a-h/templ"
)

func BaseLayout(reverse web.Reverse, page Page, user dto.User) templ.Component {
	cml := mt.CourseManagerLayout{
		HomeURL:    "/",
		Page:       page.HTMXResponse(),
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
	return cml.Component()
}

func BaseLayout2(reverse web.Reverse, user dto.User) mt.CourseManagerLayout {
	return mt.CourseManagerLayout{
		HomeURL:    "/",
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
}

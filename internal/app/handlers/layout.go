package handlers

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/a-h/templ"
)

func BaseLayout(reverse web.Reverse, page Page, user dto.User) templ.Component {
	cml := appcomponents.CourseManagerLayout{
		HomeURL:    "/",
		Page:       page.HTMXResponse(),
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
	return cml.Component()
}

func BaseLayout2(reverse web.Reverse, user dto.User) appcomponents.CourseManagerLayout {
	return appcomponents.CourseManagerLayout{
		HomeURL:    "/",
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
}
func BaseLayout3(reverse web.Reverse, user dto.User) appcomponents.CourseManagerLayout {
	return appcomponents.CourseManagerLayout{
		HomeURL:    "/",
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
}

func BaseLayoutWithoutUser(reverse web.Reverse) appcomponents.CourseManagerLayout {
	return appcomponents.CourseManagerLayout{
		HomeURL:    "/",
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
}

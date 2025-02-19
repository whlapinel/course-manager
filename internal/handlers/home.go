package handlers

import (
	mt "gh_static_portfolio/internal/templates/manager_templates"

	"github.com/labstack/echo/v4"
)

const (
	Home RouteName = "/"
)

const (
	ShowHome = RouteHandlerName(GET + Home)
)

func (h CourseHandler) HomeHandlers() []RouteHandler {
	return []RouteHandler{
		{Home, ShowHome, GET, h.ShowHome},
	}
}

func (h CourseHandler) ShowHome(c echo.Context) error {
	pageData := mt.HomePage{
		UsersURL:  h.e.Reverse(UserAuth.String()),
		SigninURL: h.e.Reverse(GetSignin.String()),
		SignupURL: h.e.Reverse(GetSignup.String()),
	}
	template := mt.HomePageComponent(pageData)
	layout := mt.CourseManagerLayout{
		PageTitle:  "Home",
		SigninURL:  h.e.Reverse(GetSignin.String()),
		SignupURL:  h.e.Reverse(GetSignup.String()),
		SignoutURL: h.e.Reverse(PostSignout.String()),
		Page:       pageData.Component(),
		E:          h.e,
	}.Component()
	return Respond(c, "", template, layout)
}

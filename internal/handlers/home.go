package handlers

import (
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"

	"github.com/labstack/echo/v4"
)

const (
	Home RoutePath = "/"
)

const (
	ShowHome = RouteHandlerName(GET + Home)
)

type homeRouter struct {
	router
}

// GetRouter implements Router.
func (h *homeRouter) GetRouter() router {
	panic("unimplemented")
}

// SetRouter implements Router.
func (h *homeRouter) SetRouter(router router) {
	panic("unimplemented")
}

func newHomeRouter(svc service.CourseService, app *echo.Echo) homeRouter {
	return homeRouter{
		router: router{
			svc: svc,
			app: app,
		},
	}
}

func (h CourseHandler) HomeHandlers() []RouteHandler {
	return []RouteHandler{
		{Home, ShowHome, GET, h.ShowHome},
	}
}
func HomeHandlers(svc service.CourseService, echo *echo.Echo) []RouteHandler {
	hr := newHomeRouter(svc, echo)
	return []RouteHandler{
		{Home, ShowHome, GET, hr.ShowHome},
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
func (h homeRouter) ShowHome(c echo.Context) error {
	pageData := mt.HomePage{
		UsersURL:  h.app.Reverse(UserAuth.String()),
		SigninURL: h.app.Reverse(GetSignin.String()),
		SignupURL: h.app.Reverse(GetSignup.String()),
	}
	template := mt.HomePageComponent(pageData)
	layout := mt.CourseManagerLayout{
		PageTitle:  "Home",
		SigninURL:  h.app.Reverse(GetSignin.String()),
		SignupURL:  h.app.Reverse(GetSignup.String()),
		SignoutURL: h.app.Reverse(PostSignout.String()),
		Page:       pageData.Component(),
		E:          h.app,
	}.Component()
	return Respond(c, "", template, layout)
}

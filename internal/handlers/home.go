package handlers

import (
	mt "gh_static_portfolio/internal/templates/manager_templates"

	"github.com/labstack/echo/v4"
)

const (
	Home     RouteName = "/"
	Generate RouteName = "/generate"
	Sync     RouteName = "/sync"
)

const (
	ShowHome     = RouteHandlerName(GET + Home)
	GenerateSite = RouteHandlerName(POST + Generate)
	SyncSite     = RouteHandlerName(POST + Sync)
)

func (h CourseHandler) HomeHandlers() []RouteHandler {
	return []RouteHandler{
		{Home, ShowHome, GET, h.ShowHome},
		{Generate, GenerateSite, POST, h.GenerateSite},
		{Sync, SyncSite, POST, h.SyncSite},
	}
}

func (h CourseHandler) ShowHome(c echo.Context) error {
	pageData := mt.HomePage{
		ListTermsURL: h.e.Reverse(ListTerms.String()),
	}
	template := mt.HomePageComponent(pageData)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) GenerateSite(c echo.Context) error {
	h.svc.GenerateSite()
	return Respond(c, "/", mt.Confirm("Site Generation Complete!"), nil)
}

func (h CourseHandler) SyncSite(c echo.Context) error {
	err := h.svc.SyncSite()
	if err != nil {
		return err
	}
	return Respond(c, "/", mt.Confirm("Sync Complete!"), nil)
}

package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/manager_templates"

	"github.com/labstack/echo/v4"
)

const (
	Users    RouteName = "/users"
	User     RouteName = Users + RouteName(UserID)
	Generate RouteName = User + "/generate"
	Sync     RouteName = User + "/sync"
)

const (
	UserAuth     = RouteHandlerName(GET + Users)
	UserHome     = RouteHandlerName(GET + User)
	GenerateSite = RouteHandlerName(POST + Generate)
	SyncSite     = RouteHandlerName(POST + Sync)
)

func (h CourseHandler) UserHomeHandlers() []RouteHandler {
	return []RouteHandler{
		{Users, UserAuth, GET, h.UserAuth},
		{User, UserHome, GET, h.UserHome},
		{Generate, GenerateSite, POST, h.GenerateSite},
		{Sync, SyncSite, POST, h.SyncSite},
	}
}

func (h CourseHandler) UserAuth(c echo.Context) error {
	userID := c.Get("id")
	return c.Redirect(302, h.e.Reverse(UserHome.String(), userID))
}

func (h CourseHandler) UserHome(c echo.Context) error {
	params := ParseCourseIDParams(c)
	userIDParam, ok := params.UserID.Value.(string)
	if !ok {
		return fmt.Errorf("invalid userID")
	}
	if userIDParam != c.Get("id") {
		return fmt.Errorf("mismatch between param userID and authenticated userID")
	}
	page := mt.UserHomePage{
		GenerateSiteURL: h.e.Reverse(GenerateSite.String(), params.ToIntSlice()...),
		SyncSiteURL:     h.e.Reverse(SyncSite.String(), params.ToIntSlice()...),
		ListTermsURL:    h.e.Reverse(ListTerms.String(), params.ToIntSlice()...),
		User: domain.User{
			ID:        userIDParam,
			FirstName: c.Get("first").(string),
			LastName:  c.Get("last").(string),
			Email:     c.Get("email").(string),
			Picture:   c.Get("picture").(string),
		},
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, page.User)
	return Respond(c, "", component, layout)
}
func (h CourseHandler) GenerateSite(c echo.Context) error {
	userID := c.Get("id").(string)
	h.svc.GenerateSite(userID)
	return Respond(c, "/", mt.Confirm("Site Generation Complete!"), nil)
}

func (h CourseHandler) SyncSite(c echo.Context) error {
	err := h.svc.SyncSite()
	if err != nil {
		return err
	}
	return Respond(c, "/", mt.Confirm("Sync Complete!"), nil)
}

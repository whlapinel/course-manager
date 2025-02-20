package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	svc       service.CourseService
	e         *echo.Echo
	params    mt.CourseIDParams
	node      domain.CourseNode
	ancestors []domain.CourseNode
}

// AncestorPath implements NodeHandler.
func (u *userHandler) AncestorPath() []domain.CourseNode {
	panic("unimplemented")
}

// Delete implements NodeHandler.
func (u *userHandler) Delete(echo.Context) error {
	panic("unimplemented")
}

// ListChildren implements NodeHandler.
func (h *userHandler) ListChildren(c echo.Context) error {
	h.params = ParseCourseIDParams(c)
	userID := c.Get("id")
	user, err := h.svc.GetUser(userID.(string))
	if err != nil {
		return err
	}
	log.Println("user ID:", h.params.UserID.Value.(string))
	terms, err := h.svc.GetTerms(userID.(string))
	if err != nil {
		return fmt.Errorf("error in CourseHandler.ListTerms: %s", err)
	}
	user.Terms = terms
	h.node = user
	h.ancestors = []domain.CourseNode{}
	page := mt.TermsListPage{
		ShowTermCalendarRHN: ShowTermCalendar.String(),
		NodeListPage:        NodeListPage(h),
	}
	var component = page.Component()
	layout := CourseManagerLayout(h.Router(), component, user)
	return Respond(c, "", component, layout)

}

// Node implements NodeHandler.
func (u *userHandler) Node() domain.CourseNode {
	panic("unimplemented")
}

// NodeSet implements NodeHandler.
func (u *userHandler) NodeSet() []EmptyNode {
	panic("unimplemented")
}

// Params implements NodeHandler.
func (u *userHandler) Params() mt.CourseIDParams {
	panic("unimplemented")
}

// PostEdit implements NodeHandler.
func (u *userHandler) PostEdit(echo.Context) error {
	panic("unimplemented")
}

// PostNewChild implements NodeHandler.
func (u *userHandler) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// Router implements NodeHandler.
func (u *userHandler) Router() *echo.Echo {
	panic("unimplemented")
}

// ShowDetails implements NodeHandler.
func (u *userHandler) ShowDetails(echo.Context) error {
	panic("unimplemented")
}

// ShowEdit implements NodeHandler.
func (u *userHandler) ShowEdit(echo.Context) error {
	panic("unimplemented")
}

// ShowFiles implements NodeHandler.
func (u *userHandler) ShowFiles(echo.Context) error {
	panic("unimplemented")
}

// ShowNewChild implements NodeHandler.
func (u *userHandler) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeHandler.
func (u *userHandler) ViewFile(echo.Context) error {
	panic("unimplemented")
}

func NewUserHandler(svc service.CourseService, e *echo.Echo) NodeHandler {
	return &userHandler{
		svc: svc,
		e:   e,
	}

}

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

func UserHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	handler := NewUserHandler(svc, router)
	var routeHandlers []RouteHandler
	routeHandlers = append(routeHandlers, NodeHandlers(handler, EmptyNodesTerm...)...)
	return routeHandlers
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
		GenerateSiteURL: h.e.Reverse(GenerateSite.String(), params.ToSlice()...),
		SyncSiteURL:     h.e.Reverse(SyncSite.String(), params.ToSlice()...),
		ListTermsURL:    h.e.Reverse(ListTerms.String(), params.ToSlice()...),
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

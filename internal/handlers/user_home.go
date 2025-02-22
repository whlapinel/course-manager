package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"

	"github.com/labstack/echo/v4"
)

type userRouter struct {
	Router
}

// PostFile implements NodeRouter.
func (r *userRouter) PostFile(echo.Context) error {
	panic("unimplemented")
}

// DownloadFile implements NodeRouter.
func (r *userRouter) DownloadFile(echo.Context) error {
	panic("unimplemented")
}

// Router implements NodeRouter.
func (r *userRouter) GetRouter() Router {
	return r.Router
}

// Delete implements NodeHandler.
func (u *userRouter) Delete(echo.Context) error {
	panic("unimplemented")
}

// ListChildren implements NodeHandler. (implemented)
func (r *userRouter) ListChildren(c echo.Context) error {
	log.Println("userHandler.ListChildren")
	r.params = ParseCourseIDParams(c)
	userID := c.Get("id")
	user, err := r.svc.GetUser(userID.(string))
	if err != nil {
		return err
	}
	log.Println("user ID:", r.params.UserID.Value.(string))
	terms, err := r.svc.GetTerms(userID.(string))
	if err != nil {
		return fmt.Errorf("error in CourseHandler.ListTerms: %s", err)
	}
	user.Terms = terms
	r.node = user
	r.ancestors = []domain.CourseNode{}
	page := mt.TermsListPage{
		ShowTermCalendarRHN: ShowTermCalendar.String(),
		NodeListPage:        NodeListPage(r),
	}
	log.Println("TermsListPage initialized: ", page)
	var component = page.Component()
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

// PostEdit implements NodeHandler.
func (u *userRouter) PostEdit(echo.Context) error {
	panic("unimplemented")
}

// PostNewChild implements NodeHandler.
func (u *userRouter) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// ShowDetails implements NodeHandler.
func (r *userRouter) ShowDetails(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	userIDParam, ok := r.params.UserID.Value.(string)
	if !ok {
		return fmt.Errorf("invalid userID")
	}
	if userIDParam != c.Get("id") {
		return fmt.Errorf("mismatch between param userID and authenticated userID")
	}
	page := mt.UserHomePage{
		GenerateSiteURL: r.app.Reverse(GenerateSite.String(), r.params.ToSlice()...),
		SyncSiteURL:     r.app.Reverse(SyncSite.String(), r.params.ToSlice()...),
		ListTermsURL:    ListChildrenURL(r),
		User: domain.User{
			ID:        userIDParam,
			FirstName: c.Get("first").(string),
			LastName:  c.Get("last").(string),
			Email:     c.Get("email").(string),
			Picture:   c.Get("picture").(string),
		},
	}
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, page.User)
	return Respond(c, "", component, layout)

}

// ShowEdit implements NodeHandler.
func (u *userRouter) ShowEdit(echo.Context) error {
	panic("unimplemented")
}

// ShowFiles implements NodeHandler.
func (u *userRouter) ShowFiles(echo.Context) error {
	panic("unimplemented")
}

// ShowNewChild implements NodeHandler.
func (r *userRouter) ShowNewChild(c echo.Context) error {

	params := ParseCourseIDParams(c)
	user, err := r.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	nodeCreate := NodeCreateChildPage(r)
	template := mt.NodeCreateComponent(nodeCreate)
	layout := CourseManagerLayout(r.app, template, user)
	return Respond(c, "", template, layout)
}

// ViewFile implements NodeHandler.
func (u *userRouter) ViewFile(echo.Context) error {
	panic("unimplemented")
}

func (r *userRouter) UserAuth(c echo.Context) error {
	userID := c.Get("id")
	r.params = mt.CourseIDParams{
		UserID: mt.NodeIDParam{
			Valid: true,
			Value: userID,
		},
	}
	return c.Redirect(302, ShowDetailsURL(r))
}

func (r *userRouter) GenerateSite(c echo.Context) error {
	userID := c.Get("id").(string)
	r.svc.GenerateSite(userID)
	return Respond(c, "/", mt.Confirm("Site Generation Complete!"), nil)
}

func (r *userRouter) SyncSite(c echo.Context) error {
	err := r.svc.SyncSite()
	if err != nil {
		return err
	}
	return Respond(c, "/", mt.Confirm("Sync Complete!"), nil)
}

func NewUserHandler(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &userRouter{
		Router: Router{
			svc:       svc,
			app:       app,
			nodeSet:   EmptyNodesUser,
			ancestors: []domain.CourseNode{domain.RootCourseNode{}},
		},
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

func UserHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nodeRouter := NewUserHandler(svc, router)
	userRouter := nodeRouter.(*userRouter)
	var routeHandlers []RouteHandler
	userRouteHandlers := []RouteHandler{
		{Users, UserAuth, GET, userRouter.UserAuth},
		{Generate, GenerateSite, POST, userRouter.GenerateSite},
		{Sync, SyncSite, POST, userRouter.SyncSite},
	}
	routeHandlers = append(routeHandlers, userRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

package handlers

import (
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"

	"github.com/labstack/echo/v4"
)

type userRouter struct {
	Router
}

// SetRouter implements NodeRouter.
func (r *userRouter) SetRouter(router Router) {
	r.Router = router
}

// PostFile implements NodeRouter.
func (r *userRouter) PostFile(c echo.Context) error {
	return PostFile(c, r)
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
	// userID := c.Get("id")
	path, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = path
	nodes, err := r.svc.NodesWithChildren(path)
	if err != nil {
		return err
	}
	r.nodes = nodes
	page := mt.TermsListPage{
		ShowTermCalendarRHN: ShowTermCalendar.String(),
		NodeListPage:        NodeListPage(r),
	}
	var component = page.Component()
	layout := CourseManagerLayout(r.app, component, nodes.User)
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
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	page := mt.UserHomePage{
		GenerateSiteURL: r.app.Reverse(GenerateSite.String(), r.params.ToSlice()...),
		SyncSiteURL:     r.app.Reverse(SyncSite.String(), r.params.ToSlice()...),
		ListTermsURL:    ListChildrenURL(r),
		User:            r.nodes.User,
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
func (r *userRouter) ShowFiles(c echo.Context) error {
	return ShowFiles(c, r)
}

// ShowNewChild implements NodeHandler.
func (r *userRouter) ShowNewChild(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	user, err := r.svc.GetUser(params.UserID)
	if err != nil {
		return err
	}
	nodeCreate := NodeCreateChildPage(r)
	template := mt.NodeCreatePageComponent(nodeCreate)
	layout := CourseManagerLayout(r.app, template, user)
	return Respond(c, "", template, layout)
}

// ViewFile implements NodeHandler.
func (u *userRouter) ViewFile(echo.Context) error {
	panic("unimplemented")
}

func (r *userRouter) UserAuth(c echo.Context) error {
	userID := c.Get("id")
	r.params.UserID = userID.(string)
	return c.Redirect(302, ShowDetailsURL(r))
}

func (r *userRouter) GenerateSite(c echo.Context) error {
	userID := c.Get("id").(string)
	r.svc.NewGenerateSite(userID)
	return Respond(c, "/", mt.Confirm("Site Generation Complete!"), nil)
}

func NewUserHandler(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &userRouter{
		Router: Router{
			svc:          svc,
			app:          app,
			emptyNodeSet: EmptyNodesUser,
		},
	}
}

const (
	Users    RoutePath = "/users"
	User     RoutePath = Users + RoutePath(UserID)
	Generate RoutePath = User + "/generate"
	Sync     RoutePath = User + "/sync"
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
	}
	routeHandlers = append(routeHandlers, userRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

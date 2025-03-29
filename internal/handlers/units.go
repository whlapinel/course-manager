package handlers

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UnitHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nodeRouter := NewUnitRouter(svc, router)
	var routeHandlers []RouteHandler
	// ur := nodeRouter.(*unitRouter)
	unitRouteHandlers := []RouteHandler{}
	routeHandlers = append(routeHandlers, unitRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

type unitRouter struct {
	router
}

// SetRouter implements NodeRouter.
func (r *unitRouter) SetRouter(router router) {
	r.router = router
}

// Delete implements NodeRouter.
func (r *unitRouter) Delete(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	return r.svc.DeleteUnit(r.params.UnitID)
}

// ListChildren implements NodeRouter.
func (r *unitRouter) ListChildren(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.NodesWithChildren(r.params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	page := NodeListPage(r)
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)
}

// PostEdit implements NodeRouter.
func (r *unitRouter) PostEdit(c echo.Context) error {
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
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "number":
			number, err := strconv.Atoi(val[0])
			if err != nil {
				return err
			}
			r.nodes.Unit.Number = number
		case "description":
			r.nodes.Unit.Description = val[0]
		case "name":
			r.nodes.Unit.Name = val[0]
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	err = r.svc.UpdateUnit(r.nodes.Unit)
	if err != nil {
		return err
	}
	return c.Redirect(303, ShowDetailsURL(r))
}

// PostFile implements NodeRouter.
func (r *unitRouter) PostFile(c echo.Context) error {
	return PostFile(c, r)
}

// PostNewChild implements NodeRouter.
func (r *unitRouter) PostNewChild(c echo.Context) error {
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
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for key, val := range form {
		log.Println("key, val: ", key, val)
	}
	name := c.FormValue("name")
	description := c.FormValue("description")
	numberParam := c.FormValue("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		return err
	}
	lesson, err := r.svc.SaveLesson(service.SaveLessonParams{
		Lesson: domain.Lesson{
			UnitID:      params.UnitID,
			Name:        name,
			Number:      number,
			Description: description,
		},
	})
	if err != nil {
		return err
	}
	return c.Redirect(303, r.app.Reverse(string(ShowNodeDetailsRHN(EmptyNodesLesson...)), AddParams(params, lesson.ID)...))

}

// Router implements NodeRouter.
func (u *unitRouter) GetRouter() router {
	return u.router
}

// ShowDetails implements NodeRouter.
func (r *unitRouter) ShowDetails(c echo.Context) error {
	return ShowDetails(c, r)
}

// ShowEdit implements NodeRouter.
func (r *unitRouter) ShowEdit(c echo.Context) error {
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
	details := NodeDetailsPage(r, true)
	component := details.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}

// ShowFiles implements NodeRouter. (implemented)
func (r *unitRouter) ShowFiles(c echo.Context) error {
	return ShowFiles(c, r)
}

func (r *unitRouter) PostEditFile(c echo.Context) error {
	return PostEditFile(c, r, ShowDetailsURL(r))
}

// ShowNewChild implements NodeRouter.
func (r *unitRouter) ShowNewChild(c echo.Context) error {
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
	page := NodeCreateChildPage(r)
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}
func (r *unitRouter) ShowEditFile(c echo.Context) error {
	return ShowEditFile(c, r, "/")
}

// ViewFile implements NodeRouter.
func (r *unitRouter) ViewFile(c echo.Context) error {
	return ViewFile(c, r, ShowDetailsURL(r))
}

func NewUnitRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &unitRouter{
		router: router{
			svc:          svc,
			app:          app,
			emptyNodeSet: EmptyNodesUnit,
		},
	}

}

const (
	Units            RoutePath = Course + "/units"
	Unit             RoutePath = Units + RoutePath(UnitID)
	NewUnit          RoutePath = Units + "/new"
	EditUnit         RoutePath = Unit + "/edit"
	UnitFiles        RoutePath = Unit + "/files/*"
	UnitViewMarkdown RoutePath = Unit + "/view-markdown/files/*"
)
const (
	ListCourseUnits     = RouteHandlerName(GET + Units)
	UnitDetails         = RouteHandlerName(GET + Unit)
	ShowUnitFiles       = RouteHandlerName(GET + UnitFiles)
	GetUnitViewMarkdown = RouteHandlerName(GET + UnitViewMarkdown)
	PostUnitFile        = RouteHandlerName(POST + UnitFiles)
	ShowEditUnit        = RouteHandlerName(GET + EditUnit)
	PostEditUnit        = RouteHandlerName(POST + EditUnit)
	ShowNewUnit         = RouteHandlerName(GET + NewUnit)
	PostNewUnit         = RouteHandlerName(POST + NewUnit)
	DeleteUnit          = RouteHandlerName(DELETE + Unit)
)

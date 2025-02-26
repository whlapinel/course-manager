package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"gh_static_portfolio/internal/util"
	"log"
	"path/filepath"

	"github.com/a-h/templ"
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
	Router
}

// SetRouter implements NodeRouter.
func (r *unitRouter) SetRouter(router Router) {
	r.Router = router
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
		case "description":
			r.nodes.Unit.Description = val[0]
			err := r.svc.UpdateUnit(r.nodes.Unit)
			if err != nil {
				return err
			}
		case "name":
			r.nodes.Unit.Name = val[0]
			err := r.svc.UpdateUnit(r.nodes.Unit)
			if err != nil {
				return err
			}
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	return c.Redirect(303, ShowDetailsURL(r))
}

// PostFile implements NodeRouter.
func (r *unitRouter) PostFile(c echo.Context) error {
	return PostFile(c, r)
}

// PostNewChild implements NodeRouter.
func (u *unitRouter) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// Router implements NodeRouter.
func (u *unitRouter) GetRouter() Router {
	return u.Router
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

	queryParam := c.QueryParam("field")
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := NodeDetailsPage(r, true)
	var component templ.Component
	if queryParam == util.KebabCase(string(Description)) {
		component = mt.EditDescriptionComponent(details)
	} else if queryParam == util.KebabCase(string(Name)) {
		component = mt.EditNameComponent(details)
	} else if queryParam == util.KebabCase(string(Number)) {
		component = mt.EditNumberComponent(details)
	} else {
		errText := "field value is not expected"
		log.Println(errText)
		return fmt.Errorf("%s %s", errText, queryParam)
	}
	layout := CourseManagerLayout(r.app, details.Component(), r.nodes.User)
	return Respond(c, "", component, layout)

}

// ShowFiles implements NodeRouter. (implemented)
func (r *unitRouter) ShowFiles(c echo.Context) error {
	return ShowFiles(c, r)
}

// ShowNewChild implements NodeRouter.
func (u *unitRouter) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeRouter.
func (r *unitRouter) ViewFile(c echo.Context) error {
	path := c.Param("*")
	log.Println("path: ", path)
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
	err = r.svc.CreateNodeFilesDir(r.nodes.ToSlice()...)
	if err != nil {
		return err
	}
	pathRoot := data.NodeFilesDirPath(r.nodes.ToSlice()...)
	path = filepath.Join(pathRoot, path)
	content, err := RenderMarkdownFile(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	data := mt.MarkdownDocument{
		Title:   filepath.Base(path),
		Content: string(content),
		Static:  false,
	}
	err = mt.DocLayout(data).Render(context.Background(), &buf)
	if err != nil {
		return err
	}
	data.Content = buf.String()
	component := mt.MarkdownIFrame(data)
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}

func NewUnitRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &unitRouter{
		Router: Router{
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

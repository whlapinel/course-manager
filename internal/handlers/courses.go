package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"
	templates "gh_static_portfolio/internal/templates/shared"
	"log"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func CourseHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nodeRouter := NewCourseRouter(svc, router)
	var routeHandlers []RouteHandler
	cr := nodeRouter.(*courseRouter)
	courseRouteHandlers := []RouteHandler{
		{CopyCourse, GetCopyCourse, GET, cr.GetCopyCourse},
		{CopyCourseToTerm, PostCopyCourseToTerm, POST, cr.PostCopyCourseToTerm},
		{StandardSet, PostSelectStandardSet, POST, cr.PostSelectStandardSet},
	}
	routeHandlers = append(routeHandlers, courseRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

type courseRouter struct {
	Router
}

// SetRouter implements NodeRouter.
func (r *courseRouter) SetRouter(router Router) {
	r.Router = router
}

// PostFile implements NodeRouter.
func (r *courseRouter) PostFile(c echo.Context) error {
	return PostFile(c, r)

}

func (r *courseRouter) GetRouter() Router {
	return r.Router
}

// Delete implements NodeRouter.
func (r *courseRouter) Delete(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	return r.svc.DeleteCourse(nodes.Course.ID)
}

// ListChildren implements NodeRouter. (implemented)
func (r *courseRouter) ListChildren(c echo.Context) error {
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
	layout := CourseManagerLayout(r.app, component, nodes.User)
	return Respond(c, "", component, layout)
}

// PostEdit implements NodeRouter.
func (c *courseRouter) PostEdit(echo.Context) error {
	panic("unimplemented")
}

// PostNewChild implements NodeRouter.
func (c *courseRouter) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// ShowDetails implements NodeRouter.
func (r *courseRouter) ShowDetails(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	sets, err := r.svc.GetStandardSets()
	if err != nil {
		return err
	}
	page := mt.CourseDetailsPage{
		GetCopyCourseURL:         r.app.Reverse(GetCopyCourse.String(), r.params.ToSlice()...),
		StandardSets:             sets,
		PostSelectStandardSetURL: r.app.Reverse(string(PostSelectStandardSet), r.params.ToSlice()...),
		NodeDetailsPage: mt.NodeDetailsPage{
			Params:            r.params,
			Node:              nodes.Course,
			CourseCalendarURL: r.app.Reverse(ShowCourseCalendar.String(), r.params.ToSlice()...),
			GetEditNodeURL:    r.app.Reverse(ShowEditCourse.String(), r.params.ToSlice()...),
			PostEditNodeURL:   r.app.Reverse(PostEditCourse.String(), r.params.ToSlice()...),
			ListChildrenURL:   r.app.Reverse(ListCourseUnits.String(), r.params.ToSlice()...),
			UpNavURL:          r.app.Reverse(ListTermCourses.String(), r.params.ToSlice()...),
			IsEdit:            false,
			BreadCrumbsData:   BreadCrumbs(r.app, r.params, nodes.ToSlice()...),
			ServerFilesURL:    r.app.Reverse(ShowCourseFiles.String(), AddParams(params, "")...),
		},
	}
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, nodes.User)
	return Respond(c, "", component, layout)

}

// ShowEdit implements NodeRouter.
func (r *courseRouter) ShowEdit(c echo.Context) error {
	queryParam := c.QueryParam("field")
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := NodeDetailsPage(r, true)
	respond := func(component templ.Component) error {
		return Respond(c, r.app.Reverse(string(CourseDetails), r.params.ToSlice()...), component, nil)
	}
	if queryParam == templates.KebabCase(string(Description)) {
		return respond(mt.EditDescriptionComponent(details))
	} else if queryParam == templates.KebabCase(string(Name)) {
		return respond(mt.EditNameComponent(details))
	}
	errText := "field value is not expected"
	log.Println(errText)
	return fmt.Errorf("%s %s", errText, queryParam)

}

// ShowFiles implements NodeRouter.
func (r *courseRouter) ShowFiles(c echo.Context) error {
	return ShowFiles(c, r)
}

// ShowNewChild implements NodeRouter.
func (c *courseRouter) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeRouter.
func (r *courseRouter) ViewFile(c echo.Context) error {
	return ViewFile(c, r, ShowDetailsURL(r))
}

func NewCourseRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &courseRouter{
		Router: Router{
			svc:          svc,
			app:          app,
			emptyNodeSet: EmptyNodesCourse,
		},
	}
}

const (
	Courses            RoutePath = Term + "/courses"
	Course             RoutePath = Courses + RoutePath(CourseID)
	CourseFiles        RoutePath = Course + "/files/*"
	CourseViewMarkdown RoutePath = Course + "/view-markdown/files/*"
	CourseImage        RoutePath = Course + "/image"
	NewCourse          RoutePath = Courses + "/new"
	EditCourse         RoutePath = Course + "/edit"
	CopyCourse         RoutePath = Course + "/copy-to-term"
	CopyCourseToTerm   RoutePath = CopyCourse
	StandardSet        RoutePath = Course + "/standard-set"
)
const (
	ListTermCourses       = RouteHandlerName(GET + Courses)
	CourseDetails         = RouteHandlerName(GET + Course)
	ShowCourseFiles       = RouteHandlerName(GET + CourseFiles)
	GetCourseViewMarkdown = RouteHandlerName(GET + CourseViewMarkdown)
	PostCourseFiles       = RouteHandlerName(POST + CourseFiles)
	ShowEditCourse        = RouteHandlerName(GET + EditCourse)
	PostEditCourse        = RouteHandlerName(POST + EditCourse)
	ShowNewCourse         = RouteHandlerName(GET + NewCourse)
	PostNewCourse         = RouteHandlerName(POST + NewCourse)
	DeleteCourse          = RouteHandlerName(DELETE + Course)
	GetCopyCourse         = RouteHandlerName(GET + CopyCourse)
	PostCopyCourseToTerm  = RouteHandlerName(POST + CopyCourseToTerm)
	PostSelectStandardSet = RouteHandlerName(POST + StandardSet)
)

func (r *courseRouter) GetCopyCourse(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	terms, err := r.svc.GetTerms(nodes.User.ID)
	if err != nil {
		return err
	}
	data := mt.CopyCourseData{
		TermID:                  r.params.TermID,
		CourseID:                r.params.CourseID,
		Terms:                   terms,
		E:                       r.app,
		PostCopyCourseToTermRHN: string(PostCopyCourseToTerm),
	}
	component := data.Component()
	return Respond(c, r.app.Reverse(ListTermCourses.String(), r.params.ToSlice()...), component, nil)
}

func (r *courseRouter) PostCopyCourseToTerm(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	if r.params.CourseID != 0 && r.params.TermID != 0 {
		err := c.Request().ParseForm()
		if err != nil {
			return err
		}
		termIDParam := c.Request().Form.Get("term-id")
		termID, err := strconv.Atoi(termIDParam)
		if err != nil {
			return err
		}
		_, err = r.svc.CopyCourseToTerm(r.params.CourseID, termID)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("params not valid: courseID: %d and termID: %d", r.params.CourseID, r.params.TermID)
	}
	return c.Redirect(302, r.app.Reverse(ListTermCourses.String(), r.params.ToSlice()...))
}

func (r *courseRouter) PostSelectStandardSet(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	standardSetParam := c.Request().Form.Get("standard-set")
	setID, err := strconv.Atoi(standardSetParam)
	if err != nil {
		return err
	}
	log.Println("selected set: ", standardSetParam)
	err = r.svc.SetStandardSet(r.params.CourseID, setID)
	if err != nil {
		return err
	}
	return c.Redirect(302, r.app.Reverse(CourseDetails.String(), r.params.ToSlice()...))
}

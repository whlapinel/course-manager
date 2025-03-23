package handlers

import (
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type termRouter struct {
	router
}

// SetRouter implements NodeRouter.
func (r *termRouter) SetRouter(router router) {
	r.router = router
}

// PostFile implements NodeRouter.
func (r *termRouter) PostFile(c echo.Context) error {
	return PostFile(c, r)
}

func NewTermRouter(svc service.CourseService, e *echo.Echo) NodeRouter {
	return &termRouter{
		router: router{
			svc:          svc,
			app:          e,
			emptyNodeSet: EmptyNodesTerm,
		},
	}

}

func (r *termRouter) GetRouter() router {
	return r.router
}

// PostNewChild implements NodeRouter. (implemented)
func (r *termRouter) PostNewChild(c echo.Context) error {
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
	course, err := r.svc.SaveCourse(service.SaveCourseParams{
		TermID:      params.TermID,
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}
	return c.Redirect(303, r.app.Reverse(string(ShowNodeDetailsRHN(EmptyNodesCourse...)), AddParams(params, course.ID)...))
}

// ShowNewChild implements NodeRouter. (implemented)
func (r *termRouter) ShowNewChild(c echo.Context) error {
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
	nodeCreate := NodeCreateChildPage(r)
	component := nodeCreate.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}

// Delete implements NodeRouter. (implemented)
func (r *termRouter) Delete(c echo.Context) error {
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
	err = r.svc.DeleteTerm(r.params.TermID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// ListChildren implements NodeRouter. (implemented)
func (r *termRouter) ListChildren(c echo.Context) error {
	newParams, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = newParams
	nodes, err := r.svc.NodesWithChildren(r.params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	page := mt.CourseListPage{
		ShowAssessmentsRHN: string(GetCourseAssessments),
		ShowCalendarRHN:    string(ShowCourseCalendar),
		NodeListPage:       NodeListPage(r),
	}
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)
}

// PostEdit implements NodeRouter. (implemented)
func (r *termRouter) PostEdit(c echo.Context) error {
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
			r.nodes.Term.Description = val[0]
			err := r.svc.UpdateTerm(r.nodes.Term)
			if err != nil {
				return err
			}
		case "name":
			r.nodes.Term.Name = val[0]
			err := r.svc.UpdateTerm(r.nodes.Term)
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

// ShowDetails implements NodeRouter. (implemented)
func (r *termRouter) ShowDetails(c echo.Context) error {
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

	pageData := mt.TermDetailsPage{
		NodeDetailsPage:      NodeDetailsPage(r, false),
		ShowEditTermDatesURL: r.app.Reverse(ShowEditTermDates.String(), r.params.ToSlice()...),
	}
	component := pageData.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)
}

// ShowEdit implements NodeRouter.  (implemented)
func (r *termRouter) ShowEdit(c echo.Context) error {
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
	details := mt.TermDetailsPage{
		NodeDetailsPage: NodeDetailsPage(r, true),
	}
	component := details.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)
}

// ShowFiles implements NodeRouter.
func (r *termRouter) ShowFiles(c echo.Context) error {
	return ShowFiles(c, r)
}

// ViewFile implements NodeRouter.
func (r *termRouter) ViewFile(c echo.Context) error {
	return ViewFile(c, r, ShowDetailsURL(r))
}

const (
	Terms        RoutePath = User + "/terms"
	Term         RoutePath = Terms + RoutePath(TermID)
	TermCalendar RoutePath = Term + "/calendar"
	Occasions    RoutePath = Term + "/occasions"
	Occasion     RoutePath = Occasions + RoutePath(OccasionID)
	TermDates    RoutePath = Term + "/dates"
	EditTerm     RoutePath = Term + "/edit"
	NewTerm      RoutePath = Terms + "/new"
)

const (
	// ListTerms         = RouteHandlerName(GET + Terms)
	TermDetails       = RouteHandlerName(GET + Term)
	ShowTermCalendar  = RouteHandlerName(GET + TermCalendar)
	CreateOccasion    = RouteHandlerName(POST + Occasions)
	DeleteOccasion    = RouteHandlerName(DELETE + Occasion)
	ShowEditOccasion  = RouteHandlerName(GET + Occasion)
	PostEditOccasion  = RouteHandlerName(POST + Occasion)
	ShowEditTermDates = RouteHandlerName(GET + TermDates)
	PostEditTermDates = RouteHandlerName(POST + TermDates)
	ShowEditTerm      = RouteHandlerName(GET + EditTerm)
	PostEditTerm      = RouteHandlerName(POST + EditTerm)
	ShowNewTerm       = RouteHandlerName(GET + NewTerm)
	PostNewTerm       = RouteHandlerName(POST + NewTerm)
	DeleteTerm        = RouteHandlerName(POST + Term)
)

func TermHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nodeRouter := NewTermRouter(svc, router)
	var routeHandlers []RouteHandler
	termRouter := nodeRouter.(*termRouter)
	termRouteHandlers := []RouteHandler{
		{TermCalendar, ShowTermCalendar, GET, termRouter.ShowTermCalendar},
		{Occasions, CreateOccasion, POST, termRouter.CreateOccasion},
		{Occasion, ShowEditOccasion, GET, termRouter.ShowEditOccasion},
		{Occasion, PostEditOccasion, POST, termRouter.PostEditOccasion},
		{Occasion, DeleteOccasion, DELETE, termRouter.DeleteOccasion},
		{TermDates, ShowEditTermDates, GET, termRouter.ShowEditTermDates},
		{TermDates, PostEditTermDates, POST, termRouter.PostEditTermDates},
	}
	routeHandlers = append(routeHandlers, termRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

func (r *termRouter) ShowTermCalendar(c echo.Context) error {
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
	termsListURL := ListSiblingsURL(r)
	log.Println("termsListURL:", termsListURL)
	data := mt.TermCalendar{
		Params:              r.params,
		Term:                r.nodes.Term,
		GetEditOccasionRHN:  ShowEditOccasion.String(),
		PostEditOccasionRHN: PostEditOccasion.String(),
		DeleteOccasionRHN:   DeleteOccasion.String(),
		ListTermsURL:        ListSiblingsURL(r),
		TermDetailsURL:      ShowDetailsURL(r),
		CreateOccasionURL:   r.app.Reverse(CreateOccasion.String(), r.params.ToSlice()...),
		E:                   r.app,
	}

	component := data.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}

func (r *termRouter) CreateOccasion(c echo.Context) error {
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

	name := form.Get("name")
	dateParam := form.Get("date")
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	log.Println("name and date:", name, dateParam)
	occasion, err := r.svc.CreateOccasion(date, name, r.params.TermID)
	if err != nil {
		return err
	}
	log.Println(occasion)
	return c.Redirect(303, r.app.Reverse(ShowTermCalendar.String(), r.params.ToSlice()...))
}

func (r *termRouter) ShowEditOccasion(c echo.Context) error {
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
	occasionID, err := ParseRouteParam(c, OccasionID)
	if err != nil {
		return err
	}
	occasion, err := r.svc.GetOccasion(occasionID)
	if err != nil {
		return err
	}
	component := mt.TermOccasionEditor{
		Occasion:            occasion,
		IsEditing:           true,
		GetEditOccasionURL:  r.app.Reverse(ShowEditOccasion.String(), AddParams(r.params, occasion.ID)...),
		PostEditOccasionURL: r.app.Reverse(PostEditOccasion.String(), AddParams(r.params, occasion.ID)...),
	}.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}

func (r *termRouter) PostEditOccasion(c echo.Context) error {
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
	occasionID, err := ParseRouteParam(c, OccasionID)
	if err != nil {
		return err
	}
	form := c.Request().Form
	name := form.Get("name")
	err = r.svc.UpdateOccasion(name, occasionID)
	if err != nil {
		return err
	}
	return c.Redirect(303, r.app.Reverse(ShowTermCalendar.String(), r.params.ToSlice()...))
}

func (r *termRouter) DeleteOccasion(c echo.Context) error {
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
	occasionID, err := ParseRouteParam(c, OccasionID)
	if err != nil {
		return err
	}
	err = r.svc.DeleteOccasion(occasionID)
	if err != nil {
		return err
	}
	return c.NoContent(200)

}

func (r *termRouter) ShowEditTermDates(c echo.Context) error {
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

	if r.nodes.Term.InstructionalDays == nil {
		log.Println("term.InstructionalDays is nil")
	}
	for _, d := range r.nodes.Term.InstructionalDays {
		log.Println("date: ", d.Format(time.DateOnly))
	}
	pageData := mt.AddNonInstructDayPage{
		Term:           r.nodes.Term,
		GetAddDayURL:   r.app.Reverse(ShowEditTermDates.String(), r.params.TermID),
		PostAddDayURL:  r.app.Reverse(PostEditTermDates.String(), r.params.TermID),
		TermDetailsURL: ShowDetailsURL(r),
	}
	template := pageData.Component()
	layout := CourseManagerLayout(r.app, template, r.nodes.User)
	return Respond(c, "", template, layout)
}

func (r *termRouter) PostEditTermDates(c echo.Context) error {
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
	formDate := form.Get("date")
	date, err := time.Parse(time.DateOnly, formDate)
	if err != nil {
		return err
	}
	err = r.svc.AddNonInstructDay(params.TermID, date)
	if err != nil {
		return err
	}
	return c.Redirect(303, ShowDetailsURL(r))
}

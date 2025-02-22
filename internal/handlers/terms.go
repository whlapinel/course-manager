package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"gh_static_portfolio/internal/templates"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type termRouter struct {
	Router
}

// PostFile implements NodeRouter.
func (r *termRouter) PostFile(echo.Context) error {
	panic("unimplemented")
}

func NewTermRouter(svc service.CourseService, e *echo.Echo) NodeRouter {
	return &termRouter{
		Router: Router{
			svc:     svc,
			app:     e,
			nodeSet: EmptyNodesTerm,
		},
	}

}

func (r *termRouter) GetRouter() Router {
	return r.Router
}

// PostNewChild implements NodeRouter. (implemented)
func (r *termRouter) PostNewChild(c echo.Context) error {
	params := ParseCourseIDParams(c)
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for key, val := range form {
		log.Println("key, val: ", key, val)
	}
	termID := params.TermID.Value
	name := c.FormValue("name")
	description := c.FormValue("description")
	course, err := r.svc.SaveCourse(service.SaveCourseParams{
		TermID:      termID.(int),
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}
	return c.Redirect(303, r.app.Reverse(string(ShowNodeDetailsRHN(EmptyNodesCourse...)), params.ToSlice(course.ID)...))
}

// ShowNewChild implements NodeRouter. (implemented)
func (r *termRouter) ShowNewChild(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := r.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	r.node = term
	nodeCreate := NodeCreateChildPage(r)
	// nodeCreate := mt.NodeCreatePage{
	// 	ParentNode:        h.Node(),
	// 	NodeType:          domain.NodeTypeName(h.node.ChildTypeName()),
	// 	Params:            h.Params(),
	// 	PostCreateNodeURL: PostNewChildURL(h),
	// 	CancelURL:         ListChildrenURL(h),
	// 	BreadCrumbsData:   BreadCrumbs(h.e, params, user, term),
	// }
	component := mt.NodeCreateComponent(nodeCreate)
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

// Delete implements NodeRouter. (implemented)
func (r *termRouter) Delete(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	err = r.svc.DeleteTerm(termID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// ListChildren implements NodeRouter. (implemented)
func (r *termRouter) ListChildren(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		log.Println("error in GetUser:", err)
		return err
	}
	termID := r.params.TermID.Value.(int)
	term, err := r.svc.GetTerm(termID)
	if err != nil {
		log.Println("error in GetTerm:", err)
		return err
	}
	courses, err := r.svc.GetCourses(termID)
	if err != nil {
		log.Println("error in GetCourses:", err)
		return err
	}
	term.Courses = courses
	r.node = term
	r.ancestors = []domain.CourseNode{user, term}
	page := mt.CourseListPage{
		ShowAssessmentsRHN: string(GetCourseAssessments),
		ShowCalendarRHN:    string(ShowCourseCalendar),
		NodeListPage:       NodeListPage(r),
	}
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)
}

// PostEdit implements NodeRouter. (implemented)
func (r *termRouter) PostEdit(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(r.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	r.node = term
	r.ancestors = []domain.CourseNode{user}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	var updateTerm = func(term domain.Term) (domain.Term, error) {
		log.Println("updating: ", term.ID, term.Name, term.Description)
		err := r.svc.UpdateTerm(term)
		if err != nil {
			return domain.Term{}, err
		}
		updatedTerm, err := r.svc.GetTerm(r.params.TermID.Value.(int))
		if err != nil {
			return domain.Term{}, err
		}
		log.Println("retrieved term: ", updatedTerm.ID, updatedTerm.Name, updatedTerm.Description)
		return updatedTerm, nil
	}
	var pageData = func() mt.NodeDetailsPage {
		return NodeDetailsPage(r, false)
	}
	var template templ.Component
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "description":
			term.Description = val[0]
			term, err := updateTerm(term)
			if err != nil {
				return err
			}
			r.node = term
			details := pageData()
			template = mt.EditDescriptionComponent(details)
		case "name":
			term.Name = val[0]
			term, err := updateTerm(term)
			if err != nil {
				return err
			}
			r.node = term
			details := pageData()
			template = mt.EditNameComponent(details)
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	if template == nil {
		panic("template is nil!")
	}
	return Respond(c, ShowDetailsURL(r), template, nil)

}

// ShowDetails implements NodeRouter. (implemented)
func (r *termRouter) ShowDetails(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(r.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	r.node = term
	r.ancestors = []domain.CourseNode{user, term}
	pageData := mt.TermDetailsPage{
		NodeDetailsPage:      NodeDetailsPage(r, false),
		ShowEditTermDatesURL: r.app.Reverse(ShowEditTermDates.String(), r.params.ToSlice()...),
	}
	component := pageData.Component()
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)
}

// ShowEdit implements NodeRouter.  (implemented)
func (r *termRouter) ShowEdit(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	queryParam := c.QueryParam("field")
	term, err := r.svc.GetTerm(r.params.TermID.Value.(int))
	if err != nil {
		log.Println(err)
		return err
	}
	r.node = term
	r.ancestors = []domain.CourseNode{user}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := mt.TermDetailsPage{
		NodeDetailsPage: NodeDetailsPage(r, true),
	}

	respond := func(component templ.Component) error {
		return Respond(c, ShowDetailsURL(r), component, nil)
	}
	if queryParam == templates.KebabCase(string(Description)) {
		return respond(mt.EditDescriptionComponent(details.NodeDetailsPage))
	} else if queryParam == templates.KebabCase(string(Name)) {
		return respond(mt.EditNameComponent(details.NodeDetailsPage))
	}
	errText := "field value is not expected"
	log.Println(errText)
	return fmt.Errorf("%s %s", errText, queryParam)

}

// ShowFiles implements NodeRouter.
func (h *termRouter) ShowFiles(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeRouter.
func (h *termRouter) ViewFile(echo.Context) error {
	panic("unimplemented")
}

const (
	Terms        RouteName = User + "/terms"
	Term         RouteName = Terms + RouteName(TermID)
	TermCalendar RouteName = Term + "/calendar"
	Occasions    RouteName = Term + "/occasions"
	Occasion     RouteName = Occasions + RouteName(OccasionID)
	TermDates    RouteName = Term + "/dates"
	EditTerm     RouteName = Term + "/edit"
	NewTerm      RouteName = Terms + "/new"
)

const (
	// ListTerms         = RouteHandlerName(GET + Terms)
	TermDetails       = RouteHandlerName(GET + Term)
	ShowTermCalendar  = RouteHandlerName(GET + TermCalendar)
	CreateOccasion    = RouteHandlerName(POST + Occasions)
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
		{TermDates, ShowEditTermDates, GET, termRouter.ShowEditTermDates},
		{TermDates, PostEditTermDates, POST, termRouter.PostEditOccasion},
	}
	routeHandlers = append(routeHandlers, termRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

func (r *termRouter) ShowTermCalendar(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	log.Println("len params: ", len(r.params.ToSlice()))
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}

	log.Println("termID", r.params.TermID.Value)
	term, err := r.svc.GetTerm(r.params.TermID.Value.(int))
	if err != nil {
		log.Println(err)
		return err
	}
	termsListURL := ListSiblingsURL(r)
	log.Println("termsListURL:", termsListURL)
	data := mt.TermCalendar{
		Params:              r.params,
		Term:                term,
		GetEditOccasionRHN:  ShowEditOccasion.String(),
		PostEditOccasionRHN: PostEditOccasion.String(),
		ListTermsURL:        ListSiblingsURL(r),
		TermDetailsURL:      ShowDetailsURL(r),
		CreateOccasionURL:   r.app.Reverse(CreateOccasion.String(), r.params.ToSlice()...),
		E:                   r.app,
	}

	component := data.Component()
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

func (r *termRouter) CreateOccasion(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	err := c.Request().ParseForm()
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
	occasion, err := r.svc.CreateOccasion(date, name, r.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	log.Println(occasion)
	return c.Redirect(303, r.app.Reverse(ShowTermCalendar.String(), r.params.ToSlice()...))
}

func (r *termRouter) ShowEditOccasion(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}

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
		GetEditOccasionURL:  r.app.Reverse(ShowEditOccasion.String(), r.params.ToSlice(occasion.ID)...),
		PostEditOccasionURL: r.app.Reverse(PostEditOccasion.String(), r.params.ToSlice(occasion.ID)...),
	}.Component()
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

func (r *termRouter) PostEditOccasion(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	err := c.Request().ParseForm()
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

func (r *termRouter) ShowEditTermDates(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}

	termID, err := TermIDParam(r.params)
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	if term.InstructionalDays == nil {
		log.Println("term.InstructionalDays is nil")
	}
	for _, d := range term.InstructionalDays {
		log.Println("date: ", d.Format(time.DateOnly))
	}
	pageData := mt.AddNonInstructDayPage{
		Term:           term,
		GetAddDayURL:   r.app.Reverse(ShowEditTermDates.String(), termID),
		PostAddDayURL:  r.app.Reverse(PostEditTermDates.String(), termID),
		TermDetailsURL: ShowDetailsURL(r),
	}
	template := pageData.Component()
	layout := CourseManagerLayout(r.app, template, user)
	return Respond(c, "", template, layout)
}

func (r *termRouter) PostEditTermDates(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
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
	err = r.svc.AddNonInstructDay(termID, date)
	if err != nil {
		return err
	}
	return c.Redirect(303, ShowDetailsURL(r))
}

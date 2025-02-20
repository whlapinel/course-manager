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

type termHandler struct {
	svc       service.CourseService
	e         *echo.Echo
	params    mt.CourseIDParams
	node      domain.CourseNode
	ancestors []domain.CourseNode
}

func NewTermHandler(svc service.CourseService, e *echo.Echo) NodeHandler {
	return &termHandler{
		svc: svc,
		e:   e,
	}

}

func (h *termHandler) AncestorPath() []domain.CourseNode {
	return h.ancestors
}

func (h *termHandler) Node() domain.CourseNode {
	return h.node
}

func (h *termHandler) NodeSet() []EmptyNode {
	return EmptyNodesTerm
}

func (h *termHandler) Router() *echo.Echo {
	return h.e
}

func (h *termHandler) Params() mt.CourseIDParams {
	return h.params
}

// PostNewChild implements NodeHandler.
func (h *termHandler) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// ShowNewChild implements NodeHandler.
func (h *termHandler) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// Delete implements NodeHandler.
func (h *termHandler) Delete(echo.Context) error {
	panic("unimplemented")
}

// ListChildren implements NodeHandler.
func (h *termHandler) ListChildren(c echo.Context) error {
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	termID := h.params.TermID.Value.(int)
	term, err := h.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	courses, err := h.svc.GetCourses(termID)
	if err != nil {
		return err
	}
	term.Courses = courses
	h.node = term
	h.ancestors = []domain.CourseNode{user}
	page := mt.CourseListPage{
		ShowAssessmentsRHN: string(GetCourseAssessments),
		ShowCalendarRHN:    string(ShowCourseCalendar),
		NodeListPage:       NodeListPage(h),
	}
	component := page.Component()
	layout := CourseManagerLayout(h.e, component, user)
	return Respond(c, "", component, layout)
}

// NewChild implements NodeHandler.
func (h *termHandler) NewChild(echo.Context) error {
	panic("unimplemented")
}

// PostEdit implements NodeHandler.
func (h *termHandler) PostEdit(c echo.Context) error {
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(h.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	h.node = term
	h.ancestors = []domain.CourseNode{user}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	var updateTerm = func(term domain.Term) (domain.Term, error) {
		log.Println("updating: ", term.ID, term.Name, term.Description)
		err := h.svc.UpdateTerm(term)
		if err != nil {
			return domain.Term{}, err
		}
		updatedTerm, err := h.svc.GetTerm(h.params.TermID.Value.(int))
		if err != nil {
			return domain.Term{}, err
		}
		log.Println("retrieved term: ", updatedTerm.ID, updatedTerm.Name, updatedTerm.Description)
		return updatedTerm, nil
	}
	var pageData = func() mt.NodeDetailsPage {
		return NodeDetailsPage(h, false)
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
			h.node = term
			details := pageData()
			template = mt.EditDescriptionComponent(details)
		case "name":
			term.Name = val[0]
			term, err := updateTerm(term)
			if err != nil {
				return err
			}
			h.node = term
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
	return Respond(c, ShowDetailsURL(h), template, nil)

}

// ShowDetails implements NodeHandler.
func (h *termHandler) ShowDetails(c echo.Context) error {
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(h.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	h.node = term
	h.ancestors = []domain.CourseNode{user}
	pageData := mt.TermDetailsPage{
		NodeDetailsPage:      NodeDetailsPage(h, false),
		ShowEditTermDatesURL: h.e.Reverse(ShowEditTermDates.String(), h.params.ToSlice()...),
	}
	component := pageData.Component()
	layout := CourseManagerLayout(h.e, component, user)
	return Respond(c, "", component, layout)
}

// ShowEdit implements NodeHandler.
func (h *termHandler) ShowEdit(c echo.Context) error {
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	queryParam := c.QueryParam("field")
	term, err := h.svc.GetTerm(h.params.TermID.Value.(int))
	if err != nil {
		log.Println(err)
		return err
	}
	h.node = term
	h.ancestors = []domain.CourseNode{user}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := mt.TermDetailsPage{
		NodeDetailsPage: NodeDetailsPage(h, true),
	}

	respond := func(component templ.Component) error {
		return Respond(c, ShowDetailsURL(h), component, nil)
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

// ShowFiles implements NodeHandler.
func (h *termHandler) ShowFiles(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeHandler.
func (h *termHandler) ViewFile(echo.Context) error {
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
	ListTerms         = RouteHandlerName(GET + Terms)
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

func (h CourseHandler) TermHandlers() []RouteHandler {
	return []RouteHandler{
		// Terms handlers
		{User, UserHome, GET, h.UserHome},
		{Terms, ListTerms, GET, h.ListTerms},
		// {Term, TermDetails, GET, h.TermDetails},
		{TermCalendar, ShowTermCalendar, GET, h.ShowTermCalendar},
		{Occasions, CreateOccasion, POST, h.CreateOccasion},
		{Occasion, ShowEditOccasion, GET, h.ShowEditOccasion},
		{Occasion, PostEditOccasion, POST, h.PostEditOccasion},
		{TermDates, ShowEditTermDates, GET, h.ShowEditTermDates},
		{TermDates, PostEditTermDates, POST, h.PostEditTermDates},
		{NewTerm, ShowNewTerm, GET, h.ShowNewTerm},
		{NewTerm, PostNewTerm, POST, h.PostNewTerm},
		// {EditTerm, ShowEditTerm, GET, h.ShowEditTerm},
		// {EditTerm, PostEditTerm, POST, h.PostEditTerm},
		{Term, DeleteTerm, DELETE, h.DeleteTerm},
	}
}

func TermHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	handler := NewTermHandler(svc, router)
	var routeHandlers []RouteHandler
	routeHandlers = append(routeHandlers, NodeHandlers(handler, EmptyNodesTerm...)...)
	return routeHandlers
}

func (h CourseHandler) ListTerms(c echo.Context) error {
	userIDParam := c.Param("user-id")
	log.Println("useridparam: ", userIDParam)
	userID := c.Get("id")
	user, err := h.svc.GetUser(userID.(string))
	if err != nil {
		return err
	}
	log.Println(userID)
	params := ParseCourseIDParams(c)
	if !params.UserID.Valid {
		return fmt.Errorf("error parsing user ID: %s", params.UserID.Value)
	}
	log.Println("user ID:", params.UserID.Value.(string))
	terms, err := h.svc.GetTerms(userID.(string))
	if err != nil {
		return fmt.Errorf("error in CourseHandler.ListTerms: %s", err)
	}
	var termNodes []domain.CourseNode
	for _, term := range terms {
		termNodes = append(termNodes, term)
	}
	log.Println("RHN:", ShowTermCalendar.String())
	log.Println(h.e.Reverse(ShowTermCalendar.String(), terms[0].GetID()))
	page := mt.TermsListPage{
		ShowTermCalendarRHN: ShowTermCalendar.String(),
		NodeListPage: mt.NodeListPage{
			Params:           params,
			ParentNode:       user,
			Children:         termNodes,
			ChildDetailsRHN:  TermDetails.String(),
			ChildChildrenRHN: ListTermCourses.String(),
			CreateChildRHN:   ShowNewTerm.String(),
			DeleteChildRHN:   DeleteTerm.String(),
			UpNavURL:         h.e.Reverse(UserHome.String(), userID),
			E:                h.e,
		},
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

// func (h CourseHandler) TermDetails(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	user, err := h.svc.GetUser(params.UserID.Value.(string))
// 	if err != nil {
// 		return err
// 	}

// 	termID, err := TermIDParam(params)
// 	if err != nil {
// 		return err
// 	}
// 	term, err := h.svc.GetTerm(termID)
// 	if err != nil {
// 		return err
// 	}
// 	pageData := mt.TermDetailsPage{
// 		NodeDetailsPage: mt.NodeDetailsPage{
// 			Params:          params,
// 			Node:            term,
// 			GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), params.ToIntSlice()...),
// 			PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), params.ToIntSlice()...),
// 			ListChildrenURL: h.e.Reverse(ListTermCourses.String(), params.ToIntSlice()...),
// 			UpNavURL:        h.e.Reverse(ListTerms.String()),
// 			BreadCrumbsData: mt.BreadCrumbs{
// 				User:           user,
// 				UserDetailsURL: h.e.Reverse(UserHome.String(), params.ToIntSlice()...),
// 				Term:           term,
// 				TermDetailsURL: h.e.Reverse(TermDetails.String(), params.ToIntSlice()...),
// 			},
// 		},
// 		ShowEditTermDatesURL: h.e.Reverse(ShowEditTermDates.String(), termID),
// 	}
// 	template := pageData.Component()
// 	layout := h.CourseManagerLayout(template, user)
// 	return Respond(c, "", pageData.Component(), layout)
// }

func (h CourseHandler) ShowTermCalendar(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	log.Println("termID", params.TermID.Value)
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("create occasion URL: ", h.e.Reverse(CreateOccasion.String(), params.ToSlice()...))
	data := mt.TermCalendar{
		Params:              params,
		Term:                term,
		GetEditOccasionRHN:  ShowEditOccasion.String(),
		PostEditOccasionRHN: PostEditOccasion.String(),
		ListTermsURL:        h.e.Reverse(ListTerms.String(), params.ToSlice()...),
		TermDetailsURL:      h.e.Reverse(TermDetails.String(), params.ToSlice()...),
		CreateOccasionURL:   h.e.Reverse(CreateOccasion.String(), params.ToSlice()...),
		E:                   h.e,
	}

	component := data.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)

}

func (h CourseHandler) CreateOccasion(c echo.Context) error {
	params := ParseCourseIDParams(c)
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
	occasion, err := h.svc.CreateOccasion(date, name, params.TermID.Value.(int))
	if err != nil {
		return err
	}
	log.Println(occasion)
	return c.Redirect(303, h.e.Reverse(ShowTermCalendar.String(), params.ToSlice()...))
}

func (h CourseHandler) ShowEditOccasion(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	occasionID, err := ParseRouteParam(c, OccasionID)
	if err != nil {
		return err
	}
	occasion, err := h.svc.GetOccasion(occasionID)
	if err != nil {
		return err
	}
	component := mt.TermOccasionEditor{
		Occasion:            occasion,
		IsEditing:           true,
		GetEditOccasionURL:  h.e.Reverse(ShowEditOccasion.String(), params.ToSlice(occasion.ID)...),
		PostEditOccasionURL: h.e.Reverse(PostEditOccasion.String(), params.ToSlice(occasion.ID)...),
	}.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)

}

func (h CourseHandler) PostEditOccasion(c echo.Context) error {
	params := ParseCourseIDParams(c)
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
	err = h.svc.UpdateOccasion(name, occasionID)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(ShowTermCalendar.String(), params.ToSlice()...))
}

func (h CourseHandler) ShowNewTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	nodeCreate := mt.NodeCreatePage{
		ParentNode:        user,
		NodeType:          domain.TermTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewTerm.String(), params.ToSlice()...),
		CancelURL:         h.e.Reverse(ListTerms.String(), params.ToSlice()...),
	}
	template := mt.NodeCreateComponent(nodeCreate)
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)

}

func (h CourseHandler) PostNewTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

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
	startDateStr := c.FormValue("start-date")
	startDate, err := time.Parse(time.DateOnly, startDateStr)
	if err != nil {
		return err
	}
	endDateStr := c.FormValue("end-date")
	endDate, err := time.Parse(time.DateOnly, endDateStr)
	if err != nil {
		return err
	}
	term := domain.Term{
		Name:        name,
		Description: description,
		Start:       startDate,
		End:         endDate,
	}
	term.ID, err = h.svc.SaveTerm(service.SaveTermParams{
		Name:        name,
		Description: description,
		Start:       startDate,
		End:         endDate,
	})
	if err != nil {
		return err
	}
	page := mt.NodeDetailsPage{
		Node:            term,
		GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), term.ID),
		PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), term.ID),
		UpNavURL:        h.e.Reverse(ListTerms.String()),
	}
	template := page.Component()
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

// func (h CourseHandler) ShowEditTerm(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	user, err := h.svc.GetUser(params.UserID.Value.(string))
// 	if err != nil {
// 		return err
// 	}

// 	queryParam := c.QueryParam("field")
// 	termID, err := TermIDParam(params)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	term, err := h.svc.GetTerm(termID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	if queryParam == "" {
// 		log.Println(err)
// 		return fmt.Errorf("field query param is missing")
// 	}
// 	details := mt.TermDetailsPage{
// 		NodeDetailsPage: mt.NodeDetailsPage{
// 			Params:          params,
// 			Node:            term,
// 			GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), params.ToSlice()...),
// 			PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), params.ToSlice()...),
// 			UpNavURL:        h.e.Reverse(ListTerms.String(), params.ToSlice()...),
// 			IsEdit:          true,
// 			BreadCrumbsData: mt.BreadCrumbs{
// 				User:           user,
// 				UserDetailsURL: h.e.Reverse(UserHome.String(), params.ToSlice()...),

// 				Term:           term,
// 				TermDetailsURL: h.e.Reverse(TermDetails.String(), params.ToSlice()...),
// 			},
// 		},
// 	}

// 	respond := func(component templ.Component) error {
// 		return Respond(c, h.e.Reverse(TermDetails.String(), params.ToSlice()...), component, nil)
// 	}
// 	if queryParam == templates.KebabCase(string(Description)) {
// 		return respond(mt.EditDescriptionComponent(details.NodeDetailsPage))
// 	} else if queryParam == templates.KebabCase(string(Name)) {
// 		return respond(mt.EditNameComponent(details.NodeDetailsPage))
// 	}
// 	errText := "field value is not expected"
// 	log.Println(errText)
// 	return fmt.Errorf("%s %s", errText, queryParam)

// }

// func (h CourseHandler) PostEditTerm(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	user, err := h.svc.GetUser(params.UserID.Value.(string))
// 	if err != nil {
// 		return err
// 	}

// 	termID, err := TermIDParam(params)
// 	if err != nil {
// 		return err
// 	}
// 	term, err := h.svc.GetTerm(termID)
// 	if err != nil {
// 		return err
// 	}
// 	err = c.Request().ParseForm()
// 	if err != nil {
// 		return err
// 	}
// 	form := c.Request().Form
// 	var updateTerm = func(term domain.Term) (domain.Term, error) {
// 		log.Println("updating: ", term.ID, term.Name, term.Description)
// 		err := h.svc.UpdateTerm(term)
// 		if err != nil {
// 			return domain.Term{}, err
// 		}
// 		updatedTerm, err := h.svc.GetTerm(termID)
// 		if err != nil {
// 			return domain.Term{}, err
// 		}
// 		log.Println("retrieved term: ", updatedTerm.ID, updatedTerm.Name, updatedTerm.Description)
// 		return updatedTerm, nil
// 	}
// 	var pageData = func(unit domain.Term) mt.NodeDetailsPage {
// 		return mt.NodeDetailsPage{
// 			Node:            unit,
// 			Params:          params,
// 			GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), params.ToSlice()...),
// 			PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), params.ToSlice()...),
// 			IsEdit:          false,
// 			BreadCrumbsData: mt.BreadCrumbs{
// 				User:           user,
// 				UserDetailsURL: h.e.Reverse(UserHome.String(), params.ToSlice()...),

// 				Term:           term,
// 				TermDetailsURL: h.e.Reverse(TermDetails.String(), params.ToSlice()...),
// 			},
// 		}
// 	}
// 	var template templ.Component
// 	for key, val := range form {
// 		log.Println(key, val)
// 		switch key {
// 		case "description":
// 			term.Description = val[0]
// 			term, err := updateTerm(term)
// 			if err != nil {
// 				return err
// 			}
// 			details := pageData(term)
// 			template = mt.EditDescriptionComponent(details)
// 		case "name":
// 			term.Name = val[0]
// 			term, err := updateTerm(term)
// 			if err != nil {
// 				return err
// 			}
// 			details := pageData(term)
// 			template = mt.EditNameComponent(details)
// 		default:
// 			log.Println("form key:", key)
// 			panic("form key not expected!")
// 		}

// 	}
// 	if template == nil {
// 		panic("template is nil!")
// 	}
// 	return Respond(c, h.e.Reverse(string(TermDetails), params.ToSlice()...), template, nil)

// }

func (h CourseHandler) DeleteTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	err = h.svc.DeleteTerm(termID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h CourseHandler) ShowEditTermDates(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(termID)
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
		GetAddDayURL:   h.e.Reverse(ShowEditTermDates.String(), termID),
		PostAddDayURL:  h.e.Reverse(PostEditTermDates.String(), termID),
		TermDetailsURL: h.e.Reverse(TermDetails.String(), termID),
	}
	template := pageData.Component()
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) PostEditTermDates(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(termID)
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
	err = h.svc.AddNonInstructDay(termID, date)
	if err != nil {
		return err
	}
	updatedTerm, err := h.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	term = updatedTerm

	var pageData = mt.AddNonInstructDayPage{
		Term:           term,
		GetAddDayURL:   h.e.Reverse(ShowEditTermDates.String(), term.ID),
		PostAddDayURL:  h.e.Reverse(PostEditTermDates.String(), term.ID),
		TermDetailsURL: h.e.Reverse(TermDetails.String(), term.ID),
	}
	var template = pageData.Component()
	return Respond(c, h.e.Reverse(string(TermDetails), params.ToSlice()...), template, nil)

}

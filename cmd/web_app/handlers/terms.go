package handlers

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/templates"
	mt "gh_static_portfolio/cmd/templates/manager_templates"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	Terms    RouteName = "/terms"
	Term     RouteName = Terms + RouteName(TermID)
	EditTerm RouteName = Term + "/edit"
	NewTerm  RouteName = Terms + "/new"
)

const (
	ListTerms    = RouteHandlerName(GET + Terms)
	TermDetails  = RouteHandlerName(GET + Term)
	ShowEditTerm = RouteHandlerName(GET + EditTerm)
	PostEditTerm = RouteHandlerName(POST + EditTerm)
	ShowNewTerm  = RouteHandlerName(GET + NewTerm)
	PostNewTerm  = RouteHandlerName(POST + NewTerm)
	DeleteTerm   = RouteHandlerName(POST + Term)
)

func (h CourseHandler) TermHandlers() []RouteHandler {
	return []RouteHandler{
		// Terms handlers
		{Terms, ListTerms, GET, h.ListTerms},
		{Term, TermDetails, GET, h.TermDetails},
		{NewTerm, ShowNewTerm, GET, h.ShowNewTerm},
		{NewTerm, PostNewTerm, POST, h.PostNewTerm},
		{EditTerm, ShowEditTerm, GET, h.ShowEditTerm},
		{EditTerm, PostEditTerm, POST, h.PostEditTerm},
		{Term, DeleteTerm, DELETE, h.DeleteTerm},
	}
}

func (h CourseHandler) ListTerms(c echo.Context) error {
	params := ParseCourseIDParams(c)
	terms, err := h.svc.GetTerms()
	if err != nil {
		return fmt.Errorf("error in CourseHandler.ListTerms: %s", err)
	}
	var termNodes []domain.CourseNode
	for _, term := range terms {
		termNodes = append(termNodes, term)
	}
	termsList := mt.NodeListPage{
		Params:           params,
		ParentNode:       domain.RootCourseNode{},
		Children:         termNodes,
		ChildDetailsRHN:  TermDetails.String(),
		ChildChildrenRHN: ListTermCourses.String(),
		CreateChildRHN:   ShowNewTerm.String(),
		DeleteChildRHN:   DeleteTerm.String(),
		UpNavURL:         h.e.Reverse(string(ShowHome)),
		E:                h.e,
	}
	template := mt.NodeListComponent(termsList)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) TermDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	pageData := mt.NodeDetailsPage{
		Params:          params,
		Node:            term,
		GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), params.ToIntSlice()...),
		UpNavURL:        h.e.Reverse(ListTerms.String()),
	}
	template := pageData.Component()
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", pageData.Component(), layout)
}

func (h CourseHandler) ShowNewTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	nodeCreate := mt.NodeCreatePage{
		ParentNode:        domain.RootCourseNode{},
		NodeType:          domain.TermTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewTerm.String(), params.ToIntSlice()...),
		CancelURL:         h.e.Reverse(ListTerms.String(), params.ToIntSlice()...),
	}
	template := mt.NodeCreateComponent(nodeCreate)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)

}

func (h CourseHandler) PostNewTerm(c echo.Context) error {
	err := c.Request().ParseForm()
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
	term.ID, err = h.svc.SaveTerm(term)
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
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShowEditTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	queryParam := c.QueryParam("field")
	termID, err := TermIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	term, err := h.svc.GetTerm(termID)
	if err != nil {
		log.Println(err)
		return err
	}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := mt.NodeDetailsPage{
		Params:          params,
		Node:            term,
		GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), params.ToIntSlice()...),
		UpNavURL:        h.e.Reverse(ListTerms.String(), params.ToIntSlice()...),
		IsEdit:          true,
	}
	respond := func(component templ.Component) error {
		return Respond(c, h.e.Reverse(TermDetails.String(), params.ToIntSlice()...), component, nil)
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

func (h CourseHandler) PostEditTerm(c echo.Context) error {
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
	var updateUnit = func() error {
		err := h.svc.UpdateTerm(term)
		if err != nil {
			return err
		}
		updatedTerm, err := h.svc.GetTerm(termID)
		if err != nil {
			return err
		}
		term = updatedTerm
		return nil
	}
	var pageData = func(unit domain.Term) mt.NodeDetailsPage {
		return mt.NodeDetailsPage{
			Node:            unit,
			Params:          params,
			GetEditNodeURL:  h.e.Reverse(ShowEditTerm.String(), params.ToIntSlice()...),
			PostEditNodeURL: h.e.Reverse(PostEditTerm.String(), params.ToIntSlice()...),
			IsEdit:          false,
		}
	}
	var template templ.Component
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "description":
			term.Description = val[0]
			err := updateUnit()
			if err != nil {
				return err
			}
			details := pageData(term)
			template = mt.EditDescriptionComponent(details)
		case "name":
			term.Name = val[0]
			err := updateUnit()
			if err != nil {
				return err
			}
			details := pageData(term)
			template = mt.EditNameComponent(details)
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	if template == nil {
		panic("template is nil!")
	}
	return Respond(c, h.e.Reverse(string(TermDetails), params.ToIntSlice()...), template, nil)

}

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

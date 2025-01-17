package handlers

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/templates"
	mt "gh_static_portfolio/cmd/templates/manager_templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	Courses    RouteName = Term + "/courses"
	Course     RouteName = Courses + RouteName(CourseID)
	NewCourse  RouteName = Courses + "/new"
	EditCourse RouteName = Course + "/edit"
)
const (
	ListTermCourses = RouteHandlerName(GET + Courses)
	CourseDetails   = RouteHandlerName(GET + Course)
	ShowEditCourse  = RouteHandlerName(GET + EditCourse)
	PostEditCourse  = RouteHandlerName(POST + EditCourse)
	ShowNewCourse   = RouteHandlerName(GET + NewCourse)
	PostNewCourse   = RouteHandlerName(POST + NewCourse)
	DeleteCourse    = RouteHandlerName(DELETE + Course)
)

func (h CourseHandler) CourseHandlers() []RouteHandler {
	return []RouteHandler{
		// Courses handlers
		{Courses, ListTermCourses, GET, h.ListTermCourses},
		{Course, CourseDetails, GET, h.CourseDetails},
		{NewCourse, ShowNewCourse, GET, h.ShowNewCourse},
		{NewCourse, PostNewCourse, POST, h.PostNewCourse},
		{EditCourse, ShowEditCourse, GET, h.ShowEditCourse},
		{EditCourse, PostEditCourse, POST, h.PostEditCourse},
		{Course, DeleteCourse, DELETE, h.DeleteCourse},
	}
}

func (h CourseHandler) ListTermCourses(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := ParseRouteParam(c, TermID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("termID: ", termID)
	term, err := h.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	courses, err := h.svc.GetCourses(termID)
	if err != nil {
		log.Println(err)
		return err
	}
	term.Courses = courses
	coursesList := mt.NodeListPage{
		Params:           params,
		ParentNode:       term,
		Children:         term.Children(),
		ChildDetailsRHN:  CourseDetails.String(),
		CreateChildRHN:   ShowNewCourse.String(),
		ChildChildrenRHN: ListCourseUnits.String(),
		UpNavURL:         h.e.Reverse(ListTerms.String()),
		E:                h.e,
	}
	template := mt.NodeListComponent(coursesList)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) CourseDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	courseID, err := CourseIDParam(params)
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(courseID)
	if err != nil {
		return err
	}
	courseDetails := mt.NodeDetailsPage{
		Params:          params,
		Node:            *course,
		GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), params.ToIntSlice()...),
		UpNavURL:        h.e.Reverse(ListTermCourses.String(), params.ToIntSlice()...),
		IsEdit:          false,
	}
	template := mt.NodeDetailsComponent(courseDetails)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShowNewCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	nodeCreate := mt.NodeCreatePage{
		ParentNode:        term,
		NodeType:          domain.CourseTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewCourse.String(), params.ToIntSlice()...),
		CancelURL:         h.e.Reverse(ListTermCourses.String(), params.ToIntSlice()...),
	}
	template := mt.NodeCreateComponent(nodeCreate)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) PostNewCourse(c echo.Context) error {
	return nil
}

func (h CourseHandler) ShowEditCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	queryParam := c.QueryParam("field")
	courseID, err := CourseIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	course, err := h.svc.GetCourse(courseID)
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
		Node:            *course,
		GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), params.ToIntSlice()...),
		IsEdit:          true,
	}
	respond := func(component templ.Component) error {
		return Respond(c, h.e.Reverse(string(UnitDetails), params.ToIntSlice()...), component, nil)
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

func (h CourseHandler) PostEditCourse(c echo.Context) error {
	return nil
}

func (h CourseHandler) DeleteCourse(c echo.Context) error {
	panic("not implemented")
}

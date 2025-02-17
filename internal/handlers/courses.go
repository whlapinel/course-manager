package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"gh_static_portfolio/internal/templates"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	Courses          RouteName = Term + "/courses"
	Course           RouteName = Courses + RouteName(CourseID)
	CourseImage      RouteName = Course + "/image"
	NewCourse        RouteName = Courses + "/new"
	EditCourse       RouteName = Course + "/edit"
	CopyCourse       RouteName = Course + "/copy-to-term"
	CopyCourseToTerm RouteName = CopyCourse
	StandardSet      RouteName = Course + "/standard-set"
)
const (
	ListTermCourses       = RouteHandlerName(GET + Courses)
	CourseDetails         = RouteHandlerName(GET + Course)
	ShowEditCourse        = RouteHandlerName(GET + EditCourse)
	PostEditCourse        = RouteHandlerName(POST + EditCourse)
	ShowNewCourse         = RouteHandlerName(GET + NewCourse)
	PostNewCourse         = RouteHandlerName(POST + NewCourse)
	DeleteCourse          = RouteHandlerName(DELETE + Course)
	GetCopyCourse         = RouteHandlerName(GET + CopyCourse)
	PostCopyCourseToTerm  = RouteHandlerName(POST + CopyCourseToTerm)
	PostSelectStandardSet = RouteHandlerName(POST + StandardSet)
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
		{CopyCourse, GetCopyCourse, GET, h.GetCopyCourse},
		{CopyCourseToTerm, PostCopyCourseToTerm, POST, h.PostCopyCourseToTerm},
		{StandardSet, PostSelectStandardSet, POST, h.PostSelectStandardSet},
	}
}

func (h CourseHandler) ListTermCourses(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID := params.TermID.Value.(int)
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
	page := mt.CourseListPage{
		ShowAssessmentsRHN: string(GetCourseAssessments),
		ShowCalendarRHN:    string(ShowCourseCalendar),
		NodeListPage: mt.NodeListPage{
			Params:           params,
			ParentNode:       term,
			Children:         term.Children(),
			ChildDetailsRHN:  CourseDetails.String(),
			CreateChildRHN:   ShowNewCourse.String(),
			ChildChildrenRHN: ListCourseUnits.String(),
			DeleteChildRHN:   DeleteCourse.String(),
			UpNavURL:         h.e.Reverse(ListTerms.String(), params.ToIntSlice()...),
			E:                h.e,
			BreadCrumbsData: mt.BreadCrumbs{
				Term:           term,
				TermDetailsURL: h.e.Reverse(TermDetails.String(), params.ToIntSlice()...),
			},
		},
	}

	component := page.Component()
	layout := h.CourseManagerLayout(component)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) CourseDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	sets, err := h.svc.GetStandardSets()
	if err != nil {
		return err
	}
	page := mt.CourseDetailsPage{
		GetCopyCourseURL:         h.e.Reverse(GetCopyCourse.String(), params.ToIntSlice()...),
		StandardSets:             sets,
		PostSelectStandardSetURL: h.e.Reverse(string(PostSelectStandardSet), params.ToIntSlice()...),
		NodeDetailsPage: mt.NodeDetailsPage{
			Params:          params,
			Node:            course,
			GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), params.ToIntSlice()...),
			PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), params.ToIntSlice()...),
			ListChildrenURL: h.e.Reverse(ListCourseUnits.String(), params.ToIntSlice()...),
			UpNavURL:        h.e.Reverse(ListTermCourses.String(), params.ToIntSlice()...),
			IsEdit:          false,
			BreadCrumbsData: mt.BreadCrumbs{
				Term:             course.Term,
				TermDetailsURL:   h.e.Reverse(TermDetails.String(), params.ToIntSlice()...),
				Course:           course,
				CourseDetailsURL: h.e.Reverse(CourseDetails.String(), params.ToIntSlice()...),
			},
		},
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component)
	return Respond(c, "", component, layout)
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
		BreadCrumbsData: mt.BreadCrumbs{
			Term:           term,
			TermDetailsURL: h.e.Reverse(TermDetails.String(), params.ToIntSlice()...),
		},
	}
	template := mt.NodeCreateComponent(nodeCreate)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) PostNewCourse(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for key, val := range form {
		log.Println("key, val: ", key, val)
	}
	termID := ParseCourseIDParams(c).TermID
	name := c.FormValue("name")
	description := c.FormValue("description")
	course, err := h.svc.SaveCourse(service.SaveCourseParams{
		TermID:      termID.Value.(int),
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}
	page := mt.NodeDetailsPage{
		Node:            course,
		GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), course.ID),
		PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), course.ID),
		UpNavURL:        h.e.Reverse(ListTermCourses.String(), termID),
	}
	template := page.Component()
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
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
		Node:            course,
		GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), params.ToIntSlice()...),
		ListChildrenURL: h.e.Reverse(ListCourseUnits.String(), params.ToIntSlice()...),
		IsEdit:          true,
		BreadCrumbsData: mt.BreadCrumbs{
			Term:             course.Term,
			TermDetailsURL:   h.e.Reverse(TermDetails.String(), params.ToIntSlice()...),
			Course:           course,
			CourseDetailsURL: h.e.Reverse(CourseDetails.String(), params.ToIntSlice()...),
		},
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
	params := ParseCourseIDParams(c)
	courseId := params.CourseID
	return h.svc.DeleteCourse(courseId.Value.(int))
}

func (h CourseHandler) GetCopyCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	terms, err := h.svc.GetTerms(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	data := mt.CopyCourseData{
		TermID:                  params.TermID.Value.(int),
		CourseID:                params.CourseID.Value.(int),
		Terms:                   terms,
		E:                       h.e,
		PostCopyCourseToTermRHN: string(PostCopyCourseToTerm),
	}
	component := data.Component()
	return Respond(c, h.e.Reverse(ListTermCourses.String(), params.ToIntSlice()...), component, nil)
}

func (h CourseHandler) PostCopyCourseToTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	if params.CourseID.Valid && params.TermID.Valid {
		err := c.Request().ParseForm()
		if err != nil {
			return err
		}
		termIDParam := c.Request().Form.Get("term-id")
		log.Println("selected termID: ", termIDParam)
		termID, err := strconv.Atoi(termIDParam)
		if err != nil {
			return err
		}
		_, err = h.svc.CopyCourseToTerm(params.CourseID.Value.(int), termID)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("params not valid: courseID: %d and termID: %d", params.CourseID.Value, params.TermID.Value)
	}
	return c.Redirect(302, h.e.Reverse(ListTermCourses.String(), params.ToIntSlice()...))
}

func (h CourseHandler) PostSelectStandardSet(c echo.Context) error {
	params := ParseCourseIDParams(c)
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	standardSetParam := c.Request().Form.Get("standard-set")
	setID, err := strconv.Atoi(standardSetParam)
	if err != nil {
		return err
	}
	log.Println("selected set: ", standardSetParam)
	err = h.svc.SetStandardSet(params.CourseID.Value.(int), setID)
	if err != nil {
		return err
	}
	return c.Redirect(302, h.e.Reverse(CourseDetails.String(), params.ToIntSlice()...))
}

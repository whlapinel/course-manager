package handlers

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	CourseAssessments   RoutePath = Course + "/assessments"
	LessonAssessments   RoutePath = Lesson + "/assessments"
	NewLessonAssessment RoutePath = LessonAssessments + "/new"
	Assessment          RoutePath = LessonAssessments + RoutePath(AssessmentID)
)

const (
	GetCourseAssessments   = RouteHandlerName(GET + CourseAssessments)
	ShowNewAsssessmentForm = RouteHandlerName(GET + NewLessonAssessment)
	PostAssessment         = RouteHandlerName(POST + LessonAssessments)
	GetEditAssessment      = RouteHandlerName(GET + Assessment)
	PostEditAssessment     = RouteHandlerName(POST + Assessment)
	DeleteAssessment       = RouteHandlerName(DELETE + Assessment)
)

func (h CourseHandler) AssessmentHandlers() []RouteHandler {
	return []RouteHandler{
		{CourseAssessments, GetCourseAssessments, GET, h.GetCourseAssessments},
		{LessonAssessments, PostAssessment, POST, h.PostAssessment},
		{Assessment, GetEditAssessment, GET, h.GetEditAssessment},
		{Assessment, PostEditAssessment, POST, h.PostEditAssessment},
		{Assessment, DeleteAssessment, DELETE, h.DeleteAssessment},
	}
}

func (h CourseHandler) GetCourseAssessments(c echo.Context) error {
	var err error
	var category domain.AssessmentCategory
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.Nodes(params)
	if err != nil {
		return err
	}
	categoryParam := c.QueryParam("category")
	startDateParam := c.QueryParam("start")
	var start time.Time
	if startDateParam != "" {
		start, err = time.Parse(time.DateOnly, startDateParam)
		if err != nil {
			return err
		}
	} else {
		start = time.Now().AddDate(-10, 0, 0)
	}
	endDateParam := c.QueryParam("end")
	log.Println(endDateParam)
	var end time.Time
	if endDateParam != "" {
		end, err = time.Parse(time.DateOnly, endDateParam)
		if err != nil {
			return err
		}
	} else {
		end = time.Now().AddDate(10, 0, 0)
	}
	var assessments = []domain.Assessment{}
	if categoryParam != "" {
		// category filter
		category, err = domain.ParseCategories(categoryParam)
		if err != nil {
			return err
		}
		assessments, err = h.svc.FilterAssessmentsByCategoryAndDate(category, params.CourseID, start, end)
		if err != nil {
			return err
		}
	} else {
		// no category filter
		assessments, err = h.svc.GetAllCourseAssessments(params.CourseID)
		if err != nil {
			return err
		}
	}
	var queryParams = make(map[string]string)
	queryParams["category"] = category.String()
	queryParams["start"] = startDateParam
	queryParams["end"] = endDateParam
	baseURL := h.e.Reverse(GetCourseAssessments.String(), params.ToSlice()...)
	urlWithParams, err := AddQueryParams(baseURL, queryParams)
	if err != nil {
		return err
	}
	page := mt.CourseAssessmentsPage{
		GetAssessmentsURL: urlWithParams,
		Assessments:       assessments,
		CourseListURL:     h.e.Reverse(ListTermCourses.String(), params.ToSlice()...),
		BreadCrumbsData:   BreadCrumbs(h.e, params, nodes.ToSlice()...),
	}
	return Respond(c, "", page.Component(), h.CourseManagerLayout(page.Component(), nodes.User))
}

func (h CourseHandler) PostAssessment(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	name := form.Get("name")
	file := form.Get("file")
	instructions := form.Get("instructions")
	categoryParam := form.Get("category")
	category, err := strconv.Atoi(categoryParam)
	if err != nil {
		return err
	}
	assignedParam := form.Get("date-assigned")
	assigned, err := time.Parse(time.DateOnly, assignedParam)
	if err != nil {
		return err
	}
	dueParam := form.Get("date-due")
	due, err := time.Parse(time.DateOnly, dueParam)
	if err != nil {
		return err
	}
	_, err = h.svc.SaveAssessment(service.SaveAssessmentParams{
		Assessment: domain.Assessment{
			LessonID:     params.LessonID,
			Name:         name,
			Instructions: instructions,
			File:         file,
			Category:     domain.AssessmentCategory(category),
			DateAssigned: assigned,
			DateDue:      due,
			Dropped:      false,
		},
	})
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(string(ShowNodeDetailsRHN(EmptyNodesLesson...)), params.ToSlice()...))
}

func (h CourseHandler) GetEditAssessment(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.Nodes(params)
	if err != nil {
		return err
	}
	assessmentID, err := ParseRouteParam(c, AssessmentID)
	if err != nil {
		return err
	}
	assessment, err := h.svc.GetAssessment(assessmentID)
	if err != nil {
		return err
	}
	data := mt.EditAssessmentForm{
		Params:                params,
		Assessment:            assessment,
		PostEditAssessmentURL: h.e.Reverse(PostEditAssessment.String(), AddParams(params, assessmentID)...),
		LessonDetailsURL:      h.e.Reverse(string(ShowNodeDetailsRHN(EmptyNodesLesson...)), params.ToSlice()...),
	}
	return Respond(c, "", data.Component(), h.CourseManagerLayout(data.Component(), nodes.User))
}

func (h CourseHandler) PostEditAssessment(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	assessmentID, err := ParseRouteParam(c, AssessmentID)
	if err != nil {
		return err
	}
	form := c.Request().Form
	name := form.Get("name")
	instructions := form.Get("instructions")
	file := form.Get("file")
	categoryParam := form.Get("category")
	category, err := strconv.Atoi(categoryParam)
	if err != nil {
		return err
	}
	assignedParam := form.Get("date-assigned")
	assigned, err := time.Parse(time.DateOnly, assignedParam)
	if err != nil {
		return err
	}
	dueParam := form.Get("date-due")
	due, err := time.Parse(time.DateOnly, dueParam)
	if err != nil {
		return err
	}
	droppedParam := form.Get("dropped")
	var dropped bool
	if droppedParam == "true" {
		dropped = true
	}
	updateParams := service.UpdateAssessmentParams{
		Assessment: domain.Assessment{
			ID:           assessmentID,
			LessonID:     params.LessonID,
			Name:         name,
			Instructions: instructions,
			File:         file,
			Category:     domain.AssessmentCategory(category),
			DateAssigned: assigned,
			DateDue:      due,
			Dropped:      dropped,
		},
	}
	log.Println("Handler: PostEditAssessment: ", updateParams.Assessment)
	err = h.svc.UpdateAssessment(updateParams)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(string(ShowNodeDetailsRHN(EmptyNodesLesson...)), params.ToSlice()...))
}

func (h CourseHandler) DeleteAssessment(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	id, err := ParseRouteParam(c, AssessmentID)
	if err != nil {
		return err
	}
	err = h.svc.DeleteAssessment(id)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(string(ShowNodeDetailsRHN(EmptyNodesLesson...)), params.ToSlice()...))
}

package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/app"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	Calendar              RoutePath = Course + "/calendar"
	ShiftLessonRouteName  RoutePath = Lesson + RoutePath(ShiftDirection)
	ExtendLessonRouteName RoutePath = ShiftLessonRouteName + "/extend"
	CalendarDate          RoutePath = Calendar + RoutePath(Date)
	ListLessons           RoutePath = CalendarDate + RoutePath(UnitID)
	LessonDates           RoutePath = Lesson + "/dates"
	LessonDate            RoutePath = LessonDates + RoutePath(Date)
)
const (
	ShowCourseCalendar    = RouteHandlerName(GET + Calendar)
	ShiftLessonRHN        = RouteHandlerName(POST + ShiftLessonRouteName)
	ExtendLessonRHN       = RouteHandlerName(POST + ExtendLessonRouteName)
	ShowAddLessonDatePage = RouteHandlerName(GET + CalendarDate)
	ListLessonsRHN        = RouteHandlerName(GET + ListLessons)
	RemoveLessonDate      = RouteHandlerName(DELETE + LessonDate)
	PostAddLessonDate     = RouteHandlerName(POST + LessonDate)
)

func (h CourseHandler) CalendarHandlers() []RouteHandler {
	return []RouteHandler{
		{Calendar, ShowCourseCalendar, GET, h.ShowCourseCalendar},
		{ShiftLessonRouteName, ShiftLessonRHN, POST, h.ShiftLesson},
		{ExtendLessonRouteName, ExtendLessonRHN, POST, h.ExtendLesson},
		{CalendarDate, ShowAddLessonDatePage, GET, h.ShowAddLessonDatePage},
		{ListLessons, ListLessonsRHN, GET, h.ListUnitLessons},
		{LessonDate, RemoveLessonDate, DELETE, h.RemoveLessonDate},
		{LessonDate, PostAddLessonDate, POST, h.AddLessonDate},
	}
}

func (h CourseHandler) ShowCourseCalendar(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.Nodes(params)
	if err != nil {
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component, nodes.User)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ShiftLesson(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.Nodes(params)
	if err != nil {
		return err
	}

	cdParam := ParseRouteStringParam(c, ShiftDirection)
	cd, err := domain.ParseDirection(cdParam)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.svc.WebShift(params.TermID, params.CourseID, params.LessonID, cd)
	if err != nil {
		log.Println(err)
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component, nodes.User)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ExtendLesson(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.Nodes(params)
	if err != nil {
		return err
	}

	cdParam := ParseRouteStringParam(c, ShiftDirection)
	cd, err := domain.ParseDirection(cdParam)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.svc.Extend(params.TermID, params.CourseID, params.LessonID, cd)
	if err != nil {
		log.Println(err)
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component, nodes.User)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ShowAddLessonDatePage(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.Nodes(params)
	if err != nil {
		return err
	}

	date, err := ParseDateParam(c)
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourseWithChildren(params.CourseID)
	if err != nil {
		return err
	}
	page := mt.AddLessonToDatePage{
		Date:             date,
		Params:           params,
		Course:           course,
		E:                h.e,
		ListLessonsRHN:   ListLessonsRHN.String(),
		AddLessonDateRHN: string(PostAddLessonDate),
		BreadCrumbsData:  BreadCrumbs(h.e, params, nodes.ToSlice()...),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, nodes.User)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) RemoveLessonDate(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	date, err := ParseDateParam(c)
	if err != nil {
		return err
	}
	err = h.svc.RemoveLessonDate(params.LessonID, date)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(ShowCourseCalendar.String(), params.ToSlice()...))
}

func (h CourseHandler) AddLessonDate(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	date, err := ParseDateParam(c)
	if err != nil {
		return err
	}
	err = h.svc.AddLessonDate(params.LessonID, params.TermID, date)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(ShowCourseCalendar.String(), params.ToSlice()...))
}

// for selecting lesson to add to calendar
func (h CourseHandler) ListUnitLessons(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.svc.NodesWithChildren(params)
	if err != nil {
		return err
	}
	dateStr := ParseRouteStringParam(c, Date)
	if dateStr == "" {
		return fmt.Errorf("date param empty")
	}
	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return err
	}
	component := mt.LessonPicker{
		Date:            date,
		Params:          params,
		ListUnitsURL:    h.e.Reverse(string(ShowAddLessonDatePage), params.UserID, params.TermID, params.CourseID, dateStr),
		Lessons:         nodes.Unit.Lessons,
		SelectLessonRHN: string(PostAddLessonDate),
		Echo:            h.e,
	}.Component()
	return Respond(c, h.e.Reverse(ShowAddLessonDatePage.String(), AddParams(params, dateStr)...), component, nil)

}

func (h CourseHandler) calendarPage(params domain.NodePath) (mt.CourseCalendar, error) {
	course, err := h.svc.GetCourseForCalendar(params.CourseID)
	if err != nil {
		log.Println(err)
		return mt.CourseCalendar{}, err
	}
	return mt.CourseCalendar{
		Admin:                         true,
		Params:                        params,
		Course:                        course,
		LessonDetailsRouteHandlerName: string(ShowNodeDetailsRHN(EmptyNodesLesson...)),
		TermDetailsURL:                h.e.Reverse(TermDetails.String(), params.ToSlice()...),
		CourseDetailsURL:              h.e.Reverse(CourseDetails.String(), params.ToSlice()...),
		ShiftLessonRouteHandlerName:   ShiftLessonRHN.String(),
		ListTermCoursesRHN:            string(ListTermCourses),
		ShowAddLessonDateRHN:          string(ShowAddLessonDatePage),
		RemoveLessonDateRHN:           RemoveLessonDate.String(),
		E:                             h.e,
	}, nil

}

func ParseDateParam(c echo.Context) (time.Time, error) {
	dateParam := ParseRouteStringParam(c, Date)
	if dateParam == "" {
		return time.Time{}, fmt.Errorf("date param is empty: %s", dateParam)
	}
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

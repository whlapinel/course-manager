package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	Calendar              RoutePath = Course + "/calendar"
	ShiftLessonRouteName  RoutePath = Lesson + RoutePath(ShiftDirection)
	ExtendLessonRouteName RoutePath = ShiftLessonRouteName + "/extend"
	CalendarDate          RoutePath = Calendar + RoutePath(Date)
	LessonDates           RoutePath = Lesson + "/dates"
	LessonDate            RoutePath = LessonDates + RoutePath(Date)
)
const (
	ShowCourseCalendar    = RouteHandlerName(GET + Calendar)
	ShiftLessonRHN        = RouteHandlerName(POST + ShiftLessonRouteName)
	ExtendLessonRHN       = RouteHandlerName(POST + ExtendLessonRouteName)
	ShowAddLessonDatePage = RouteHandlerName(GET + CalendarDate)
	RemoveLessonDate      = RouteHandlerName(DELETE + LessonDate)
	PostAddLessonDate     = RouteHandlerName(POST + LessonDate)
)

func (h CourseHandler) CalendarHandlers() []RouteHandler {
	return []RouteHandler{
		{Calendar, ShowCourseCalendar, GET, h.ShowCourseCalendar},
		{ShiftLessonRouteName, ShiftLessonRHN, POST, h.ShiftLesson},
		{ExtendLessonRouteName, ExtendLessonRHN, POST, h.ExtendLesson},
		{CalendarDate, ShowAddLessonDatePage, GET, h.ShowAddLessonDatePage},
		{LessonDate, RemoveLessonDate, DELETE, h.RemoveLessonDate},
		{LessonDate, PostAddLessonDate, POST, h.AddLessonDate},
	}
}

func (h CourseHandler) ShowCourseCalendar(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ShiftLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	cdParam := ParseRouteStringParam(c, ShiftDirection)
	cd, err := domain.ParseDirection(cdParam)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.svc.WebShift(params.TermID.Value.(int), params.CourseID.Value.(int), params.LessonID.Value.(int), cd)
	if err != nil {
		log.Println(err)
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ExtendLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	cdParam := ParseRouteStringParam(c, ShiftDirection)
	cd, err := domain.ParseDirection(cdParam)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.svc.Extend(params.TermID.Value.(int), params.CourseID.Value.(int), params.LessonID.Value.(int), cd)
	if err != nil {
		log.Println(err)
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ShowAddLessonDatePage(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	date, err := ParseDateParam(c)
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourseWithChildren(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	page := mt.AddLessonToDatePage{
		Date:             date,
		Params:           params,
		Course:           course,
		E:                h.e,
		AddLessonDateRHN: string(PostAddLessonDate),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) RemoveLessonDate(c echo.Context) error {
	params := ParseCourseIDParams(c)
	date, err := ParseDateParam(c)
	if err != nil {
		return err
	}
	err = h.svc.RemoveLessonDate(params.LessonID.Value.(int), date)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(ShowCourseCalendar.String(), params.ToSlice()...))
}

func (h CourseHandler) AddLessonDate(c echo.Context) error {
	params := ParseCourseIDParams(c)
	date, err := ParseDateParam(c)
	if err != nil {
		return err
	}
	err = h.svc.AddLessonDate(params.LessonID.Value.(int), params.TermID.Value.(int), date)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(ShowCourseCalendar.String(), params.ToSlice()...))
}

func (h CourseHandler) calendarPage(params mt.NodePath) (mt.CourseCalendar, error) {
	course, err := h.svc.GetCourseForCalendar(params.CourseID.Value.(int))
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

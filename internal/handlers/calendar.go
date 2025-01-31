package handlers

import (
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"

	"github.com/labstack/echo/v4"
)

const (
	Calendar             RouteName = Course + "/calendar"
	ShiftLessonRouteName RouteName = Lesson + RouteName(ShiftDirection)
)
const (
	ShowCourseCalendar = RouteHandlerName(GET + Calendar)
	ShiftLessonRHN     = RouteHandlerName(POST + ShiftLessonRouteName)
)

func (h CourseHandler) CalendarHandlers() []RouteHandler {
	return []RouteHandler{
		{Calendar, ShowCourseCalendar, GET, h.ShowCourseCalendar},
		{ShiftLessonRouteName, ShiftLessonRHN, POST, h.ShiftLesson},
	}
}

func (h CourseHandler) ShowCourseCalendar(c echo.Context) error {
	params := ParseCourseIDParams(c)
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	template := mt.CourseCalendarTemplate(calendarData)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShiftLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	cdParam := ParseRouteStringParam(c, ShiftDirection)
	cd, err := domain.ParseDirection(cdParam)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.svc.WebShift(params.TermID.Value, params.CourseID.Value, params.LessonID.Value, cd)
	if err != nil {
		log.Println(err)
		return err
	}
	calendarData, err := h.calendarPage(params)
	if err != nil {
		return err
	}
	component := calendarData.Component()
	layout := h.CourseManagerLayout(component)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) calendarPage(params mt.CourseIDParams) (mt.CourseCalendar, error) {
	course, err := h.svc.GetCourseForCalendar(params.CourseID.Value)
	if err != nil {
		log.Println(err)
		return mt.CourseCalendar{}, err
	}
	return mt.CourseCalendar{
		Course:                        course,
		LessonDetailsRouteHandlerName: LessonDetails.String(),
		TermDetailsURL:                h.e.Reverse(TermDetails.String(), params.ToIntSlice()...),
		CourseDetailsURL:              h.e.Reverse(CourseDetails.String(), params.ToIntSlice()...),
		ShiftLessonRouteHandlerName:   ShiftLessonRHN.String(),
		ListTermCoursesRHN:            string(ListTermCourses),
		E:                             h.e,
	}, nil

}

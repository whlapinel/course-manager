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
	ShowCourseCalendar          = RouteHandlerName(GET + Calendar)
	ShiftLessonRouteHandlerName = RouteHandlerName(POST + ShiftLessonRouteName)
)

func (h CourseHandler) CalendarHandlers() []RouteHandler {
	return []RouteHandler{
		{Calendar, ShowCourseCalendar, GET, h.ShowCourseCalendar},
		{ShiftLessonRouteName, ShiftLessonRouteHandlerName, POST, h.ShiftLesson},
	}
}

func (h CourseHandler) ShowCourseCalendar(c echo.Context) error {
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		log.Println(err)
		return err
	}
	course, err := h.svc.GetCourseForCalendar(courseID)
	if err != nil {
		log.Println(err)
		return err
	}
	template := mt.CourseCalendarTemplate(*course, string(LessonDetails), string(ShiftLessonRouteHandlerName), ListTermCourses.String(), h.e)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShiftLesson(c echo.Context) error {
	termID, err := ParseRouteParam(c, TermID)
	if err != nil {
		log.Println(err)
		return err
	}
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		log.Println(err)
		return err
	}
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	cdParam := ParseRouteStringParam(c, ShiftDirection)
	cd, err := domain.ParseDirection(cdParam)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.svc.WebShift(termID, courseID, lessonID, cd)
	if err != nil {
		log.Println(err)
		return err
	}
	course, err := h.svc.GetCourseForCalendar(courseID)
	if err != nil {
		log.Println(err)
		return err
	}
	template := mt.CourseCalendarTemplate(*course, string(LessonDetails), string(ShiftLessonRouteHandlerName), ListTermCourses.String(), h.e)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

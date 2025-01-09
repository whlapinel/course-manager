package main

import (
	"context"
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/service"
	"gh_static_portfolio/cmd/templates"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	e   *echo.Echo
	svc service.CourseService
}

func NewCourseHandler(e *echo.Echo, svc service.CourseService) CourseHandler {
	return CourseHandler{
		e:   e,
		svc: svc,
	}
}

type RouteName string

type MethodName string

type RouteParam string

const (
	GET    = "GET: "
	POST   = "POST: "
	PUT    = "PUT: "
	PATCH  = "PATCH: "
	DELETE = "DELETE: "
)

const (
	TermID   RouteParam = "/:term-id"
	CourseID RouteParam = "/:course-id"
	UnitID   RouteParam = "/:unit-id"
	LessonID RouteParam = "/:lesson-id"
)

func (p RouteParam) Name() string {
	return string(p[2:])

}

func ParseRouteParam(c echo.Context, param RouteParam) (int, error) {
	return strconv.Atoi(c.Param(param.Name()))

}

const (
	Home     RouteName = "/"
	Terms    RouteName = "/terms"
	Courses  RouteName = Terms + RouteName(TermID) + "/courses"
	Units    RouteName = Courses + RouteName(CourseID) + "/units"
	Calendar RouteName = Courses + RouteName(CourseID) + "/calendar"
	Lessons  RouteName = Units + RouteName(UnitID) + "/lessons"
	Lesson   RouteName = Lessons + RouteName(LessonID)
	Slides   RouteName = Lesson + "/slides"
)

type RouteHandlerName string

const (
	ShowHome           = RouteHandlerName(GET + Home)
	ListTerms          = RouteHandlerName(GET + Terms)
	ListTermCourses    = RouteHandlerName(GET + Courses)
	ListCourseUnits    = RouteHandlerName(GET + Units)
	ShowCourseCalendar = RouteHandlerName(GET + Calendar)
	ListUnitLessons    = RouteHandlerName(GET + Lessons)
	LessonDetails      = RouteHandlerName(GET + Lesson)
	LessonSlides       = RouteHandlerName(GET + Slides)
)

func (h CourseHandler) Mount() {
	h.e.GET(string(Home), h.ShowHome).Name = string(ShowHome)
	h.e.GET(string(Terms), h.ListTerms).Name = string(ListTerms)
	h.e.GET(string(Courses), h.ListTermCourses).Name = string(ListTermCourses)
	h.e.GET(string(Units), h.ListCourseUnits).Name = string(ListCourseUnits)
	h.e.GET(string(Calendar), h.ShowCourseCalendar).Name = string(ShowCourseCalendar)
	h.e.GET(string(Lessons), h.ListUnitLessons).Name = string(ListUnitLessons)
	h.e.GET(string(Lesson), h.LessonDetails).Name = string(LessonDetails)
	h.e.GET(string(Slides), h.LessonSlides).Name = string(LessonSlides)

}

func (h CourseHandler) ShowHome(c echo.Context) error {
	template := templates.ManagerHomePage()
	if !IsHTMX(c) {
		template = templates.CourseManagerLayout("Home", template)
	}
	return template.Render(context.Background(), c.Response())
}

func (h CourseHandler) ListTerms(c echo.Context) error {
	terms, err := h.svc.GetTerms()
	if err != nil {
		return fmt.Errorf("error in CourseHandler.ListTerms: %s", err)
	}
	if len(terms) == 0 {
		return fmt.Errorf("error in CourseHandler.ListTerms: terms is empty")
	}
	template := templates.TermsListTemplate(terms, string(ListTermCourses), h.e)
	if !IsHTMX(c) {
		template = templates.CourseManagerLayout("Terms", template)
	}
	return template.Render(context.Background(), c.Response())

}

func (h CourseHandler) ListTermCourses(c echo.Context) error {
	termID, err := ParseRouteParam(c, TermID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("termID: ", termID)
	courses, err := h.svc.GetCourses(termID)
	if err != nil {
		log.Println(err)
		return err
	}
	template := templates.CoursesListTemplate(termID, courses, string(ListCourseUnits), string(ShowCourseCalendar), h.e)
	if !IsHTMX(c) {
		template = templates.CourseManagerLayout("Courses", template)
	}
	return template.Render(context.Background(), c.Response())
}

func (h CourseHandler) ListCourseUnits(c echo.Context) error {
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
	units, err := h.svc.GetUnits(courseID)
	if err != nil {
		log.Println(err)
		return err
	}
	unitsListTemplate := templates.UnitsListTemplate(termID, courseID, units, string(ListUnitLessons), h.e)
	return unitsListTemplate.Render(context.Background(), c.Response())
}

func (h CourseHandler) ShowCourseCalendar(c echo.Context) error {
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		log.Println(err)
		return err
	}
	course, err := h.svc.GetCourse(courseID)
	if err != nil {
		log.Println(err)
		return err
	}
	calendarTemplate := templates.CourseCalendarTemplate(*course)
	return calendarTemplate.Render(context.Background(), c.Response())
}

func (h CourseHandler) ListUnitLessons(c echo.Context) error {
	TermID, err := ParseRouteParam(c, TermID)
	if err != nil {
		log.Println(err)
		return err
	}
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		log.Println(err)
		return err
	}
	unitID, err := ParseRouteParam(c, UnitID)
	if err != nil {
		log.Println(err)
		return err
	}
	lessons, err := h.svc.GetLessons(unitID)
	if err != nil {
		log.Println(err)
		return err
	}
	lessonsListTemplate := templates.LessonListTemplate(TermID, courseID, unitID, lessons, string(LessonDetails), h.e)
	return lessonsListTemplate.Render(context.Background(), c.Response())
}

func (h CourseHandler) LessonDetails(c echo.Context) error {
	TermID, err := ParseRouteParam(c, TermID)
	if err != nil {
		log.Println(err)
		return err
	}
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		log.Println(err)
		return err
	}
	unitID, err := ParseRouteParam(c, UnitID)
	if err != nil {
		log.Println(err)
		return err
	}
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		log.Println("error getting lesson:", err)
		return err
	}
	unit, err := h.svc.GetUnit(unitID)
	if err != nil {
		log.Println("error getting unit:", err)
		return err
	}
	course, err := h.svc.GetCourse(courseID)
	if err != nil {
		log.Println("error getting course:", err)
		return err
	}
	lessonDetailsTemplate := templates.LessonDetailsTemplate(TermID, courseID, unitID, lessonID, lesson, unit, course, string(LessonSlides), h.e)
	return lessonDetailsTemplate.Render(context.Background(), c.Response())

}

func (h CourseHandler) LessonSlides(c echo.Context) error {
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	slidesPath := data.SlidesHTMLFilePath(lesson.Slides)
	log.Println(slidesPath)
	return c.File(slidesPath)
}

func IsHTMX(e echo.Context) bool {
	// Check for "HX-Request" header
	return e.Request().Header.Get("Hx-Request") != ""
}

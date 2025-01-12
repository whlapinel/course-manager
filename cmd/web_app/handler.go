package main

import (
	"context"
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"gh_static_portfolio/cmd/templates"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"
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
	TermID         RouteParam = "/:term-id"
	CourseID       RouteParam = "/:course-id"
	UnitID         RouteParam = "/:unit-id"
	LessonID       RouteParam = "/:lesson-id"
	ShiftDirection RouteParam = "/:shift-direction" // string param
)

func (p RouteParam) Name() string {
	return string(p[2:])

}

func ParseCourseIDParams(c echo.Context) templates.CourseIDParams {
	var params templates.CourseIDParams
	termID, err := ParseRouteParam(c, TermID)
	if err == nil {
		params.TermID.Valid = true
		params.TermID.Value = termID
	}
	courseID, err := ParseRouteParam(c, CourseID)
	if err == nil {
		params.CourseID.Valid = true
		params.CourseID.Value = courseID
	}
	unitID, err := ParseRouteParam(c, UnitID)
	if err == nil {
		params.UnitID.Valid = true
		params.UnitID.Value = unitID
	}
	lessonID, err := ParseRouteParam(c, LessonID)
	if err == nil {
		params.LessonID.Valid = true
		params.LessonID.Value = lessonID
	}
	return params
}

func CourseIDParam(params templates.CourseIDParams) (int, error) {
	if params.CourseID.Valid {
		return params.CourseID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func UnitIDParam(params templates.CourseIDParams) (int, error) {
	if params.UnitID.Valid {
		return params.UnitID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func LessonIDParam(params templates.CourseIDParams) (int, error) {
	if params.LessonID.Valid {
		return params.LessonID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func TermIDParam(params templates.CourseIDParams) (int, error) {
	if params.TermID.Valid {
		return params.TermID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}

func ParseRouteParam(c echo.Context, param RouteParam) (int, error) {
	return strconv.Atoi(c.Param(param.Name()))

}

func ParseRouteStringParam(c echo.Context, param RouteParam) string {
	return c.Param(param.Name())
}

const (
	Home                 RouteName = "/"
	Terms                RouteName = "/terms"
	Courses              RouteName = Terms + RouteName(TermID) + "/courses"
	Units                RouteName = Courses + RouteName(CourseID) + "/units"
	Calendar             RouteName = Courses + RouteName(CourseID) + "/calendar"
	Lessons              RouteName = Units + RouteName(UnitID) + "/lessons"
	Lesson               RouteName = Lessons + RouteName(LessonID)
	EditLesson           RouteName = Lesson + "/edit"
	Slides               RouteName = Lesson + "/slides"
	EditSlides           RouteName = Slides + "/edit"
	ShiftLessonRouteName RouteName = Lesson + RouteName(ShiftDirection)
)

type RouteHandlerName string

const (
	ShowHome                    = RouteHandlerName(GET + Home)
	ListTerms                   = RouteHandlerName(GET + Terms)
	ListTermCourses             = RouteHandlerName(GET + Courses)
	ListCourseUnits             = RouteHandlerName(GET + Units)
	ShowCourseCalendar          = RouteHandlerName(GET + Calendar)
	ListUnitLessons             = RouteHandlerName(GET + Lessons)
	LessonDetails               = RouteHandlerName(GET + Lesson)
	ShowEditLesson              = RouteHandlerName(GET + EditLesson)
	PostEditLesson              = RouteHandlerName(POST + EditLesson)
	ViewLessonSlides            = RouteHandlerName(GET + Slides)
	ShowEditSlides              = RouteHandlerName(GET + EditSlides)
	PostEditSlides              = RouteHandlerName(POST + EditSlides)
	ShiftLessonRouteHandlerName = RouteHandlerName(POST + ShiftLessonRouteName)
)

func (h CourseHandler) Mount() {
	h.e.GET(string(Home), h.ShowHome).Name = string(ShowHome)
	h.e.GET(string(Terms), h.ListTerms).Name = string(ListTerms)
	h.e.GET(string(Courses), h.ListTermCourses).Name = string(ListTermCourses)
	h.e.GET(string(Units), h.ListCourseUnits).Name = string(ListCourseUnits)
	h.e.GET(string(Calendar), h.ShowCourseCalendar).Name = string(ShowCourseCalendar)
	h.e.GET(string(Lessons), h.ListUnitLessons).Name = string(ListUnitLessons)
	h.e.GET(string(Lesson), h.LessonDetails).Name = string(LessonDetails)
	h.e.GET(string(EditLesson), h.ShowEditLesson).Name = string(ShowEditLesson)
	h.e.POST(string(EditLesson), h.PostEditLesson).Name = string(PostEditLesson)
	h.e.GET(string(Slides), h.ViewLessonSlides).Name = string(ViewLessonSlides)
	h.e.GET(string(EditSlides), h.ShowEditSlides).Name = string(ShowEditSlides)
	h.e.POST(string(EditSlides), h.PostEditSlides).Name = string(PostEditSlides)
	h.e.POST(string(ShiftLessonRouteName), h.ShiftLesson).Name = string(ShiftLessonRouteHandlerName)

}

func (h CourseHandler) ShowHome(c echo.Context) error {
	template := templates.ManagerHomePage()
	return Respond(c, "", template, templates.CourseManagerLayout("Home", template))
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
	return Respond(c, "", template, templates.CourseManagerLayout("Terms", template))
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
	return Respond(c, "", template, templates.CourseManagerLayout("Courses", template))
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
	template := templates.UnitsListTemplate(termID, courseID, units, string(ListUnitLessons), h.e)
	return Respond(c, "", template, templates.CourseManagerLayout("Units", template))
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
	template := templates.CourseCalendarTemplate(*course, string(LessonDetails), string(ShiftLessonRouteHandlerName), h.e)
	return Respond(c, "", template, templates.CourseManagerLayout("Calendar", template))
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
	template := templates.LessonListTemplate(TermID, courseID, unitID, lessons, string(LessonDetails), h.e)
	return Respond(c, "", template, templates.CourseManagerLayout("Lessons", template))
}

func (h CourseHandler) LessonDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("termID:", termID)
	courseID, err := CourseIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	unitID, err := UnitIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	lessonID, err := LessonIDParam(params)
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
	editor := templates.NewLessonEditor(params, string(ShowEditLesson), string(PostEditLesson), h.e, *lesson)
	template := templates.LessonDetailsTemplate(params, lesson, unit, course, string(ViewLessonSlides), string(ShowEditSlides), string(ShowEditLesson), string(PostEditLesson), editor, h.e)
	return Respond(c, "", template, templates.CourseManagerLayout("Lesson Details", template))
}

func (h CourseHandler) ShowEditLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	queryParam := c.QueryParam("field")
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	editor := templates.NewLessonEditor(params, string(ShowEditLesson), string(PostEditLesson), h.e, *lesson)
	respond := func(comp templ.Component) error {
		return Respond(c, h.e.Reverse(string(LessonDetails), params.ToIntSlice()...), comp, nil)
	}
	if queryParam == string(editor.DescriptionInputID) {
		template := editor.LessonDescription(true)
		return respond(template)
	} else if queryParam == string(editor.NameInputID) {
		template := editor.LessonName(true)
		return respond(template)
	}
	errText := "field value is not expected"
	log.Println(errText)
	return fmt.Errorf("%s %s", errText, queryParam)
}

// TODO: maybe the UpdateLesson functions should return an updated lesson instead of having to call GetLesson again
func (h CourseHandler) PostEditLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	lessonID, err := LessonIDParam(params)
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	lessonDescr := c.FormValue(string(templates.EditLessonDescID))
	lessonName := c.FormValue(string(templates.EditLessonNameID))
	log.Println(lessonDescr)
	var editor templates.LessonEditor
	respond := func(comp templ.Component) error {
		return Respond(c, h.e.Reverse(string(LessonDetails), params.ToIntSlice()...), comp, nil)
	}
	updateLesson := func() error {
		updatedLesson, err := h.svc.GetLesson(lessonID)
		if err != nil {
			return err
		}
		*lesson = *updatedLesson
		return nil
	}

	if lessonDescr != "" {
		lesson.Description = lessonDescr
		err := h.svc.UpdateLesson(*lesson)
		if err != nil {
			return err
		}
		updateLesson()
		editor = templates.NewLessonEditor(params, string(ShowEditLesson), string(PostEditLesson), h.e, *lesson)
		return respond(editor.LessonDescription(false))
	}
	if lessonName != "" {
		lesson.Name = lessonName
		err := h.svc.UpdateLesson(*lesson)
		if err != nil {
			return err
		}
		updateLesson()
		editor = templates.NewLessonEditor(params, string(ShowEditLesson), string(PostEditLesson), h.e, *lesson)
		return respond(editor.LessonName(false))
	}
	return nil
}

func (h CourseHandler) ViewLessonSlides(c echo.Context) error {
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	slidesPath := data.NewSlidesHTMLFilePath(lessonID)
	log.Println(slidesPath)
	return c.File(slidesPath)
}

func (h CourseHandler) ShowEditSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	markdownPath := data.NewSlidesMarkdownFilePath(lessonID)
	markdownFile, err := os.Open(markdownPath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer markdownFile.Close()
	bytes, err := io.ReadAll(markdownFile)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(len(bytes), "bytes read")
	log.Println(string(bytes))
	template := templates.EditSlidesTemplate(params, string(bytes), string(PostEditSlides), h.e)
	if !IsHTMX(c) {
		template = templates.CourseManagerLayout("Lessons", template)
	}
	return template.Render(context.Background(), c.Response())
}

func (h CourseHandler) PostEditSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	log.Println(params)
	content := c.FormValue(string(templates.EditSlidesTextAreaID))
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	path := data.NewSlidesMarkdownFilePath(lessonID)
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
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
	template := templates.CourseCalendarTemplate(*course, string(LessonDetails), string(ShiftLessonRouteHandlerName), h.e)
	if !IsHTMX(c) {
		template = templates.CourseManagerLayout("Calendar", template)
	}
	return template.Render(context.Background(), c.Response())
}

func IsHTMX(c echo.Context) bool {
	// Check for "HX-Request" header
	return c.Request().Header.Get("Hx-Request") != ""
}

// This function sends the component passed in. In case request is not an HTMX request
// provide either a redirect or an alternative component (not both). if redirect is empty string
// the alt component will be rendered and sent to the client. If redirect is not empty, the alt
// component will be ignored. Produces error if neither is provided.
func Respond(c echo.Context, redirect string, component, altComponent templ.Component) error {
	if altComponent == nil && redirect == "" {
		return fmt.Errorf("neither redirect or alt component provided in function call")
	}
	if !IsHTMX(c) {
		log.Println("request is NOT an HTMX request:", c.Request().Header.Get("Hx-Request"))
		if redirect != "" {
			log.Println("redirecting to: ", redirect)
			return c.Redirect(http.StatusFound, redirect)
		}
		if altComponent != nil {
			return altComponent.Render(context.Background(), c.Response())

		}
	}
	log.Println("request IS an HTMX request:", c.Request().Header.Get("Hx-Request"))
	return component.Render(context.Background(), c.Response())
}

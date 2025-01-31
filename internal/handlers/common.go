package handlers

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"
	"net/http"
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

type NodeFieldName string

const (
	Name        NodeFieldName = "Name"
	Description NodeFieldName = "Description"
)

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
	StandardID     RouteParam = "/:standard-id"
	AssessmentID   RouteParam = "/:assessment-id"
	ShiftDirection RouteParam = "/:shift-direction" // string param
)

// strips the '/:' off RouteParam
func (p RouteParam) Name() string {
	return string(p[2:])

}

func ParseCourseIDParams(c echo.Context) mt.CourseIDParams {
	var params mt.CourseIDParams
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

func CourseIDParam(params mt.CourseIDParams) (int, error) {
	if params.CourseID.Valid {
		return params.CourseID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func UnitIDParam(params mt.CourseIDParams) (int, error) {
	if params.UnitID.Valid {
		return params.UnitID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func LessonIDParam(params mt.CourseIDParams) (int, error) {
	if params.LessonID.Valid {
		return params.LessonID.Value, nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func TermIDParam(params mt.CourseIDParams) (int, error) {
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

type RouteHandlerName string

func (rhn RouteHandlerName) String() string {
	return string(rhn)
}

type RouteHandler struct {
	RouteName   RouteName
	HandlerName RouteHandlerName
	Method      MethodName
	HandlerFunc echo.HandlerFunc
}

func (h CourseHandler) Mount() {
	h.mountHandlers(h.TermHandlers())
	h.mountHandlers(h.CourseHandlers())
	h.mountHandlers(h.UnitHandlers())
	h.mountHandlers(h.LessonHandlers())
	h.mountHandlers(h.AssessmentHandlers())
	h.mountHandlers(h.CalendarHandlers())
	h.mountHandlers(h.HomeHandlers())

}

func (h CourseHandler) mountHandlers(newHandlers []RouteHandler) {
	for _, handler := range newHandlers {
		switch handler.Method {
		case GET:
			h.e.GET(string(handler.RouteName), handler.HandlerFunc).Name = handler.HandlerName.String()
		case POST:
			h.e.POST(string(handler.RouteName), handler.HandlerFunc).Name = handler.HandlerName.String()
		case DELETE:
			h.e.DELETE(string(handler.RouteName), handler.HandlerFunc).Name = handler.HandlerName.String()
		default:
			log.Fatal("http method in route handler not expected")
		}
	}
}

func IsHTMX(c echo.Context) bool {
	// Check for "HX-Request" header
	return c.Request().Header.Get("Hx-Request") != ""
}

// Sends component passed in. If not an HTMX request will use redirect or an alternative component.
// if redirect is empty the alt component will be sent. If redirect is not empty, altcomponent will be ignored.
// Returns error if neither is provided.
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

func (h CourseHandler) CourseManagerLayout(page templ.Component) templ.Component {
	cml := mt.CourseManagerLayout{
		Page:            page,
		ListTermsRHN:    ListTerms.String(),
		GenerateSiteRHN: GenerateSite.String(),
		SyncSiteRHN:     SyncSite.String(),
		E:               h.e,
	}
	return mt.CourseManagerLayoutComponent(cml)

}

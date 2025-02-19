package handlers

import (
	"context"
	"fmt"
	auth "gh_static_portfolio/internal/authentication"
	"gh_static_portfolio/internal/authorization"
	"gh_static_portfolio/internal/domain"
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

type RouteMethod string

const (
	GET    = "GET: "
	POST   = "POST: "
	PUT    = "PUT: "
	PATCH  = "PATCH: "
	DELETE = "DELETE: "
)

const (
	UserID         RouteParam = "/:user-id"
	TermID         RouteParam = "/:term-id"
	OccasionID     RouteParam = "/:occasion-id"
	CourseID       RouteParam = "/:course-id"
	UnitID         RouteParam = "/:unit-id"
	LessonID       RouteParam = "/:lesson-id"
	StandardID     RouteParam = "/:standard-id"
	AssessmentID   RouteParam = "/:assessment-id"
	ShiftDirection RouteParam = "/:shift-direction" // string param
	Date           RouteParam = "/:date"
)

// strips the '/:' off RouteParam
func (p RouteParam) Name() string {
	return string(p[2:])

}

func ParseCourseIDParams(c echo.Context) mt.CourseIDParams {
	var params mt.CourseIDParams
	userID := ParseRouteStringParam(c, UserID)
	params.UserID.Valid = userID != ""
	params.UserID.Value = userID
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
		return params.CourseID.Value.(int), nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func UnitIDParam(params mt.CourseIDParams) (int, error) {
	if params.UnitID.Valid {
		return params.UnitID.Value.(int), nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func LessonIDParam(params mt.CourseIDParams) (int, error) {
	if params.LessonID.Valid {
		return params.LessonID.Value.(int), nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}
func TermIDParam(params mt.CourseIDParams) (int, error) {
	if params.TermID.Valid {
		return params.TermID.Value.(int), nil
	} else {
		return -1, fmt.Errorf("invalid param")
	}
}

func ParseRouteParam(c echo.Context, param RouteParam) (int, error) {
	return strconv.Atoi(c.Param(param.Name()))
}

func ParseRouteStringParam(c echo.Context, param RouteParam) string {
	log.Println("ParseRouteStringParam: ", c.Param(param.Name()))
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
	h.MountHandlers(h.HomeHandlers())
	h.MountHandlers(h.AuthenticationHandlers())
	ProtectedGroup := h.e.Group("", auth.AddCookieToHeader, auth.JWTMiddlewareProtectedNew(h.e, GetSignin.String()), auth.GetClaims, authorization.Authorization(h.svc.GetUser))
	h.ProtectRoutes(h.UserHomeHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.TermHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.CourseHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.UnitHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.LessonHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.AssessmentHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.CalendarHandlers(), ProtectedGroup)
}

func (h CourseHandler) ProtectRoutes(handlers []RouteHandler, group *echo.Group) {
	for _, handler := range handlers {
		switch handler.Method {
		case GET:
			group.GET(string(handler.RouteName), handler.HandlerFunc).Name = handler.HandlerName.String()
		case POST:
			group.POST(string(handler.RouteName), handler.HandlerFunc).Name = handler.HandlerName.String()
		case DELETE:
			group.DELETE(string(handler.RouteName), handler.HandlerFunc).Name = handler.HandlerName.String()
		default:
			log.Fatal("http method in route handler not expected")
		}

	}
}

func (h CourseHandler) MountHandlers(newHandlers []RouteHandler) {
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

func (h CourseHandler) CourseManagerLayout(page templ.Component, user domain.User) templ.Component {
	cml := mt.CourseManagerLayout{
		Page: page,
		User: user,
		E:    h.e,
	}
	return mt.CourseManagerLayoutComponent(cml)

}

func (h CourseHandler) BreadCrumbs(params mt.CourseIDParams, nodes ...domain.CourseNode) mt.BreadCrumbs {
	var breadCrumbs mt.BreadCrumbs
	for _, node := range nodes {
		if user, ok := node.(domain.User); ok {
			breadCrumbs.User = user
			breadCrumbs.UserDetailsURL = h.e.Reverse(UserHome.String(), params.ToIntSlice()...)
		} else if term, ok := node.(domain.Term); ok {
			breadCrumbs.Term = term
			breadCrumbs.TermDetailsURL = h.e.Reverse(TermDetails.String(), params.ToIntSlice()...)
		} else if course, ok := node.(domain.Course); ok {
			breadCrumbs.Course = course
			breadCrumbs.CourseDetailsURL = h.e.Reverse(CourseDetails.String(), params.ToIntSlice()...)
		} else if unit, ok := node.(domain.Unit); ok {
			breadCrumbs.Unit = unit
			breadCrumbs.UnitDetailsURL = h.e.Reverse(UnitDetails.String(), params.ToIntSlice()...)
		} else if lesson, ok := node.(domain.Lesson); ok {
			breadCrumbs.Lesson = lesson
			breadCrumbs.LessonDetailsURL = h.e.Reverse(LessonDetails.String(), params.ToIntSlice()...)
		}
	}
	return breadCrumbs

}

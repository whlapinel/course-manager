package handlers

import (
	"context"
	"fmt"
	auth "gh_static_portfolio/internal/authentication"
	"gh_static_portfolio/internal/authorization"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"
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

type RoutePath string

type MethodName string

type RouteParam string

type NodeFieldName string

const (
	Name        NodeFieldName = "Name"
	Number      NodeFieldName = "Number"
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

type NodeParam interface {
	int | string
}

func ParseNodePath(c echo.Context) (domain.NodePath, error) {
	parsingError := func(param RouteParam, err error) (domain.NodePath, error) {
		return domain.NodePath{}, fmt.Errorf("error parsing %s: %v", param, err)
	}
	userID := ParseRouteStringParam(c, UserID)
	termID, err := ParseRouteParam(c, TermID)
	if err != nil {
		return parsingError(TermID, err)
	}
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		return parsingError(CourseID, err)
	}
	unitID, err := ParseRouteParam(c, UnitID)
	if err != nil {
		return parsingError(UnitID, err)
	}
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		return parsingError(LessonID, err)
	}
	return domain.NodePath{
		UserID:   userID,
		TermID:   termID,
		CourseID: courseID,
		UnitID:   unitID,
		LessonID: lessonID,
	}, nil

}

func ParseRouteParam(c echo.Context, param RouteParam) (int, error) {
	val := c.Param(param.Name())
	if val == "" {
		return 0, nil
	}
	return strconv.Atoi(val)
}

func ParseRouteStringParam(c echo.Context, param RouteParam) string {
	return c.Param(param.Name())
}

type RouteHandlerName string

func (rhn RouteHandlerName) String() string {
	return string(rhn)
}

type RouteHandler struct {
	RoutePath   RoutePath
	HandlerName RouteHandlerName
	Method      MethodName
	HandlerFunc echo.HandlerFunc
}

type router struct {
	svc          service.CourseService
	app          *echo.Echo
	params       domain.NodePath
	nodes        domain.Nodes
	emptyNodeSet []EmptyNode
}

func (h CourseHandler) Mount() {
	// h.MountHandlers(h.HomeHandlers())
	// h.MountHandlers(h.AuthenticationHandlers())
	ProtectedGroup := h.e.Group("", auth.AddCookieToHeader, auth.JWTMiddlewareProtectedNew(h.e, GetSignin.String()), auth.GetClaims, authorization.Authorization(h.svc.GetUser))
	// h.ProtectRoutes(h.UserHomeHandlers(), ProtectedGroup)
	// h.ProtectRoutes(h.TermHandlers(), ProtectedGroup)
	// h.ProtectRoutes(h.CourseHandlers(), ProtectedGroup)
	// h.ProtectRoutes(h.UnitHandlers(), ProtectedGroup)
	// h.ProtectRoutes(h.LessonHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.AssessmentHandlers(), ProtectedGroup)
	h.ProtectRoutes(h.CalendarHandlers(), ProtectedGroup)
}

func (h CourseHandler) ProtectRoutes(handlers []RouteHandler, group *echo.Group) {
	for _, handler := range handlers {
		switch handler.Method {
		case GET:
			group.GET(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case POST:
			group.POST(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case DELETE:
			group.DELETE(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		default:
			log.Fatal("http method in route handler not expected")
		}

	}
}

func (h CourseHandler) MountHandlers(newHandlers []RouteHandler) {
	for _, handler := range newHandlers {
		switch handler.Method {
		case GET:
			h.e.GET(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case POST:
			h.e.POST(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case DELETE:
			h.e.DELETE(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		default:
			log.Fatal("http method in route handler not expected")
		}
	}
}

func Mount(svc service.CourseService, app *echo.Echo) {
	MountHandlers(AuthHandlers(svc, app), app)
	MountHandlers(HomeHandlers(svc, app), app)
	ProtectedGroup := app.Group("", auth.AddCookieToHeader, auth.JWTMiddlewareProtectedNew(app, GetSignin.String()), auth.GetClaims, authorization.Authorization(svc.GetUser))
	// ProtectRoutes(h.UserHomeHandlers(), ProtectedGroup)
	ProtectRoutes(UserHandlers(svc, app), ProtectedGroup)
	ProtectRoutes(TermHandlers(svc, app), ProtectedGroup)
	ProtectRoutes(CourseHandlers(svc, app), ProtectedGroup)
	ProtectRoutes(UnitHandlers(svc, app), ProtectedGroup)
	ProtectRoutes(LessonHandlers(svc, app), ProtectedGroup)

	// ProtectRoutes(h.CourseHandlers(), ProtectedGroup)
	// ProtectRoutes(h.UnitHandlers(), ProtectedGroup)
	// ProtectRoutes(h.LessonHandlers(), ProtectedGroup)
	// ProtectRoutes(h.AssessmentHandlers(), ProtectedGroup)
	// ProtectRoutes(h.CalendarHandlers(), ProtectedGroup)
}

func ProtectRoutes(handlers []RouteHandler, group *echo.Group) {
	for _, handler := range handlers {
		switch handler.Method {
		case GET:
			group.GET(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case POST:
			group.POST(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case DELETE:
			group.DELETE(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		default:
			log.Fatal("http method in route handler not expected")
		}

	}
}

func MountHandlers(newHandlers []RouteHandler, echo *echo.Echo) {
	for _, handler := range newHandlers {
		switch handler.Method {
		case GET:
			echo.GET(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case POST:
			echo.POST(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
		case DELETE:
			echo.DELETE(string(handler.RoutePath), handler.HandlerFunc).Name = handler.HandlerName.String()
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
		Page:       page,
		User:       user,
		E:          h.e,
		SigninURL:  h.e.Reverse(GetSignin.String()),
		SignupURL:  h.e.Reverse(GetSignup.String()),
		SignoutURL: h.e.Reverse(PostSignout.String()),
	}
	return cml.Component()

}

func CourseManagerLayout(router *echo.Echo, page templ.Component, user domain.User) templ.Component {
	cml := mt.CourseManagerLayout{
		Page:       page,
		User:       user,
		E:          router,
		SigninURL:  router.Reverse(GetSignin.String()),
		SignupURL:  router.Reverse(GetSignup.String()),
		SignoutURL: router.Reverse(PostSignout.String()),
	}
	return cml.Component()

}

func AddParams(path domain.NodePath, params ...any) []any {
	paramSlice := path.ToSlice()
	paramSlice = append(paramSlice, params...)
	return paramSlice
}

func BreadCrumbs(router *echo.Echo, params domain.NodePath, nodes ...domain.CourseNode) mt.BreadCrumbs {
	var breadCrumbs mt.BreadCrumbs
	for _, node := range nodes {
		if user, ok := node.(domain.User); ok {
			breadCrumbs.User = user
			breadCrumbs.UserDetailsURL = router.Reverse(string(ShowNodeDetailsRHN(EmptyNodesUser...)), params.ToSlice()...)
		} else if term, ok := node.(domain.Term); ok {
			breadCrumbs.Term = term
			breadCrumbs.TermDetailsURL = router.Reverse(string(ShowNodeDetailsRHN(EmptyNodesTerm...)), params.ToSlice()...)
		} else if course, ok := node.(domain.Course); ok {
			breadCrumbs.Course = course
			breadCrumbs.CourseDetailsURL = router.Reverse(string(ShowNodeDetailsRHN(EmptyNodesCourse...)), params.ToSlice()...)
		} else if unit, ok := node.(domain.Unit); ok {
			breadCrumbs.Unit = unit
			breadCrumbs.UnitDetailsURL = router.Reverse(string(ShowNodeDetailsRHN(EmptyNodesUnit...)), params.ToSlice()...)
		} else if lesson, ok := node.(domain.Lesson); ok {
			breadCrumbs.Lesson = lesson
			breadCrumbs.LessonDetailsURL = router.Reverse(string(ShowNodeDetailsRHN(EmptyNodesLesson...)), params.ToSlice()...)
		}
	}
	return breadCrumbs

}

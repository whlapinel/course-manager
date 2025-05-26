package web

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/features/user"
	"log"
	"net/http"
	"net/url"

	"slices"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type RouteParam string

// strips the '/:' off RouteParam
func (p RouteParam) Name() string {
	return string(p[2:])
}

type AddParams func(params ...any) string

type Reverse func(name string, params ...any) string

func URLFunc(rhn HandlerName, reverse Reverse, params ...any) AddParams {
	return func(additional ...any) string {
		// Create a new slice with the original params
		full := slices.Clone(params)
		full = append(full, additional...)
		return reverse(string(rhn), full...)
	}
}

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

// returns the page or fragment content with base layout
type LayoutFunc func(templ.Component) templ.Component

func IsHTMX(c echo.Context) bool {
	// Check for "HX-Request" header
	return c.Request().Header.Get("Hx-Request") != ""
}

func EmptyUser() user.User {
	return user.User{
		Picture: AssetsURLFunc("signin.png"),
	}
}

func AssetsURLFunc(path ...string) string {
	URL, err := url.JoinPath("/dist", path...)
	if err != nil {
		return "error joining paths"
	}
	return URL
}

type RouteHandler struct {
	Method
	RoutePath
	HandlerName
	echo.HandlerFunc
}

// Constructor for convenience
func NewRouteHandler(method Method, path RoutePath, name HandlerName, handlerFunc echo.HandlerFunc) RouteHandler {
	return RouteHandler{method, path, name, handlerFunc}
}

type RoutePath string

type HandlerName string

func (name HandlerName) String() string {
	return string(name)
}

func RegisterRoute(group *echo.Group, handler RouteHandler) error {
	switch handler.Method {
	case GET:
		group.GET(string(handler.RoutePath), handler.HandlerFunc).Name = string(handler.HandlerName)
	case POST:
		group.POST(string(handler.RoutePath), handler.HandlerFunc).Name = string(handler.HandlerName)
	case DELETE:
		group.DELETE(string(handler.RoutePath), handler.HandlerFunc).Name = string(handler.HandlerName)
	default:
		return fmt.Errorf("unsupported HTTP method: %s", handler.Method)
	}
	return nil
}

type Method string

const (
	GET    = http.MethodGet
	POST   = http.MethodPost
	DELETE = http.MethodDelete
	PUT    = http.MethodPut
	PATCH  = http.MethodPatch
)

// Route Parameter
const (
	ID = "/:id"
)

func NewHandlerName(method Method, path RoutePath) HandlerName {
	return HandlerName(string(method) + ": " + string(path))
}

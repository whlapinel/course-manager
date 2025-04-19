package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

type RouteParam string

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

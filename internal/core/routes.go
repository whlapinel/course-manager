package core

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	Method string
	RoutePath
	HandlerName
	echo.HandlerFunc
}

// Constructor for convenience
func NewRouteHandler(method string, path RoutePath, name HandlerName, handlerFunc echo.HandlerFunc) RouteHandler {
	return RouteHandler{method, path, name, handlerFunc}
}

type RoutePath string

type HandlerName string

func RegisterRoute(group *echo.Group, handler RouteHandler, middleWare ...echo.MiddlewareFunc) error {
	switch handler.Method {
	case http.MethodGet:
		group.GET(string(handler.RoutePath), handler.HandlerFunc, middleWare...).Name = string(handler.HandlerName)
	case http.MethodPost:
		group.POST(string(handler.RoutePath), handler.HandlerFunc, middleWare...).Name = string(handler.HandlerName)
	case http.MethodDelete:
		group.DELETE(string(handler.RoutePath), handler.HandlerFunc, middleWare...).Name = string(handler.HandlerName)
	default:
		return fmt.Errorf("unsupported HTTP method: %s", handler.Method)
	}
	return nil
}

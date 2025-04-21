package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service *services.UserService
	reverse web.Reverse
}

func NewUserHandler(service *services.UserService, e *echo.Echo) *userHandler {
	return &userHandler{
		service: service,
		reverse: e.Reverse,
	}
}

func RegisterUserRoutes(group *echo.Group, h *userHandler) {
	for _, handler := range userRouteHandlers(h) {
		web.RegisterRoute(group, handler)
	}
}

func userRouteHandlers(h *userHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Terms, HandlerName: routes.GetTerms, HandlerFunc: h.listTerms},
	}
}

func (h *userHandler) listTerms(c echo.Context) error {
	log.Println("UserHandler.listTerms running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	userDTO, err := h.service.ListTerms(path.UserID)
	if err != nil {
		return err
	}
	nodePage := managertemplates.NodeListPage{
		ParentNode:       userDTO,
		Children:         userDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetTerm, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetCourses, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteTerm, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewTerm.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetUser.String(), path.ToSlice()...),
		BreadCrumbsData: managertemplates.BreadCrumbs{
			User:           userDTO,
			UserDetailsURL: h.reverse(routes.GetUser.String(), path.ToSlice()...),
		},
	}
	component := managertemplates.TermsListPage{
		ShowTermCalendarURL: web.URLFunc(routes.GetTermCalendar, h.reverse, path.ToSlice()...),
		NodeListPage:        nodePage,
	}.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

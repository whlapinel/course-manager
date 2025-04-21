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

type termHandler struct {
	service *services.TermService
	reverse web.Reverse
}

func NewTermHandler(service *services.TermService, reverse web.Reverse) *termHandler {
	return &termHandler{
		service: service,
		reverse: reverse,
	}
}

func RegisterTermRoutes(group *echo.Group, h *termHandler) error {
	for _, handler := range termRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func termRouteHandlers(h *termHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Terms, HandlerName: routes.GetTerms, HandlerFunc: h.listCourses},
	}
}

func (h *termHandler) listCourses(c echo.Context) error {
	log.Println("TermHandler.listTerms running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	termDTO, err := h.service.ListCourses(path.TermID)
	if err != nil {
		return err
	}
	nodePage := managertemplates.NodeListPage{
		ParentNode:       termDTO,
		Children:         termDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetTerm, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetCourses, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteTerm, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewTerm.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		BreadCrumbsData: managertemplates.BreadCrumbs{
			Term:           termDTO,
			TermDetailsURL: h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		},
	}
	component := managertemplates.TermsListPage{
		ShowTermCalendarURL: web.URLFunc(routes.GetTermCalendar, h.reverse, path.ToSlice()...),
		NodeListPage:        nodePage,
	}.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

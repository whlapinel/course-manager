package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type unitHandler struct {
	service *services.UnitService
	reverse web.Reverse
}

func NewUnitHandler(service *services.UnitService, reverse web.Reverse) *unitHandler {
	return &unitHandler{
		service: service,
		reverse: reverse,
	}
}

func RegisterUnitRoutes(group *echo.Group, h *unitHandler) error {
	for _, handler := range unitRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func unitRouteHandlers(h *unitHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Lessons, HandlerName: routes.GetLessons, HandlerFunc: h.listLessons},
	}
}

func (h *unitHandler) listLessons(c echo.Context) error {
	log.Println("UnitHandler.listLessons running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	unitDTO, err := h.service.ListLessons(path.UnitID)
	if err != nil {
		return err
	}
	component := managertemplates.NodeListPage{
		ParentNode:       unitDTO,
		Children:         unitDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetLesson, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetLessons, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteLesson, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewLesson.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		BreadCrumbsData: managertemplates.BreadCrumbs{
			Nodes: templates.Nodes{
				Unit: unitDTO,
			},
			UnitDetailsURL: h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		},
	}.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

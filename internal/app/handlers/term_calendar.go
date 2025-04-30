package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type termCalendarHandler struct {
	service     *services.TermCalendarService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewTermCalHandler(svc *services.TermCalendarService, nodeService *services.NodeService, reverse web.Reverse) *termCalendarHandler {
	return &termCalendarHandler{
		service:     svc,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterTermCalRoutes(group *echo.Group, h *termCalendarHandler) error {
	for _, handler := range termCalRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func termCalRouteHandlers(h *termCalendarHandler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.TermCalendar, routes.GetTermCalendar, h.showTermCalendar),
	}
}

func (h *termCalendarHandler) showTermCalendar(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	page := managertemplates.TermCalendar{}
	return web.Respond(c, "", page.Component(), managertemplates.BaseLayout(h.reverse, page.Component(), nodes.User.(dto.User)))
}

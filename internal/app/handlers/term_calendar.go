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
		web.NewRouteHandler(web.GET, routes.TermOccasions, routes.CreateTermOccasion, h.createOccasion),
		web.NewRouteHandler(web.GET, routes.TermOccasion, routes.ShowEditTermOccasion, h.showEditOccasion),
		web.NewRouteHandler(web.POST, routes.TermOccasion, routes.PostEditTermOccasion, h.postEditOccasion),
		web.NewRouteHandler(web.DELETE, routes.TermOccasion, routes.DeleteTermOccasion, h.deleteOccasion),
	}
}

func (h *termCalendarHandler) deleteOccasion(c echo.Context) error {
	panic("not implemented")
}
func (h *termCalendarHandler) postEditOccasion(c echo.Context) error {
	panic("not implemented")
}
func (h *termCalendarHandler) showEditOccasion(c echo.Context) error {
	panic("not implemented")
}

func (h *termCalendarHandler) createOccasion(c echo.Context) error {
	panic("not implemented")
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
	term, err := h.service.Term(path.TermID)
	if err != nil {
		return err
	}
	dates, err := h.service.CalendarDates(path.TermID)
	if err != nil {
		return err
	}
	page := managertemplates.TermCalendar{
		Term:                term,
		ListTermsURL:        h.reverse(routes.GetTerms.String(), path.ToSlice()...),
		TermDetailsURL:      h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		CreateOccasionURL:   h.reverse(routes.CreateTermOccasion.String(), path.ToSlice()...),
		GetEditOccasionURL:  web.URLFunc(routes.ShowEditTermOccasion, h.reverse, path.ToSlice()...),
		PostEditOccasionURL: web.URLFunc(routes.PostEditTermOccasion, h.reverse, path.ToSlice()...),
		DeleteOccasionURL:   web.URLFunc(routes.DeleteTermOccasion, h.reverse, path.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(nodes, path, h.reverse),
		CalendarDates:       dates,
	}
	component := page.Component()
	return web.Respond(c, "", component, managertemplates.BaseLayout(h.reverse, component, nodes.User.(dto.User)))
}

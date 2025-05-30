package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
	"gh_static_portfolio/internal/features/termoccasion"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type termCalendarHandler struct {
	service     *services.TermCalendarService
	nodeService *services.NodeService
	occasions   *termoccasion.Service
	reverse     web.Reverse
}

func NewTermCalHandler(
	svc *services.TermCalendarService,
	nodeService *services.NodeService,
	occasions *termoccasion.Service,
	reverse web.Reverse,
) *termCalendarHandler {
	return &termCalendarHandler{
		service:     svc,
		nodeService: nodeService,
		occasions:   occasions,
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
		web.NewRouteHandler(web.POST, routes.TermOccasions, routes.CreateTermOccasion, h.createOccasion),
		web.NewRouteHandler(web.GET, routes.TermOccasion, routes.ShowEditTermOccasion, h.showEditOccasion),
		web.NewRouteHandler(web.POST, routes.TermOccasion, routes.PostEditTermOccasion, h.postEditOccasion),
		web.NewRouteHandler(web.DELETE, routes.TermOccasion, routes.DeleteTermOccasion, h.deleteOccasion),
	}
}

func (h *termCalendarHandler) deleteOccasion(c echo.Context) error {
	occasionIDParam := c.Param("occasion-id")
	occasionID, err := strconv.Atoi(occasionIDParam)
	if err != nil {
		return err
	}
	err = h.occasions.Delete(occasionID)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}
func (h *termCalendarHandler) postEditOccasion(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	occasionIDParam := c.Param("occasion-id")
	occasionID, err := strconv.Atoi(occasionIDParam)
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	err = h.occasions.Update(occasionID, name)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetTermCalendar.String(), info.NodePath.ToSlice()...))

}
func (h *termCalendarHandler) showEditOccasion(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	occasionIDParam := c.Param("occasion-id")
	occasionID, err := strconv.Atoi(occasionIDParam)
	if err != nil {
		return err
	}
	occ, err := h.occasions.ByID(occasionID)
	if err != nil {
		return err
	}
	component := calendarviews.OccasionEditor{
		Occasion:            occ,
		IsEditing:           true,
		PostEditOccasionURL: web.URLFunc(routes.PostEditTermOccasion, h.reverse, info.NodePath.ToSlice(occasionID)...)(),
	}.Component()
	return web.Respond(c, h.reverse(routes.GetTermCalendar.String(), info.NodePath.ToSlice()...), component, nil)
}

func (h *termCalendarHandler) createOccasion(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	dateParam := c.FormValue("date")
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	_, err = h.occasions.Create(date, name, info.TermID)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetTermCalendar.String(), info.NodePath.ToSlice()...))

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
	page := calendarviews.TermCalendar{
		Term:                term,
		ListTermsURL:        h.reverse(routes.GetTerms.String(), path.ToSlice()...),
		TermDetailsURL:      h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		CreateOccasionURL:   h.reverse(routes.CreateTermOccasion.String(), path.ToSlice()...),
		GetEditOccasionURL:  web.URLFunc(routes.ShowEditTermOccasion, h.reverse, path.ToSlice()...),
		PostEditOccasionURL: web.URLFunc(routes.PostEditTermOccasion, h.reverse, path.ToSlice()...),
		DeleteOccasionURL:   web.URLFunc(routes.DeleteTermOccasion, h.reverse, path.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(nodes, path, h.reverse),
		CalendarDates:       dates,
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

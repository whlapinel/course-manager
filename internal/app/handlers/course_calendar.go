package handlers

import (
	"fmt"
	managertemplates "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type courseCalendarHandler struct {
	service     *services.CourseCalendarService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewCourseCalHandler(svc *services.CourseCalendarService, nodeService *services.NodeService, reverse web.Reverse) *courseCalendarHandler {
	return &courseCalendarHandler{
		service:     svc,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterCourseCalRoutes(group *echo.Group, h *courseCalendarHandler) error {
	for _, handler := range courseCalRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func courseCalRouteHandlers(h *courseCalendarHandler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.CourseCalendar, routes.GetCourseCalendar, h.getCourseCalendar),
		web.NewRouteHandler(web.POST, routes.ShiftLesson, routes.PostShiftLesson, h.shiftLesson),
	}
}

func (h *courseCalendarHandler) shiftLesson(c echo.Context) error {
	log.Println("shiftLesson running")
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return fmt.Errorf("error retrieving node info: %w", err)
	}
	dateString := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateString)
	direction := c.Param("shift-direction")
	if err != nil {
		return err
	}
	switch direction {
	case dto.Left.String():
		err = h.service.ShiftLesson(date, info.LessonID, info.TermID, dto.Left)
	case dto.Right.String():
		err = h.service.ShiftLesson(date, info.LessonID, info.TermID, dto.Right)
	default:
		return fmt.Errorf("direction not defined or value not expected: %s", direction)
	}
	if err != nil {
		return fmt.Errorf("error in handler's call to service.ShiftLesson: %w", err)
	}
	return c.Redirect(303, h.reverse(routes.GetCourseCalendar.String(), info.NodePath.ToSlice()...))
}

func (h *courseCalendarHandler) getCourseCalendar(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	datesMap, err := h.service.CalendarDates(path.CourseID)
	if err != nil {
		return err
	}
	log.Println("len(datesMap):", len(datesMap))
	page := calendarviews.CourseCalendar{
		Course:                nodes.Course.(dto.Course),
		Term:                  nodes.Term.(dto.Term),
		TermDetailsURL:        h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		CourseDetailsURL:      h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		ListTermCoursesURL:    h.reverse(routes.GetCourses.String(), path.ToSlice()...),
		LessonDetailsFunc:     h.URLFunc(routes.GetLesson, path.ToSlice()...),
		ShiftLessonFunc:       h.URLFunc(routes.PostShiftLesson, path.ToSlice()...),
		CreateOccasionFunc:    h.URLFunc(routes.CreateTermOccasion, path.ToSlice()...),
		ShowAddLessonDateFunc: h.URLFunc(routes.GetAddLessonDate, path.ToSlice()...),
		RemoveLessonDateFunc:  h.URLFunc(routes.DeleteLessonDate, path.ToSlice()...),
		CalendarDates:         datesMap,
		BreadCrumbsData:       h.BreadCrumbs(nodes, path),
		CourseManagerLayout:   BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *courseCalendarHandler) BreadCrumbs(nodes ports.Nodes, path routes.NodePath) managertemplates.BreadCrumbs {
	return BreadCrumbs(nodes, path, h.reverse)
}

func (h *courseCalendarHandler) URLFunc(name web.HandlerName, params ...any) web.AddParams {
	return web.URLFunc(name, h.reverse, params...)
}

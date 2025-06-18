package handlers

import (
	"fmt"
	managertemplates "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
	"gh_static_portfolio/internal/features/courseoccasion"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type courseCalendarHandler struct {
	service     *services.CourseCalendarService
	occasions   *courseoccasion.Service
	nodeService *services.NodeService
	lessons     *services.LessonService
	units       *services.UnitService
	reverse     web.Reverse
}

func NewCourseCalHandler(
	svc *services.CourseCalendarService,
	occasions *courseoccasion.Service,
	nodeService *services.NodeService,
	lessons *services.LessonService,
	units *services.UnitService,
	reverse web.Reverse,
) *courseCalendarHandler {
	return &courseCalendarHandler{
		service:     svc,
		occasions:   occasions,
		nodeService: nodeService,
		lessons:     lessons,
		units:       units,
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
		web.NewRouteHandler(web.GET, routes.CourseMonthCalendar, routes.GetCourseMonthCalendar, h.getCourseCalendar),
		web.NewRouteHandler(web.POST, routes.ShiftLesson, routes.PostShiftLesson, h.shiftLesson),
		web.NewRouteHandler(web.GET, routes.DateUnits, routes.GetDateUnits, h.getDateUnits),
		web.NewRouteHandler(web.GET, routes.DateLessons, routes.GetDateLessons, h.getDateLessons),
		web.NewRouteHandler(web.POST, routes.DateLesson, routes.PostAddLessonDate, h.postDateLesson),
		web.NewRouteHandler(web.DELETE, routes.DateLesson, routes.DeleteLessonDate, h.deleteLessonDate),

		web.NewRouteHandler(web.POST, routes.CourseOccasions, routes.CreateCourseOccasion, h.createOccasion),
		web.NewRouteHandler(web.GET, routes.CourseOccasion, routes.ShowEditCourseOccasion, h.showEditOccasion),
		web.NewRouteHandler(web.POST, routes.CourseOccasion, routes.PostEditCourseOccasion, h.postEditOccasion),
		web.NewRouteHandler(web.DELETE, routes.CourseOccasion, routes.DeleteCourseOccasion, h.deleteOccasion),
	}
}

func (h *courseCalendarHandler) createOccasion(c echo.Context) error {
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
	_, err = h.occasions.Create(date, name, info.CourseID)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetCourseCalendar.String(), info.NodePath.ToSlice()...))
}

func (h *courseCalendarHandler) showEditOccasion(c echo.Context) error {
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
		PostEditOccasionURL: web.URLFunc(routes.PostEditCourseOccasion, h.reverse, info.NodePath.ToSlice(occasionID)...)(),
	}.Component()
	return web.Respond(c, h.reverse(routes.GetCourseCalendar.String(), info.NodePath.ToSlice()...), component, nil)
}

func (h *courseCalendarHandler) postEditOccasion(c echo.Context) error {
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
	return c.Redirect(303, h.reverse(routes.GetCourseCalendar.String(), info.NodePath.ToSlice()...))
}

func (h *courseCalendarHandler) deleteOccasion(c echo.Context) error {
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

func (h *courseCalendarHandler) deleteLessonDate(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	dateParam := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	err = h.service.DeleteLessonDate(date, info.LessonID, info.TermID)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}

func (h *courseCalendarHandler) postDateLesson(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	dateParam := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	err = h.service.AddLessonToDate(date, info.LessonID, info.TermID)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetCourseCalendar.String(), info.NodePath.ToSlice()...))

}

func (h *courseCalendarHandler) getDateLessons(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	dateParam := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	lessons, err := h.lessons.ByParentID(info.UnitID)
	if err != nil {
		return err
	}
	partial := calendarviews.LessonPicker{
		Date:            date,
		ListUnitsURL:    h.reverse(routes.GetDateUnits.String(), []any{info.UserID, info.TermID, info.CourseID, dateParam}...),
		Lessons:         lessons,
		SelectLessonURL: web.URLFunc(routes.PostAddLessonDate, h.reverse, []any{info.UserID, info.TermID, info.CourseID, dateParam, info.UnitID}...),
	}
	return web.Respond(c, routes.GetCourseCalendar.String(), partial.Component(), nil)

}

func (h *courseCalendarHandler) getDateUnits(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	dateParam := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	units, err := h.units.ByParentID(info.CourseID)
	if err != nil {
		return err
	}
	page := calendarviews.AddLessonToDatePage{
		Date:                date,
		Course:              info.Course.(dto.Course),
		ListLessonsURL:      web.URLFunc(routes.GetDateLessons, h.reverse, info.NodePath.ToSlice(dateParam)...),
		Units:               units,
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
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
	currMonth := date.Month()
	currYear := date.Year()
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
	return c.Redirect(303, h.reverse(routes.GetCourseMonthCalendar.String(), info.UserID, info.TermID, info.CourseID, int(currMonth), currYear))
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
	monthParam := c.Param("month")
	yearParam := c.Param("year")
	if monthParam != "" && yearParam != "" {
		month, err := strconv.Atoi(monthParam)
		if err != nil {
			return err
		}
		year, err := strconv.Atoi(yearParam)
		if err != nil {
			return err
		}
		calWeeks, err := h.service.MonthWeeks(month, year, path.CourseID)
		if err != nil {
			return err
		}
		var presentMonthURL string
		var nextMonthURL string
		var prevMonthURL string
		firstOfCurrentMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		lastDay := nodes.Term.(dto.Term).End
		firstOfNextMonth := firstOfCurrentMonth.AddDate(0, 1, 0)
		nextMonth := firstOfNextMonth.Month()
		nextMonthYear := firstOfNextMonth.Year()
		if !firstOfNextMonth.After(lastDay) {
			nextMonthURL = h.reverse(routes.GetCourseMonthCalendar.String(), path.ToSlice(int(nextMonth), nextMonthYear)...)
		}
		firstDay := nodes.Term.(dto.Term).Start
		lastOfPrevMonth := firstOfCurrentMonth.AddDate(0, 0, -1)
		prevMonth := lastOfPrevMonth.Month()
		prevMonthYear := lastOfPrevMonth.Year()
		if !lastOfPrevMonth.Before(firstDay) {
			prevMonthURL = h.reverse(routes.GetCourseMonthCalendar.String(), path.ToSlice(int(prevMonth), prevMonthYear)...)
		}
		presentYear, presentMonth, presentDate := time.Now().Date()
		today := time.Date(presentYear, presentMonth, presentDate, 0, 0, 0, 0, time.UTC)
		presentMonthURL = h.reverse(routes.GetCourseMonthCalendar.String(), path.ToSlice(int(presentMonth), presentYear)...)
		newPage := calendarviews.CourseMonthCalendar{
			Today:               today,
			Term:                nodes.Term.(dto.Term),
			Month:               time.Month(month),
			Year:                year,
			Weeks:               calWeeks,
			PresentMonthURL:     presentMonthURL,
			NextMonthURL:        nextMonthURL,
			PrevMonthURL:        prevMonthURL,
			RemoveLessonDateURL: web.URLFunc(routes.DeleteLessonDate, h.reverse, path.ToSlice()...),
			ShiftLessonURL:      web.URLFunc(routes.PostShiftLesson, h.reverse, path.ToSlice()...),
			CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
		}
		return Respond(c, newPage)
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
		ShowAddLessonDateFunc: h.URLFunc(routes.GetDateUnits, path.ToSlice()...),
		RemoveLessonDateFunc:  h.URLFunc(routes.DeleteLessonDate, path.ToSlice()...),
		CreateOccasionURL:     h.URLFunc(routes.CreateCourseOccasion, path.ToSlice()...)(),
		GetEditOccasionURL:    h.URLFunc(routes.ShowEditCourseOccasion, path.ToSlice()...),
		PostEditOccasionURL:   h.URLFunc(routes.PostEditCourseOccasion, path.ToSlice()...),
		DeleteOccasionURL:     h.URLFunc(routes.DeleteCourseOccasion, path.ToSlice()...),
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

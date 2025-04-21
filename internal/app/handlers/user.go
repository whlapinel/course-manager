package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/user"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service     *user.Service
	termService *term.Service
	reverse     web.Reverse
}

func NewUserHandler(service *user.Service, termService *term.Service, e *echo.Echo) *UserHandler {
	return &UserHandler{
		service:     service,
		termService: termService,
		reverse:     e.Reverse,
	}
}

func RegisterUserRoutes(group *echo.Group, h *UserHandler) {
	for _, handler := range RouteHandlers(h) {
		web.RegisterRoute(group, handler)
	}
}

func RouteHandlers(h *UserHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Terms, HandlerName: routes.GetTerms, HandlerFunc: h.listTerms},
	}
}

func (h *UserHandler) listTerms(c echo.Context) error {
	log.Println("UserHandler.listTerms running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	user, err := h.service.ByID(path.UserID)
	if err != nil {
		return err
	}
	terms, err := h.termService.TermsByUser(path.UserID)
	if err != nil {
		return err
	}
	var termDTOs []templates.Node
	for _, term := range terms {
		termDTO := dto.Term{
			Term: term,
		}
		termDTOs = append(termDTOs, termDTO)
	}
	userDTO := dto.User{
		User:  user,
		Terms: termDTOs,
	}
	nodePage := managertemplates.NodeListPage{
		ParentNode:       userDTO,
		Children:         termDTOs,
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
	layout := managertemplates.BaseLayout(h.reverse, component, user)
	return web.Respond(c, "", component, layout)
}

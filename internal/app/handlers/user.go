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
	service     *services.UserService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewUserHandler(service *services.UserService, nodeService *services.NodeService, reverse web.Reverse) *userHandler {
	return &userHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterUserRoutes(group *echo.Group, h *userHandler) error {
	for _, handler := range userRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func userRouteHandlers(h *userHandler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.Terms, routes.GetTerms, h.listTerms),
		web.NewRouteHandler(web.POST, routes.GenerateSite, routes.PostGenerateSite, h.generateSite),
	}
}

func (h *userHandler) generateSite(c echo.Context) error {
	panic("not implemented")
}

func (h *userHandler) listTerms(c echo.Context) error {
	log.Println("UserHandler.listTerms running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	log.Println("user ID:", path.UserID)
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
		BreadCrumbsData:  BreadCrumbs(nodes, path, h.reverse),
	}
	component := managertemplates.TermsListPage{
		ShowTermCalendarURL: web.URLFunc(routes.GetTermCalendar, h.reverse, path.ToSlice()...),
		NodeListPage:        nodePage,
	}.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

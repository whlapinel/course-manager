package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	dashboardviews "gh_static_portfolio/internal/app/views/dashboard"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type dashboardHandler struct {
	service     *services.UserService
	terms       *services.TermService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewDashboardHandler(
	service *services.UserService,
	terms *services.TermService,
	nodeService *services.NodeService,
	reverse web.Reverse,
) *dashboardHandler {
	return &dashboardHandler{
		terms:       terms,
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterDashboardRoutes(group *echo.Group, h *dashboardHandler) error {
	for _, handler := range dashboardRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func dashboardRouteHandlers(h *dashboardHandler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.Users, routes.GetUsers, h.redirectToDashboard),
		web.NewRouteHandler(web.GET, routes.User, routes.GetUser, h.showDashboard),
	}
}

// this reads the user's id from the context and redirects to the user's dashboard
func (h *dashboardHandler) redirectToDashboard(c echo.Context) error {
	userID := c.Get("id")
	if userID == "" {
		return fmt.Errorf("userID is nil")
	}
	dashboardURL := h.reverse(routes.GetUser.String(), userID)
	return c.Redirect(302, dashboardURL)
}

func (h *dashboardHandler) showDashboard(c echo.Context) error {
	log.Println("showDashboard running")
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	log.Println("params.UserID", info.UserID)
	user, err := h.service.ByID(info.UserID)
	if err != nil {
		return err
	}
	userDTO := dto.User{
		User: user.User,
	}
	page := dashboardviews.UserHomePage{
		ListTermsURL:        h.reverse(routes.GetTerms.String(), info.NodePath.ToSlice()...),
		User:                userDTO,
		CourseManagerLayout: BaseLayout3(h.reverse, user),
	}
	return Respond(c, page)
}

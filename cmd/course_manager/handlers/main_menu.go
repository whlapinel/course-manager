package handlers

import (
	"gh_static_portfolio/cmd/course_manager/templates"

	"github.com/labstack/echo/v4"
)

type MainMenuHandler interface {
	Mount()
	ShowMenu(c echo.Context) error
}

type mainMenuHandler struct {
	e *echo.Echo
}

func NewMainMenuHandler(e *echo.Echo) MainMenuHandler {
	return mainMenuHandler{e}
}

const (
	MainMenuHandlerShowMenu = "GET: /"
)

func (h mainMenuHandler) Mount() {
	h.e.GET("/", h.ShowMenu).Name = MainMenuHandlerShowMenu
}
func (h mainMenuHandler) ShowMenu(c echo.Context) error {
	component := templates.ManagerMainMenu()
	return RenderTempl(component, c, 200)

}

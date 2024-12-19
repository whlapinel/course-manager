package handlers

import (
	"gh_static_portfolio/cmd/course_manager/services"
	"gh_static_portfolio/cmd/course_manager/templates"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	service services.CourseService
	e       *echo.Echo
}

func NewCourseHandler(service services.CourseService, e *echo.Echo) CourseHandler {
	return CourseHandler{service: service, e: e}
}

const (
	CourseHandlerCreate        RouteName = "POST: /courses"
	CourseHandlerListTemplates RouteName = "GET: /courses"
	CourseHandlerListInstances RouteName = "GET: /courses/:term-id"
	CourseHandlerReadFromCSV   RouteName = "GET: /courses/csv"
	CourseHandlerUpdate        RouteName = "PUT: /courses/:id"
	CourseHandlerDelete        RouteName = "DELETE: /courses/:id"
)

func (h CourseHandler) Mount() {
	nameRoute(h.e.POST("/courses", h.Create), CourseHandlerCreate)
	nameRoute(h.e.GET("/courses", h.ListTemplates), CourseHandlerListTemplates)
}

func (h CourseHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
func (h CourseHandler) ListTemplates(c echo.Context) error {
	courses, err := h.service.GetTemplates()
	if err != nil {
		log.Println("courseHandler List():", err)
		return err
	}
	component := templates.ManageCourseComponent(courses)
	return RenderTempl(component, c, 200)
}

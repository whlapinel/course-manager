package handlers

import (
	"gh_static_portfolio/cmd/course_manager/services"
	"gh_static_portfolio/cmd/course_manager/templates"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseHandler interface {
	Handler
	ListTemplates(c echo.Context) error
	ListInstances(c echo.Context) error
}
type courseHandler struct {
	service services.CourseService
	e       *echo.Echo
}

func NewCourseHandler(service services.CourseService, e *echo.Echo) CourseHandler {
	return courseHandler{service: service, e: e}
}

const (
	CourseHandlerCreate        RouteName = "POST: /courses"
	CourseHandlerListTemplates RouteName = "GET: /courses"
	CourseHandlerListInstances RouteName = "GET: /courses/:term-id"
	CourseHandlerReadFromCSV   RouteName = "GET: /courses/csv"
	CourseHandlerUpdate        RouteName = "PUT: /courses/:id"
	CourseHandlerDelete        RouteName = "DELETE: /courses/:id"
)

func (h courseHandler) Mount() {
	nameRoute(h.e.POST("/courses", h.Create), CourseHandlerCreate)
	nameRoute(h.e.GET("/courses", h.ListTemplates), CourseHandlerListTemplates)
	nameRoute(h.e.GET("/courses/:term-id", h.ListInstances), CourseHandlerListInstances)
	nameRoute(h.e.GET("/courses/csv", h.ReadFromCSV), CourseHandlerReadFromCSV)
	nameRoute(h.e.PUT("/courses/:id", h.Update), CourseHandlerUpdate)
	nameRoute(h.e.DELETE("/courses/:id", h.Delete), CourseHandlerDelete)
}

func (h courseHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseHandler) ListInstances(c echo.Context) error {
	courses, err := h.service.GetInstances()
	if err != nil {
		log.Println("courseHandler List():", err)
		return err
	}
	component := templates.ManageCourseOOPComponent(courses)
	return RenderTempl(component, c, 200)
}

func (h courseHandler) ListTemplates(c echo.Context) error {
	courses, err := h.service.GetTemplates()
	if err != nil {
		log.Println("courseHandler List():", err)
		return err
	}
	component := templates.ManageCourseOOPComponent(courses)
	return RenderTempl(component, c, 200)
}

func (h courseHandler) ReadFromCSV(c echo.Context) error {
	courses, err := h.service.ReadFromCSV()
	if err != nil {
		return err
	}
	for _, instance := range courses {
		err := h.service.CreateInstance(instance)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
	}
	component := templates.ManageCourseOOPComponent(courses)
	return RenderTempl(component, c, 200)

}

func (h courseHandler) Delete(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseHandler) Update(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

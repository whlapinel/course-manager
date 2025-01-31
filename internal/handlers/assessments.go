package handlers

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	Assessments RouteName = Lesson + "/assessments"
	Assessment  RouteName = Assessments + RouteName(AssessmentID)
)

const (
	PostAssessment   = RouteHandlerName(POST + Assessments)
	DeleteAssessment = RouteHandlerName(DELETE + Assessment)
)

func (h CourseHandler) AssessmentHandlers() []RouteHandler {
	return []RouteHandler{
		{Assessments, PostAssessment, POST, h.PostAssessment},
		{Assessment, DeleteAssessment, DELETE, h.DeleteAssessment},
	}
}

func (h CourseHandler) PostAssessment(c echo.Context) error {
	params := ParseCourseIDParams(c)
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	lessonIDParam := form.Get("lesson-id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		return err
	}
	name := form.Get("name")
	instructions := form.Get("instructions")
	categoryParam := form.Get("category")
	category, err := strconv.Atoi(categoryParam)
	if err != nil {
		return err
	}
	assignedParam := form.Get("date-assigned")
	assigned, err := time.Parse(time.DateOnly, assignedParam)
	if err != nil {
		return err
	}
	dueParam := form.Get("date-due")
	due, err := time.Parse(time.DateOnly, dueParam)
	if err != nil {
		return err
	}
	_, err = h.svc.SaveAssessment(service.SaveAssessmentParams{
		Assessment: domain.Assessment{
			LessonID:     lessonID,
			Name:         name,
			Instructions: instructions,
			Category:     domain.AssessmentCategory(category),
			DateAssigned: assigned,
			DateDue:      due,
			Dropped:      false,
		},
	})
	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(LessonDetails.String(), params.ToIntSlice()...))
}

func (h CourseHandler) DeleteAssessment(c echo.Context) error {
	panic("delete assessment not implemented")
}

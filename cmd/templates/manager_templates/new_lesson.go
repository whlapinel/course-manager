package managertemplates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type newLessonForm struct {
	params           CourseIDParams
	postNewLessonRHN string
	e                *echo.Echo
}

func NewLessonForm(params CourseIDParams, postNewLessonRHN string, e *echo.Echo) templ.Component {
	form := newLessonForm{params: params, postNewLessonRHN: postNewLessonRHN, e: e}
	return newLessonFormComponent(form)
}

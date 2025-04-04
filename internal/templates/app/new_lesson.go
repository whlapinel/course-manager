package managertemplates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type newLessonForm struct {
	params           NodePath
	postNewLessonRHN string
	e                *echo.Echo
}

func NewLessonForm(params NodePath, postNewLessonRHN string, e *echo.Echo) templ.Component {
	form := newLessonForm{params: params, postNewLessonRHN: postNewLessonRHN, e: e}
	return newLessonFormComponent(form)
}

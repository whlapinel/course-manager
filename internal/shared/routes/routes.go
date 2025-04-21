package routes

import (
	"fmt"
	"gh_static_portfolio/internal/shared/web"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	UserID         web.RouteParam = "/:user-id"
	TermID         web.RouteParam = "/:term-id"
	OccasionID     web.RouteParam = "/:occasion-id"
	CourseID       web.RouteParam = "/:course-id"
	UnitID         web.RouteParam = "/:unit-id"
	LessonID       web.RouteParam = "/:lesson-id"
	StandardID     web.RouteParam = "/:standard-id"
	AssessmentID   web.RouteParam = "/:assessment-id"
	ShiftDirection web.RouteParam = "/:shift-direction" // string param
	Date           web.RouteParam = "/:date"
)

func ParseNodePath(c echo.Context) (NodePath, error) {
	parsingError := func(param web.RouteParam, err error) (NodePath, error) {
		return NodePath{}, fmt.Errorf("error parsing %s: %v", param, err)
	}
	userID := ParseRouteStringParam(c, UserID)
	termID, err := ParseRouteParam(c, TermID)
	if err != nil {
		return parsingError(TermID, err)
	}
	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		return parsingError(CourseID, err)
	}
	unitID, err := ParseRouteParam(c, UnitID)
	if err != nil {
		return parsingError(UnitID, err)
	}
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		return parsingError(LessonID, err)
	}
	return NodePath{
		UserID:   userID,
		TermID:   termID,
		CourseID: courseID,
		UnitID:   unitID,
		LessonID: lessonID,
	}, nil

}

func ParseRouteParam(c echo.Context, param web.RouteParam) (int, error) {
	val := c.Param(param.Name())
	if val == "" {
		return 0, nil
	}
	return strconv.Atoi(val)
}

func ParseRouteStringParam(c echo.Context, param web.RouteParam) string {
	return c.Param(param.Name())
}

type NodePath struct {
	UserID   string
	TermID   int
	CourseID int
	UnitID   int
	LessonID int
}

type Param interface {
	int | string
}

func (path NodePath) ToSlice(additionalParams ...any) []any {
	var pathSlice []any
	params := []any{path.UserID, path.TermID, path.CourseID, path.UnitID, path.LessonID}
	params = append(params, additionalParams...)
	for _, param := range params {
		switch v := param.(type) {
		case int:
			if v != 0 {
				pathSlice = append(pathSlice, param)
			}
		case string:
			if v != "" {
				pathSlice = append(pathSlice, param)

			}
		}
	}
	return pathSlice
}

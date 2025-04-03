package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	templates "gh_static_portfolio/internal/templates/shared"
	"log"
	"net/url"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ComponentData interface {
	Component() templ.Component
}

const pageElementID ElementID = "page"

type NodePath struct {
	UserID   NodeIDParam
	TermID   NodeIDParam
	CourseID NodeIDParam
	UnitID   NodeIDParam
	LessonID NodeIDParam
}

type NodeIDParam struct {
	Valid bool
	Value interface{}
}

func AddNodeChildIDToParams(params NodePath, childID any) NodePath {
	var newParams NodePath
	if params.UserID.Value.(string) == "" {
		newParams.UserID = NodeIDParam{Value: childID.(string)}
		return newParams
	} else if params.TermID.Value == nil {
		newParams = params
		newParams.TermID = NodeIDParam{Value: childID}
		return newParams
	} else if params.CourseID.Value == nil {
		newParams = params
		newParams.CourseID = NodeIDParam{Value: childID}
		return newParams
	} else if params.UnitID.Value == nil {
		newParams = params
		newParams.UnitID = NodeIDParam{Value: childID}
		return newParams
	} else if params.LessonID.Value == nil {
		newParams = params
		newParams.LessonID = NodeIDParam{Value: childID}
		return newParams
	}
	return params
}

// converts params into a slice of interfaces
func (params NodePath) ToSlice(additionalParams ...interface{}) []interface{} {
	var base []interface{}
	paramSlice := []interface{}{
		params.UserID.Value,
		params.TermID.Value,
		params.CourseID.Value,
		params.UnitID.Value,
		params.LessonID.Value,
	}
	for _, param := range paramSlice {
		if param != nil {
			base = append(base, param)
		}
	}
	return append(base, additionalParams...)
}

type ElementID string

// simply prefixes with '#'
func (i ElementID) Selector() string {
	return string("#" + i)
}

func (i ElementID) String() string {
	return string(i)
}

const (
	EditSlidesContainerID ElementID = "slides-editor-container"
	EditSlidesTextAreaID  ElementID = "slides-editor-text-area"
)

type ShiftButton struct {
	Params         domain.NodePath
	Direction      domain.CalendarDirection
	TermID         int
	CourseID       int
	ShiftLessonURL string
	e              *echo.Echo
}

func (data ShiftButton) Component() templ.Component {
	return ShiftButtonComponent(data)
}

func AddQueryParam(path, key, value string) string {
	u, err := url.Parse(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	query := u.Query()
	query.Set(key, value)
	u.RawQuery = query.Encode()
	log.Println(u.String())
	return u.String()
}

type EditField struct {
	Params           domain.NodePath
	FieldName        string
	Content          string
	GetEditFieldURL  string
	PostEditFieldURL string
	InputComponent   templ.Component
	IsEdit           bool
}

func FieldContainerID(fieldName string) string {
	return templates.KebabCase(fieldName) + "-container"
}

func FieldInputID(fieldName string) string {
	return templates.KebabCase(fieldName)
}

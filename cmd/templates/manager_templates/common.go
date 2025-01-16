package managertemplates

import (
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/templates"
	"log"
	"net/url"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const pageElementID ElementID = "page"

type CourseIDParams struct {
	TermID   IDParam
	CourseID IDParam
	UnitID   IDParam
	LessonID IDParam
}

type IDParam struct {
	Valid bool
	Value int
}

func (params CourseIDParams) ToIntSlice() []interface{} {
	return []interface{}{params.TermID.Value, params.CourseID.Value, params.UnitID.Value, params.LessonID.Value}
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

type ShiftButtonFactory struct {
	TermID                      int
	CourseID                    int
	ShiftLessonRouteHandlerName string
	e                           *echo.Echo
}

func NewShiftButtonFactory(course domain.Course, shiftLessonHandlerName string, e *echo.Echo) ShiftButtonFactory {
	return ShiftButtonFactory{
		TermID:                      course.Term.ID,
		CourseID:                    course.ID,
		ShiftLessonRouteHandlerName: shiftLessonHandlerName,
		e:                           e,
	}

}

func (sbf ShiftButtonFactory) ShiftButton(unitID, lessonID int, cd domain.CalendarDirection) templ.Component {
	return ShiftButtonTemplate(sbf.TermID, sbf.CourseID, unitID, lessonID, sbf.ShiftLessonRouteHandlerName, cd, sbf.e)
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

type HXButton struct {
	Method     HXMethod
	URL        string
	HxTargetID string
	PushURL    bool
	HxSwap     HxSwap
}

type HxSwap string

const (
	AfterEnd HxSwap = "afterend"
)

func (button HXButton) Component() templ.Component {
	return hxButton(button)
}
func NewHXButton(method HXMethod, hxSwap HxSwap, url, hxTargetID string, pushURL bool) HXButton {
	return HXButton{
		Method:     method,
		URL:        url,
		HxTargetID: hxTargetID,
		PushURL:    pushURL,
		HxSwap:     hxSwap,
	}
}

type HXMethod string

const (
	HxGet    = "hx-get"
	HxPost   = "hx-post"
	HxDelete = "hx-delete"
)

type EditField struct {
	Params           CourseIDParams
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

package managertemplates

import (
	"gh_static_portfolio/cmd/domain"
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

func (i ElementID) Selector() string {
	return string("#" + i)
}

func (i ElementID) String() string {
	return string(i)
}

const (
	EditSlidesContainerID ElementID = "editor-container"
	EditSlidesTextAreaID  ElementID = "editor-text-area"
)

type LessonEditor struct {
	Params                 CourseIDParams
	GetEditLessonRHN       string // route handler name to retrieve form upon clicking edit
	PostEditLessonRHN      string // route handler name to post edits upon clicking submit
	E                      *echo.Echo
	Lesson                 domain.Lesson
	DescriptionContainerID ElementID
	DescriptionInputID     ElementID
	NameContainerID        ElementID
	NameInputID            ElementID
}

const (
	EditLessonDescID ElementID = "lesson-desc-input"
	EditLessonNameID ElementID = "lesson-name-input"
)

func NewLessonEditor(params CourseIDParams, getEdit, postEdit string, e *echo.Echo, lesson domain.Lesson) LessonEditor {
	return LessonEditor{
		Params:                 params,
		GetEditLessonRHN:       getEdit,
		PostEditLessonRHN:      postEdit,
		E:                      e,
		Lesson:                 lesson,
		DescriptionContainerID: "lesson-desc-container",
		DescriptionInputID:     EditLessonDescID,
		NameContainerID:        "lesson-name-container",
		NameInputID:            EditLessonNameID,
	}
}

func (editor LessonEditor) LessonDescription(isEdit bool) templ.Component {
	input := DescriptionInput(editor.DescriptionInputID, editor.Lesson.Description)
	return LessonFieldTemplate(
		editor.Lesson,
		editor.Params,
		editor.GetEditLessonRHN,
		editor.PostEditLessonRHN,
		editor.DescriptionContainerID,
		editor.DescriptionInputID,
		editor.E,
		isEdit,
		input,
		editor.Lesson.Description,
	)
}

func (editor LessonEditor) LessonName(isEdit bool) templ.Component {
	input := NameInput(editor.NameInputID, editor.Lesson.Name)
	return LessonFieldTemplate(
		editor.Lesson,
		editor.Params,
		editor.GetEditLessonRHN,
		editor.PostEditLessonRHN,
		editor.NameContainerID,
		editor.NameInputID,
		editor.E,
		isEdit,
		input,
		editor.Lesson.Name,
	)
}

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
	NewTab     bool
	HxSwap     HxSwap
}

type HxSwap string

const (
	AfterEnd = "afterend"
)

func (button HXButton) Component() templ.Component {
	return hxButton(button)
}
func NewHXButton(method HXMethod, hxSwap HxSwap, url, hxTargetID string, pushURL, newTab bool) HXButton {
	return HXButton{
		Method:     method,
		URL:        url,
		HxTargetID: hxTargetID,
		PushURL:    pushURL,
		NewTab:     newTab,
		HxSwap:     hxSwap,
	}
}

type HXMethod string

const (
	HxGet  = "hx-get"
	HxPost = "hx-post"
)

package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type LessonDetailsPage struct {
	NodeDetailsPage
	Slides                                                       string
	E                                                            *echo.Echo
	Standards                                                    []domain.Standard
	GetObjectivesURL                                             string
	FileRHN                                                      string
	ViewMarkdownRHN                                              string
	PostLessonStandardURL, DeleteLessonStandardRHN               string
	GetEditAssessmentRHN, PostAssessmentURL, DeleteAssessmentRHN string
	GetSlidesURL, EditSlidesURL                                  string
}

func (page LessonDetailsPage) DeleteStandardURL(stdID int) string {
	return page.E.Reverse(page.DeleteLessonStandardRHN, AddParams(page.Params, stdID)...)
}

func (page LessonDetailsPage) DeleteAssessmentURL(assessmentID int) string {
	return page.E.Reverse(page.DeleteAssessmentRHN, AddParams(page.Params, assessmentID)...)
}

func (page LessonDetailsPage) Lesson() domain.Lesson {
	return page.Node.(domain.Lesson)
}

func (page LessonDetailsPage) EditSlidesButton() templ.Component {
	return cmp.Button{
		Text:     "Edit Slides",
		Method:   cmp.HxGet,
		URL:      page.EditSlidesURL,
		HxTarget: EditSlidesContainerID.Selector(),
		PushURL:  true,
	}.Component()

}

func (data LessonDetailsPage) ViewMarkdownButton(filepath string) templ.Component {
	return cmp.Button{
		Text:     "View As HTML",
		Method:   cmp.HxGet,
		HxTarget: "#markdown",
		URL:      data.FileURL(filepath),
		PushURL:  true,
	}.Component()

}

func (data LessonDetailsPage) FileURL(filepath string) string {
	return data.E.Reverse(data.ViewMarkdownRHN, AddParams(data.Params, filepath)...)

}

func (page LessonDetailsPage) Component() templ.Component {
	return LessonDetailsComponent(page)
}

type ObjectiveSelect struct {
	Objectives []domain.Standard
}

// capitalizes first letter
func DisplayCategory(cat domain.AssessmentCategory) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(cat.String()[:1]), cat.String()[1:])
}

type EditAssessmentForm struct {
	Params                domain.NodePath
	Assessment            domain.Assessment
	PostEditAssessmentURL string
	LessonDetailsURL      string
}

func (data EditAssessmentForm) Component() templ.Component {
	return EditAssessmentFormComponent(data)
}

func (page LessonDetailsPage) NewAssessmentForm() templ.Component {
	form := cmp.NewForm(cmp.NewFormParams{
		Title:     "New Assessment",
		Subtitle:  "Placeholder",
		PostURL:   page.PostAssessmentURL,
		CancelURL: page.CancelEditURL,
		HxTarget:  "#page",
	})
	nameInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Name",
		Type: cmp.Text,
	})
	var date string
	if page.Lesson().Dates != nil {
		date = page.Lesson().Dates[0].Format(time.DateOnly)
	}
	assignedInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "Date Assigned",
		Type:  cmp.Date,
		Value: date,
	})
	dueInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "Date Due",
		Type:  cmp.Date,
		Value: date,
	})
	var options []cmp.Option
	for _, category := range domain.Categories {
		catOption := cmp.Option{
			Value:   strconv.Itoa(int(category)),
			Content: category.String(),
		}
		options = append(options, catOption)
	}
	categorySelect := cmp.NewSelectWithLabel("Category", options)
	fileInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "File",
		Type: cmp.Text,
	})
	instructionsTextArea := cmp.NewTextAreaWithLabel(cmp.TextAreaWithLabelParams{
		Name: "Instructions",
	})
	hiddenIdInput := cmp.NewHiddenInput(cmp.HiddenInputParams{
		Name:  "Lesson ID",
		Value: strconv.Itoa(page.Lesson().GetID().(int)),
	})
	form = form.AddElement(nameInput, assignedInput, dueInput, categorySelect, fileInput, instructionsTextArea, hiddenIdInput)
	return form.Component()

}

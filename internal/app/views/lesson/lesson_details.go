package lessonviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/core/assessment"
	"gh_static_portfolio/internal/core/standard"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"path/filepath"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

type LessonDetailsPage struct {
	ac.NodeDetailsPage
	AssetsURLFunc               func(path ...string) string
	GetObjectivesURL            string
	FileURL                     web.AddParams
	ViewMarkdownURL             web.AddParams
	GetStandardsURL             string
	GetAssessmentsURL           string
	GetSlidesURL, EditSlidesURL string
}

type LessonDetailsEdit struct {
	LessonDetailsPage
}

func (page LessonDetailsPage) EditDetails() LessonDetailsEdit {
	editPage := LessonDetailsEdit{LessonDetailsPage: page}
	editPage.IsEdit = true
	return editPage
}

func (page LessonDetailsEdit) HTMXResponse() templ.Component {
	return page.NodeDetailsPage.DetailsEdit().HTMXResponse()
}

func (page LessonDetailsEdit) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.WithPage(page.Component())
}

func (page LessonDetailsPage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page LessonDetailsPage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.WithPage(page.Component())
}

func (page LessonDetailsPage) Lesson() dto.Lesson {
	return page.Node.(dto.Lesson)
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

func (page LessonDetailsPage) Component() templ.Component {
	return LessonDetailsComponent(page)
}

func (page LessonDetailsPage) Tabs() templ.Component {
	tabs := []cmp.TabLink{
		{Name: "Slides", URL: page.GetSlidesURL},
		// {Name: "Assessments", URL: page.GetAssessmentsURL},
		// {Name: "Standards", URL: page.GetStandardsURL},
	}
	set := cmp.TabSet{
		Tabs: tabs,
		AssetsURLFunc: func(s string) string {
			return page.AssetsURLFunc(filepath.Join("js", s))
		},
	}.Component()
	return set
}

type ObjectiveSelect struct {
	Objectives []standard.Standard
}

type EditAssessmentForm struct {
	Params                routes.NodePath
	Assessment            assessment.Assessment
	PostEditAssessmentURL string
	LessonDetailsURL      string
}

// for editing existing assessment, to replace old assessment form component
func (data EditAssessmentForm) NewEditAssessmentFormComponent() templ.Component {
	form := cmp.NewForm(cmp.NewFormParams{
		Title:     "Edit Assessment",
		PostURL:   data.PostEditAssessmentURL,
		CancelURL: data.LessonDetailsURL,
		HxTarget:  "#page",
	})
	nameInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "Name",
		Type:  cmp.Text,
		Value: data.Assessment.Name,
	})
	assignedInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "Date Assigned",
		Type:  cmp.Date,
		Value: data.Assessment.DateAssigned.Format(time.DateOnly),
	})
	dueInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "Date Due",
		Type:  cmp.Date,
		Value: data.Assessment.DateDue.Format(time.DateOnly),
	})
	var options []cmp.Option
	for _, category := range assessment.Categories {
		catOption := cmp.Option{
			Value:    string(category),
			Content:  string(category),
			Selected: category == data.Assessment.Category,
		}
		options = append(options, catOption)
	}
	categorySelect := cmp.NewSelectWithLabel("Category", options)
	fileInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "File",
		Type:  cmp.Text,
		Value: data.Assessment.File,
	})
	instructionsTextArea := cmp.NewTextAreaWithLabel(cmp.TextAreaWithLabelParams{
		Name:  "Instructions",
		Value: data.Assessment.Instructions,
	})
	hiddenIdInput := cmp.NewHiddenInput(cmp.HiddenInputParams{
		Name:  "Lesson ID",
		Value: strconv.Itoa(data.Params.LessonID),
	})
	form = form.AddElement(nameInput, assignedInput, dueInput, categorySelect, fileInput, instructionsTextArea, hiddenIdInput)
	return form.Component()

}

func (data EditAssessmentForm) Component() templ.Component {
	return data.NewEditAssessmentFormComponent()
}

const (
	EditSlidesContainerID string = "slides-editor-container"
	EditSlidesTextAreaID  string = "slides-editor-text-area"
)

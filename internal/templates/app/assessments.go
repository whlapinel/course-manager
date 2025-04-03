package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"
	"gh_static_portfolio/internal/util"
	"net/url"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

type CourseAssessmentsPage struct {
	GetAssessmentsURL string
	DateFilterURL     string
	Assessments       []domain.Assessment
	CourseListURL     string
	BreadCrumbsData   BreadCrumbs
	StartDate         time.Time
	EndDate           time.Time
	Category          domain.AssessmentCategory
}

func (page CourseAssessmentsPage) ApplyFilterButton() templ.Component {

	applyButton := cmp.NewButton(cmp.NewButtonParams{
		Button: cmp.Button{
			Text: "Apply",
			Element: cmp.Element{
				ID: "filter-button",
			},
			Method:   cmp.HxGet,
			URL:      page.GetAssessmentsURL,
			HxTarget: "#assessments",
			Attributes: templ.Attributes{
				"hx-select": "#assessments",
			},
			PushURL: true,
		},
	})
	return applyButton.Component()
}

func (page CourseAssessmentsPage) CategorySelectComponent() templ.Component {
	var options []cmp.Option
	options = append(options, cmp.Option{
		Content:  "Select a Category",
		Value:    "",
		Selected: page.Category == "",
	})
	for _, category := range domain.Categories {
		option := cmp.Option{
			Content:  util.Capitalize(string(category)),
			Value:    string(category),
			Selected: page.Category == category,
		}
		options = append(options, option)
	}
	catSelect := cmp.NewSelectWithLabel("Category", options)
	return catSelect.Component()
}

func (page CourseAssessmentsPage) StartDateComponent(date time.Time) templ.Component {
	var dateString string
	if !date.IsZero() {
		dateString = date.Format(time.DateOnly)
	}
	start := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  "Start",
		Type:  cmp.Date,
		Value: dateString,
	})
	return start.Component()
}

func (page CourseAssessmentsPage) DateComponent(name string, date time.Time) templ.Component {
	var dateString string
	if !date.IsZero() {
		dateString = date.Format(time.DateOnly)
	}
	return cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name:  name,
		Type:  cmp.Date,
		Value: dateString,
	}).Component()
}
func (page CourseAssessmentsPage) EndDateComponent() templ.Component {
	end := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "End",
		Type: cmp.Date,
	})
	return end.Component()
}

func (page CourseAssessmentsPage) GetAssessmentsWithCategoryFilter(category domain.AssessmentCategory) string {
	parsedURL, err := url.Parse(page.GetAssessmentsURL)
	if err != nil {
		panic("error parsing URL")
	}
	q := parsedURL.Query()
	q.Set("category", string(category))

	parsedURL.RawQuery = q.Encode()

	return parsedURL.String()
}

func (page CourseAssessmentsPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Assessments",
		UpNav: cmp.UpNav{
			URL:  page.CourseListURL,
			Text: "Up to courses",
		},
		Crumbs: page.BreadCrumbs().BreadCrumbs(),
	}

}

func (page CourseAssessmentsPage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData

}

func (page CourseAssessmentsPage) Component() templ.Component {
	return CourseAssessmentsPageComponent(page)
}

type AssessmentsFragment struct {
	NewAssessmentURL     string
	CancelEditURL        string
	GetEditAssessmentURL func(id any) string
	Assessments          []domain.Assessment
	NodeID               int
	ViewFileURL          func(relPath string) string
}

type NewAssessmentForm struct {
	CancelURL         string
	PostAssessmentURL string
	NodeID            int
}

func (data AssessmentsFragment) Component() templ.Component {
	return AssessmentsFragmentComponent(data)
}

func (data AssessmentsFragment) Info(assm domain.Assessment) cmp.EditableInfo {
	info := cmp.EditableInfo{
		Element: cmp.Element{
			ID: fmt.Sprintf("%s-%d", "assessment", assm.ID),
		},
		Title:      assm.Name,
		GetEditURL: data.GetEditAssessmentURL(assm.ID),
		Components: []cmp.EditableInfoItem{
			{Field: "Name", Value: assm.Name},
			{Field: "Instructions", Value: assm.Instructions},
			{Field: "File", Value: assm.File},
			{Field: "Assigned", Value: assm.DateAssigned.Format(time.DateOnly)},
			{Field: "Due", Value: assm.DateDue.Format(time.DateOnly)},
			{Field: "Category", Value: string(assm.Category)},
			{Field: "Dropped", Value: strconv.FormatBool(assm.Dropped)},
		},
	}
	return info
}

func (page AssessmentsFragment) ViewFileLink(filename string) templ.Component {
	return cmp.Link{
		Text: "View File",
		Attributes: templ.Attributes{
			string(cmp.HxGet): page.ViewFileURL(filename),
			"hx-target":       "#file-content",
		},
	}.Component()
}

func (page AssessmentsFragment) AddAssessmentButton() templ.Component {
	return AddAssessmentButtonComponent(page)
}

func (data NewAssessmentForm) Component() templ.Component {
	form := cmp.NewForm(cmp.NewFormParams{
		Title:     "New Assessment",
		Subtitle:  "Placeholder",
		PostURL:   data.PostAssessmentURL,
		CancelURL: data.CancelURL,
		HxTarget:  "#page",
	})
	nameInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Name",
		Type: cmp.Text,
	})
	assignedInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Date Assigned",
		Type: cmp.Date,
	})
	dueInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Date Due",
		Type: cmp.Date,
	})
	var options []cmp.Option
	for _, category := range domain.Categories {
		catOption := cmp.Option{
			Value:   string(category),
			Content: string(category),
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
		Value: strconv.Itoa(data.NodeID),
	})
	form = form.AddElement(nameInput, assignedInput, dueInput, categorySelect, fileInput, instructionsTextArea, hiddenIdInput)
	return form.Component()

}

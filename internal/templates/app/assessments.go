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
}

func (page CourseAssessmentsPage) ApplyFilterButton() templ.Component {

	applyButton := cmp.NewButton(cmp.NewButtonParams{
		Button: cmp.Button{
			Text: "Apply",
			Element: cmp.Element{
				ID: "apply-filter-button",
			},
			Method:   cmp.HxGet,
			URL:      page.GetAssessmentsURL,
			HxSelect: "#assessments",
			HxTarget: "#assessments",
			PushURL:  true,
		},
	})
	return applyButton.Component()
}

func (page CourseAssessmentsPage) CategorySelectComponent() templ.Component {
	var options []cmp.Option
	for _, category := range domain.Categories {
		option := cmp.Option{
			Content: util.Capitalize(category.String()),
			Value:   category.String(),
		}
		options = append(options, option)
	}
	catSelect := cmp.NewSelectWithLabel("Category", options)
	return catSelect.Component()
}

func (page CourseAssessmentsPage) StartDateComponent() templ.Component {
	start := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Start",
		Type: cmp.Date,
	})
	return start.Component()
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
	catString := strconv.Itoa(int(category))
	q.Set("category", catString)

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
}

type NewAssessmentForm struct {
	CancelURL         string
	PostAssessmentURL string
	NodeID            int
}

func (data AssessmentsFragment) Component() templ.Component {
	return AssessmentsFragmentComponent(data)
}

func (data AssessmentsFragment) Infos() []cmp.EditableInfo {
	var infos []cmp.EditableInfo
	for _, assm := range data.Assessments {
		info := cmp.EditableInfo{
			Element: cmp.Element{
				ID: fmt.Sprintf("%s-%d", "assessment", assm.ID),
			},
			Title:      assm.Name,
			GetEditURL: data.GetEditAssessmentURL(assm.ID),
			Components: []cmp.EditableInfoItem{
				{Field: "Name", Value: assm.Name},
				{Field: "Instructions", Value: assm.Instructions},
				{Field: "Assigned", Value: assm.DateAssigned.Format(time.DateOnly)},
				{Field: "Due", Value: assm.DateDue.Format(time.DateOnly)},
				{Field: "Category", Value: assm.Category.String()},
				{Field: "Dropped", Value: strconv.FormatBool(assm.Dropped)},
			},
		}
		infos = append(infos, info)
	}
	return infos
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
		Value: strconv.Itoa(data.NodeID),
	})
	form = form.AddElement(nameInput, assignedInput, dueInput, categorySelect, fileInput, instructionsTextArea, hiddenIdInput)
	return form.Component()

}

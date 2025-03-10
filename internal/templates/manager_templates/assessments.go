package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components"
	"gh_static_portfolio/internal/util"
	"net/url"
	"strconv"

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

	applyButton := cmp.NewHXButton(cmp.NewButtonParams{
		HXButton: cmp.HXButton{
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

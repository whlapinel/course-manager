package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
)

type CourseAssessmentsPage struct {
	GetAssessmentsURL string
	DateFilterURL     string
	Assessments       []domain.Assessment
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

func (page CourseAssessmentsPage) Component() templ.Component {
	return CourseAssessmentsPageComponent(page)
}

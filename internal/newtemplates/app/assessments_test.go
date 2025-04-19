package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"log"
	"testing"
)

func TestGetAssessmentsWithCategoryFilter(t *testing.T) {
	cat := domain.Prepare
	page := CourseAssessmentsPage{
		GetAssessmentsURL: "/terms/2/courses/5/assessments",
	}
	url := page.GetAssessmentsWithCategoryFilter(cat)
	log.Println(url)

}

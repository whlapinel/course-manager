package slides

import (
	"fmt"
	"gh_static_portfolio/internal/shared/routes"
	"net/url"
)

type Service struct {
	MarpBaseURL string
}

func NewSlidesService(marpBaseURL string) *Service {
	return &Service{
		MarpBaseURL: marpBaseURL,
	}

}

func (s *Service) marpSlidesPath(params routes.NodePath) (string, error) {
	// marpHost := os.Getenv("MARP_HOST")
	// marpPort := os.Getenv("MARP_PORT")
	// baseURL := fmt.Sprintf("http://%s:%s", marpHost, marpPort)
	userParam := fmt.Sprintf("user_%s", params.UserID)
	termParam := fmt.Sprintf("term_%d", params.TermID)
	courseParam := fmt.Sprintf("course_%d", params.CourseID)
	unitParam := fmt.Sprintf("unit_%d", params.UnitID)
	lessonParam := fmt.Sprintf("lesson_%d", params.LessonID)
	return url.JoinPath(s.MarpBaseURL, "users", userParam, "terms", termParam, "courses", courseParam, "units", unitParam, "lessons", lessonParam, "slides.md")
}

package slides

import (
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/node"
	"net/url"
)

type Service struct {
	marpBaseURL string
	paths       ports.PathingService
}

func NewSlidesService(
	marpBaseURL string,
	paths ports.PathingService,
) *Service {
	return &Service{
		marpBaseURL: marpBaseURL,
		paths:       paths,
	}
}

func (s *Service) marpSlidesPath(nodes ...node.Node) (string, error) {
	relPath := s.paths.NodeDirPath(nodes...)
	fullPath, err := url.JoinPath(s.marpBaseURL, relPath, "slides.md")
	if err != nil {
		return "", err
	}
	return fullPath, nil
}

package services

import "gh_static_portfolio/internal/ports"

type SiteGeneratorService struct {
	ports.SiteGenerator
}

func NewSiteGeneratorService(generator ports.SiteGenerator) *SiteGeneratorService {
	return &SiteGeneratorService{
		SiteGenerator: generator,
	}
}

func (s *SiteGeneratorService) Build(user, term, course ports.Node) error {
	return s.SiteGenerator.Build(user, term, course)
}

package sitegen

import (
	"gh_static_portfolio/internal/ports"
)

type Service struct {
	ports.SiteGenerator
}

func New(generator ports.SiteGenerator) *Service {
	return &Service{
		SiteGenerator: generator,
	}
}

func (s *Service) Build() error {
	return s.SiteGenerator.Build()
}

package hugo

import (
	"errors"
	"gh_static_portfolio/internal/ports"
)

type hugoGenerator struct {
}

func New() ports.SiteGenerator {
	return &hugoGenerator{}
}

func (h *hugoGenerator) Build() error {
	return errors.New("not implemented")
}

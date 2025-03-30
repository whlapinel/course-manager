package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
)

type StandardsFragment struct {
	NodeStandards     []domain.Standard
	CourseStandards   []domain.Standard
	PostStandardURL   string
	DeleteStandardURL func(id any) string
}

func (data StandardsFragment) Component() templ.Component {
	return StandardsFragmentComponent(data)
}

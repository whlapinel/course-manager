package appcomponents

import (
	components "gh_static_portfolio/internal/newtemplates/components/base"

	"github.com/a-h/templ"
)

type Page interface {
	ComponentData
	PageLayout() components.PageLayout
}

type ComponentData interface {
	Component() templ.Component
}

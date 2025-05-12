package appcomponents

import (
	components "gh_static_portfolio/internal/base"

	"github.com/a-h/templ"
)

type Page interface {
	ComponentData
	PageLayout() components.PageLayout
}

type ComponentData interface {
	Component() templ.Component
}

// later interfaces written below here and represent latest effort to provide useful abstractions

type Layout interface {
	Component2(templ.Component) templ.Component
}

// terrible name but Component2 is the layout

type BasicHTMXPage[T HTMXPage] struct {
	Page T
}

type HTMXPage interface {
	ComponentData
	Layout
}

func (p BasicHTMXPage[HTMXPage]) HTMXResponse() templ.Component {
	return p.Page.Component()
}
func (p BasicHTMXPage[HTMXPage]) NonHTMXResponse() templ.Component {
	return p.Page.Component()
}

package authviews

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	components "gh_static_portfolio/internal/basecomponents"

	"github.com/a-h/templ"
)

type UnauthorizedPage struct {
	Message string
	Link    components.Link
	appcomponents.CourseManagerLayout
}

func (page UnauthorizedPage) Component() templ.Component {
	return UnauthorizedPageComponent(page)
}

func (p UnauthorizedPage) HTMXResponse() templ.Component {
	return p.Component()
}

func (p UnauthorizedPage) NonHTMXResponse() templ.Component {
	return p.WithPage(p.Component())
}

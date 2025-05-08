package authviews

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	components "gh_static_portfolio/internal/newtemplates/components/base"

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
	return p.Component2(p.Component())
}

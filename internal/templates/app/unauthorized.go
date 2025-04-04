package managertemplates

import (
	components "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type UnauthorizedPage struct {
	Message string
	Link    components.Link
}

func (page UnauthorizedPage) Component() templ.Component {
	return UnauthorizedPageComponent(page)
}

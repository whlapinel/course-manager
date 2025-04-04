package templates

import (
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type Layout struct {
	templ.Attributes
	PageTitle string
	Head      templ.Component
	Teacher   domain.User
	NavItems  []components.NavItem
	Page      templ.Component
}

func (data Layout) Component() templ.Component {
	return LayoutComponent(data)

}

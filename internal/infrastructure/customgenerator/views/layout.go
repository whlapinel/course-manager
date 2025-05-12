package templates

import (
	components "gh_static_portfolio/internal/base"
	"gh_static_portfolio/internal/domain"

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

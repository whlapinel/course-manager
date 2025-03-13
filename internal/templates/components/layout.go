package components

import (
	"fmt"

	"github.com/a-h/templ"
)

type Layout struct {
	PageTitle string
	Head      templ.Component
	NavItems  []templ.Component
	Page      templ.Component
}

func (data Layout) Component() templ.Component {
	return LayoutComponent(data)
}

type NavItem struct {
	Method    HXMethod
	Text      string
	URL       string
	HxConfirm string
	Class     string
}

func (data NavItem) MethodAttributes() templ.Attributes {
	var attr templ.Attributes = make(templ.Attributes)
	if data.Method == "" {
		attr["href"] = fmt.Sprintf(data.URL)
	} else {
		attr[string(data.Method)] = data.URL
	}
	if data.HxConfirm != "" {
		attr["hx-confirm"] = data.HxConfirm
	}
	return attr
}

func (data NavItem) Component() templ.Component {
	return NavItemComponent(data)
}

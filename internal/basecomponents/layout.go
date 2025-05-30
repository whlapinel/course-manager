package basecomponents

import (
	"github.com/a-h/templ"
)

type Layout struct {
	HomeURL          string // clicking logo will send user to home e.g. index.html or "/"
	PageTitle        string
	Head             templ.Component
	UserMenu         UserMenu
	NavItems         []NavItem
	Page             templ.Component
	UserImage        string
	SignoutURL       string
	AssetsURLFunc    func(...string) string
	templ.Attributes // will be added to body element
}

func (data Layout) Component() templ.Component {
	return LayoutComponent(data)
}

type NavItem struct {
	Element
	Method    HXMethod
	Text      string
	URL       string
	HxConfirm string
	Class     string
	Current   bool
}

func (data NavItem) MethodAttributes() templ.Attributes {
	var attr templ.Attributes = make(templ.Attributes)
	if data.Method == "" {
		attr["href"] = data.URL
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

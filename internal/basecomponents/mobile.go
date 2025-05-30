package basecomponents

import (
	"strconv"

	"github.com/a-h/templ"
)

type MobileMenu struct {
	Layout
}

type MobileMenuButton struct {
	UserMenu
}

func (data MobileMenu) LinksComponent() templ.Component {
	var linkComps []templ.Component
	for i, link := range data.UserMenu.Links {
		link.Class = "block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white cursor-pointer"
		attributes := templ.Attributes{
			"role":     "menuitem",
			"tabindex": "-1",
			"id":       "mobile-menu-item-" + strconv.Itoa(i),
		}
		// make sure user can override attributes by only setting those which weren't already set
		for key, val := range attributes {
			if _, exists := link.Attributes[key]; !exists {
				link.Attributes[key] = val
			}
		}
		linkComps = append(linkComps, link.Component())
	}
	return templ.Join(linkComps...)
}

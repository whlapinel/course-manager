package components

import (
	"strconv"

	"github.com/a-h/templ"
)

type UserMenu struct {
	Image string
	Name  string
	Links []Link
}

func (data UserMenu) Component() templ.Component {
	return UserMenuComponent(data)
}

func (data UserMenu) LinksComponent() templ.Component {
	var linkComps []templ.Component
	for i, link := range data.Links {
		link.Class = "block px-4 py-2 text-sm text-gray-700 cursor-pointer"
		attributes := templ.Attributes{
			"role":     "menuitem",
			"tabindex": "-1",
			"id":       "user-menu-item-" + strconv.Itoa(i),
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

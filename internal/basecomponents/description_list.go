package basecomponents

import (
	"github.com/a-h/templ"
)

type DescriptionList struct {
	Title, Subtitle string
	Items           []DescriptionListItem
}

type DescriptionListItem struct {
	Element
	Name  string
	Value string
}

func (el DescriptionList) Component() templ.Component {
	return DescriptionListComponent(el)
}

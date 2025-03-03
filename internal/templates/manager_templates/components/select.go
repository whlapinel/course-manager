package components

import (
	"github.com/a-h/templ"
)

type SelectWithLabel struct {
	Label
	Select
}

func (el SelectWithLabel) Component() templ.Component {
	return SelectWithLabelComponent(el)
}

func NewSelectWithLabel(name string, options []Option) SelectWithLabel {
	id := Kebab(name)
	return SelectWithLabel{
		Label: Label{
			Content: name,
			For:     id,
		},
		Select: Select{
			Element: Element{
				ID: id,
			},
			Name:    name,
			Options: options,
		},
	}
}

type Select struct {
	Element
	Name    string
	Options []Option
}

func (el Select) Component() templ.Component {
	return SelectComponent(el)
}

type Option struct {
	Element
	Content string
	Value   string
}

func (el Option) Component() templ.Component {
	return OptionComponent(el)
}

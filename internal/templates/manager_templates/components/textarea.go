package components

import "github.com/a-h/templ"

type TextArea struct {
	Element
	Name  string
	Value string
}

type TextAreaWithLabel struct {
	TextArea
	Label
}

func (el TextAreaWithLabel) Component() templ.Component {
	return TextAreaWithLabelComponent(el)
}

type TextAreaWithLabelParams struct {
	Name  string
	Value string
}

func NewTextAreaWithLabel(params TextAreaWithLabelParams) TextAreaWithLabel {
	id := Kebab(params.Name)
	name := Kebab(params.Name)
	return TextAreaWithLabel{
		TextArea{
			Element: Element{ID: id},
			Name:    name,
		},
		Label{
			Content: params.Name,
			For:     id,
		},
	}
}

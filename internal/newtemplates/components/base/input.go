package components

import (
	"strings"

	"github.com/a-h/templ"
)

type Input struct {
	Element
	InputType    InputType
	Name         string
	Value        string
	Placeholder  string
	Autocomplete string
	Editing      bool
}

type InputType string

func (it InputType) String() string {
	return string(it)
}

const (
	Text   InputType = "text"
	Date   InputType = "date"
	Number InputType = "number"
	File   InputType = "file"
	Email  InputType = "email"
	Hidden InputType = "hidden"
)

func (el Input) Component() templ.Component {
	return InputComponent(el)
}

type InputWithLabelParams struct {
	Name  string
	Type  InputType
	Value string
}

func Kebab(name string) string {
	id := strings.ReplaceAll(name, " ", "-")
	id = strings.ToLower(id)
	return id
}
func NewInputWithLabel(params InputWithLabelParams) InputWithLabel {
	id := Kebab(params.Name)
	return InputWithLabel{
		Input: Input{
			Element:   Element{ID: id},
			InputType: params.Type,
			Name:      Kebab(params.Name),
			Value:     params.Value,
		},
		Label: Label{
			Content: params.Name,
			For:     id,
		},
	}
}

type HiddenInputParams struct {
	Name  string
	Value string
}

func NewHiddenInput(params HiddenInputParams) Input {
	id := Kebab(params.Name)
	return Input{
		Element:   Element{ID: id},
		Name:      Kebab(params.Name),
		Value:     params.Value,
		InputType: Hidden,
	}

}

type InputWithLabel struct {
	Input
	Label
}

func (el InputWithLabel) Component() templ.Component {
	return InputWithLabelComponent(el)
}

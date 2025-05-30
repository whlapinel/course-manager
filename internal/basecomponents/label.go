package basecomponents

import "github.com/a-h/templ"

type Label struct {
	Element
	Content string // text content
	For     string // id of input element
}

func (el Label) Component() templ.Component {
	return LabelComponent(el)

}

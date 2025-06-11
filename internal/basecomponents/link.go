package basecomponents

import "github.com/a-h/templ"

type Link struct {
	Element
	Text   string
	URL    string
	Target LinkTarget
	Class  string
	HTMX   bool
	templ.Attributes
}

type LinkTarget string

const (
	NewTab LinkTarget = "_blank"
)

func (data LinkTarget) String() string {
	return string(data)
}

func (data Link) Component() templ.Component {
	return LinkComponent(data)
}

package components

import "github.com/a-h/templ"

type Link struct {
	Text string
	URL  string
}

func (data Link) Component() templ.Component {
	return LinkComponent(data)
}

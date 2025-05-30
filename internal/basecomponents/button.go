package basecomponents

import "github.com/a-h/templ"

type Button struct {
	Element
	templ.Attributes
	Name      string
	Text      string
	Image     templ.Component
	HxConfirm string
	Method    HXMethod
	URL       string
	HxTarget  string
	PushURL   bool
	HxSwap    HxSwap
	HxSelect  string
	Class     string
}

type HxSwap string

const (
	AfterEnd HxSwap = "afterend"
)

func (button Button) Component() templ.Component {
	return HxButtonComponent(button)
}

type NewButtonParams struct {
	Button
}

func NewButton(params NewButtonParams) Button {
	return params.Button
}

type HXMethod string

const (
	HxGet    HXMethod = "hx-get"
	HxPost   HXMethod = "hx-post"
	HxDelete HXMethod = "hx-delete"
)

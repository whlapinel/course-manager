package components

import "github.com/a-h/templ"

type HXButton struct {
	Element
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
}

type HxSwap string

const (
	AfterEnd HxSwap = "afterend"
)

func (button HXButton) Component() templ.Component {
	return HxButtonComponent(button)
}

type NewButtonParams struct {
	HXButton
}

func NewHXButton(params NewButtonParams) HXButton {
	return params.HXButton
}

type HXMethod string

const (
	HxGet    HXMethod = "hx-get"
	HxPost   HXMethod = "hx-post"
	HxDelete HXMethod = "hx-delete"
)

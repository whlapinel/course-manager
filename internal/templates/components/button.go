package components

import "github.com/a-h/templ"


type HXButton struct {
	Text      string
	Image     templ.Component
	HxConfirm string
	Method    HXMethod
	URL       string
	HxTarget  string
	PushURL   bool
	HxSwap    HxSwap
}

type HxSwap string

const (
	AfterEnd HxSwap = "afterend"
)

func (button HXButton) Component() templ.Component {
	return HxButtonComponent(button)
}
func NewHXButton(method HXMethod, hxSwap HxSwap, url, hxTargetID string, pushURL bool) HXButton {
	return HXButton{
		Method:   method,
		URL:      url,
		HxTarget: hxTargetID,
		PushURL:  pushURL,
		HxSwap:   hxSwap,
	}
}

type HXMethod string

const (
	HxGet    HXMethod = "hx-get"
	HxPost   HXMethod = "hx-post"
	HxDelete HXMethod = "hx-delete"
)

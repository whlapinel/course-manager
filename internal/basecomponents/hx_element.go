package basecomponents

import "github.com/a-h/templ"

type HxElement interface {
	HXMethod() HXMethod
	URL() string
}

func Attributes(el HxElement) templ.Attributes {
	var attr = make(templ.Attributes)
	attr[string(el.HXMethod())] = el.URL()
	return attr
}

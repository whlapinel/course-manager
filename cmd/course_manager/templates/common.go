package templates

import "github.com/a-h/templ"

type Page struct {
	Title     string
	Route     string
	Component templ.Component
}

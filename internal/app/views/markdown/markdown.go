package markdownviews

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/basecomponents"

	"github.com/a-h/templ"
)

type MarkdownDocument struct {
	Title   string
	Content string
	Static  bool
}

type MarkdownEditor struct {
	New       bool
	Name      string
	Contents  string
	SubmitURL string
	CancelURL string
	appcomponents.CourseManagerLayout
}

func (data MarkdownEditor) Component() templ.Component {
	return MarkdownEditorComponent(data)
}

func (data MarkdownEditor) HTMXResponse() templ.Component {
	return data.Component()
}

func (data MarkdownEditor) NonHTMXResponse() templ.Component {
	return data.WithPage(data.Component())
}

func (data MarkdownEditor) CancelButton() templ.Component {
	return basecomponents.Button{
		Text: "Cancel",
		Attributes: templ.Attributes{
			"hx-get":    data.CancelURL,
			"hx-target": "#page",
			"hx-push-url": "true",
		},
	}.Component()
}

func (data MarkdownEditor) SubmitButton() templ.Component {
	return basecomponents.Button{
		Text: "Submit",
		Attributes: templ.Attributes{
			"hx-post":    data.SubmitURL,
			"hx-target":  "#markdown",
			"hx-include": ".markdown-form",
		},
	}.Component()
}

const (
	EditSlidesContainerID string = "slides-editor-container"
	EditSlidesTextAreaID  string = "slides-editor-text-area"
)

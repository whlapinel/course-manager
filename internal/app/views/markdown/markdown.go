package markdownviews

import (
	appcomponents "gh_static_portfolio/internal/app/components"

	"github.com/a-h/templ"
)

type MarkdownDocument struct {
	Title   string
	Content string
	Static  bool
}

type MarkdownEditor struct {
	Name            string
	Contents        string
	PostEditFileURL string
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

const (
	EditSlidesContainerID string = "slides-editor-container"
	EditSlidesTextAreaID  string = "slides-editor-text-area"
)

package lessonviews

import (
	cmp "gh_static_portfolio/internal/basecomponents"

	"github.com/a-h/templ"
)

type Slides struct {
	HTML          string
	EditSlidesURL string
	LessonDetailsPage
}

func (data Slides) Component() templ.Component {
	return SlidesComponent(data)
}

func (data Slides) EditSlidesButton() templ.Component {
	return cmp.Button{
		Text:     "Edit Slides",
		Method:   cmp.HxGet,
		URL:      data.EditSlidesURL,
		HxTarget: "#slides",
		PushURL:  true,
	}.Component()

}

type SlidesEditor struct {
	Content           string
	PostEditSlidesURL string
}

func (e SlidesEditor)Component()templ.Component{
	return EditSlidesTemplate(e)
}
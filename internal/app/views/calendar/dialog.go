package calendarviews

import (
	"time"

	"github.com/a-h/templ"
)

type AddOccasionDialog struct {
	PostURL string
	Date    time.Time
}

func (data AddOccasionDialog) Component() templ.Component {
	return DialogComponent(data)
}

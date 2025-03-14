package templates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
)

type Head struct {
	User      domain.User
	AssetsURL func(relpath string) string
}

func (data Head) Component() templ.Component {
	return HeadComponent(data)
}

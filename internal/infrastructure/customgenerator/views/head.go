package templates

import (
	"gh_static_portfolio/internal/app/dto"

	"github.com/a-h/templ"
)

type Head struct {
	User      dto.User
	AssetsURL func(relpath string) string
}

func (data Head) Component() templ.Component {
	return HeadComponent(data)
}

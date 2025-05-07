package handlers

import (
	"context"
	"gh_static_portfolio/internal/shared/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// the idea behind this is, for most pages, a page should
// be given all data for a full page refresh as well as partial
// (HTMX requests) updates. This way most htmx links can support
// pushing the url to the history (if a GET route is pushed to the history
// then that route's GET handler should also support a full page refresh)
type Page interface {
	HTMXResponse() templ.Component
	NonHTMXResponse() templ.Component
}

func Respond(c echo.Context, page Page) error {
	if web.IsHTMX(c) {
		return page.HTMXResponse().Render(context.Background(), c.Response())
	}
	return page.NonHTMXResponse().Render(context.Background(), c.Response())
}

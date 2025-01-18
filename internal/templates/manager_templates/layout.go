package managertemplates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseManagerLayout struct {
	PageTitle                                  string
	Page                                       templ.Component
	ListTermsRHN, GenerateSiteRHN, SyncSiteRHN string
	E                                          *echo.Echo
}

package assets

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed dist
var Assets embed.FS

func RegisterStatic(e *echo.Echo) {
	e.GET("/dist/*", echo.WrapHandler(http.FileServerFS(Assets)))
}

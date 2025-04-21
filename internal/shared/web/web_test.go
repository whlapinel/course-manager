package web

import (
	"log"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

const testRouteName = "test_route"

func TestMain(m *testing.M) {
	e = echo.New()
	e.GET("/test/:id/:second-id", func(c echo.Context) error {
		log.Println("cool!")
		return nil
	}).Name = testRouteName
	os.Exit(m.Run())
}

func TestURLFunc(t *testing.T) {
	fn := URLFunc(testRouteName, e.Reverse, 1)
	url := fn(2)
	if url != "/test/1/2" {
		t.Error("url:", url)
	}
}

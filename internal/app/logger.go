package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var logger = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogURI:       true,
	LogStatus:    true,
	LogRoutePath: true,
	LogMethod:    true,
	LogLatency:   true,
	LogError:     true,
	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		const (
			reset  = "\033[0m"
			red    = "\033[31m"
			green  = "\033[32m"
			yellow = "\033[33m"
		)
		statusColor := reset
		if v.Status >= 400 {
			statusColor = red
		} else if v.Status >= 300 {
			statusColor = yellow
		} else {
			statusColor = green
		}
		methodWidth := 6
		uriWidth := 25
		statusWidth := 6
		customWidth := 12
		latencyWidth := 15
		errorWidth := 20
		pathWidth := 50
		value, _ := c.Get("id").(int)
		if v.Error == nil {
			v.Error = fmt.Errorf("no error")
		}
		logLine := fmt.Sprintf("*******\n%-*s\n%-*s\n%-*s\n%s%-*d%s %-*d %-*s %-*s\n*********",
			methodWidth, v.Method,
			uriWidth, v.URI,
			pathWidth, v.RoutePath,
			statusColor,
			statusWidth, v.Status,
			reset,
			customWidth, value,
			latencyWidth, v.Latency,
			errorWidth, v.Error,
		)
		fmt.Println(logLine)
		return nil
	},
})

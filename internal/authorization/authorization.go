package authorization

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"log"

	"github.com/labstack/echo/v4"
)

// middleware to verify user is in database. should pass in service method to fetch user
func Authorization(getUser func(string) (domain.User, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Get("id").(string)
			log.Println("User ID: ", userID)
			_, err := getUser(userID)
			if err != nil {
				return fmt.Errorf("access denied: user not registered or error: %s", err)
			}
			return next(c)
		}
	}
}

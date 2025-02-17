package authorization

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
)

// middleware to verify user is in database. should pass in service method to fetch user
func Authorization(getUser func(string) (domain.User, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			params := handlers.ParseCourseIDParams(c)
			log.SetPrefix("Authorization Middleware")
			userID := c.Get("id").(string)
			log.Println("User ID: ", userID)
			user, err := getUser(userID)
			if user.ID != userID || err != nil {
				return fmt.Errorf("access denied: user not registered")
			}
			if user.ID != params.UserID.Value {
				return fmt.Errorf("route userID and signed-in userID mismatch. routeID: %s; userID: %s", user.ID, params.UserID)
			}
			return next(c)
		}
	}
}

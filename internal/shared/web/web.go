package we

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/core/user"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"log"
	"net/http"
	"net/url"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Reverse func(name string, params ...any) string

func Respond(c echo.Context, redirect string, component, altComponent templ.Component) error {
	if altComponent == nil && redirect == "" {
		return fmt.Errorf("neither redirect or alt component provided in function call")
	}
	if !IsHTMX(c) {
		log.Println("request is NOT an HTMX request:", c.Request().Header.Get("Hx-Request"))
		if redirect != "" {
			log.Println("redirecting to: ", redirect)
			return c.Redirect(http.StatusFound, redirect)
		}
		if altComponent != nil {
			return altComponent.Render(context.Background(), c.Response())

		}
	}
	log.Println("request IS an HTMX request:", c.Request().Header.Get("Hx-Request"))
	return component.Render(context.Background(), c.Response())
}

func IsHTMX(c echo.Context) bool {
	// Check for "HX-Request" header
	return c.Request().Header.Get("Hx-Request") != ""
}

func EmptyUser() user.User {
	return user.User{
		Picture: AssetsURLFunc("signin.png"),
	}
}

func CourseManagerLayout(reverse Reverse, page templ.Component, user user.User) templ.Component {
	cml := managertemplates.CourseManagerLayout{
		HomeURL:    "/",
		Page:       page,
		User:       user,
		SigninURL:  reverse(routes.GetSignin.String()),
		SignupURL:  reverse(routes.GetSignup.String()),
		SignoutURL: reverse(routes.PostSignout.String()),
	}
	return cml.Component()

}

func AssetsURLFunc(path ...string) string {
	URL, err := url.JoinPath("/dist", path...)
	if err != nil {
		return "error joining paths"
	}
	return URL
}

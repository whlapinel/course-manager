package handlers

import (
	"fmt"
	auth "gh_static_portfolio/internal/authentication"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	Signin RouteName = "/signin"
	Signup RouteName = "/signup"
)

const (
	GetSignin  RouteHandlerName = RouteHandlerName(GET + Signin)
	PostSignin RouteHandlerName = RouteHandlerName(POST + Signin)
	GetSignup  RouteHandlerName = RouteHandlerName(GET + Signup)
	PostSignup RouteHandlerName = RouteHandlerName(POST + Signup)
)

func (h CourseHandler) AuthenticationHandlers() []RouteHandler {
	return []RouteHandler{
		{Signin, GetSignin, GET, h.GetSignin},
		{Signin, PostSignin, POST, h.PostSignin},
		{Signup, GetSignup, GET, h.GetSignup},
		{Signup, PostSignup, POST, h.PostSignup},
	}
}

func (h CourseHandler) GetSignin(c echo.Context) error {
	page := mt.SigninPage{
		GoogleSigninURL: h.e.Reverse(PostSignin.String()),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(page.Component())
	return Respond(c, "", component, layout)

}

func (h CourseHandler) PostSignin(c echo.Context) error {
	payload, err := auth.GoogleAuth(c)
	if err != nil {
		log.Println(err)
		return c.String(500, "Failed to authenticate")
	}
	sub := payload.Claims["sub"].(string)
	user, err := h.svc.GetUser(sub)
	if err != nil {
		log.Println(err)
		return c.String(500, err.Error())
	}
	t, err := auth.IssueToken(auth.TokenParams{User: domain.User{
		ID:        sub,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
	}})
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	expirationTime := time.Now().Add(auth.SessionLifeSpan).UnixMilli()
	expyString := strconv.Itoa(int(expirationTime))
	msg := fmt.Sprintf("%s, %d, %s", "congrats you did it!", int(expirationTime), expyString)
	return c.String(200, msg)
}

func (h CourseHandler) GetSignup(c echo.Context) error {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	page := mt.SignUpPage{
		GoogleClientID:  clientID,
		GoogleSignupURL: h.e.Reverse(GetSignup.String()),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(page.Component())
	return Respond(c, "", component, layout)
}
func (h CourseHandler) PostSignup(c echo.Context) error {
	payload, err := auth.GoogleAuth(c)
	if err != nil {
		log.Println("Failed to authenticate user: ", err)
		return c.String(401, "Failed to authenticate user")
	}
	id := payload.Claims["sub"].(string)
	email := payload.Claims["email"].(string)
	first := payload.Claims["given_name"].(string)
	last := payload.Claims["family_name"].(string)
	picture := payload.Claims["picture"].(string)
	// create user
	user, err := h.svc.SaveUser(service.SaveUserParams{
		User: domain.User{
			ID:        id,
			Email:     email,
			FirstName: first,
			LastName:  last,
			Picture:   picture,
		},
	})
	if err != nil {
		log.Println("Failed to create user: ", err)
		return c.String(500, "Failed to create user")
	}
	t, err := auth.IssueToken(auth.TokenParams{User: user})
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	expirationTime := time.Now().Add(auth.SessionLifeSpan).UnixMilli()
	expyString := strconv.Itoa(int(expirationTime))
	msg := fmt.Sprintf("%s, %d, %s", "congrats you did it!", int(expirationTime), expyString)
	return c.String(200, msg)
}

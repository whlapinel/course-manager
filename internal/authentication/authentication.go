package authentication

import (
	"context"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/domain"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

const SessionLifeSpan = time.Hour
const cushionTime = time.Minute * 5

var secret = os.Getenv("JWT_SECRET")

type JwtCustomClaims struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	ID      string `json:"id"`
	jwt.RegisteredClaims
}

func AddCookieToHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("running AddCookieToHeader middleware")
		cookie, err := c.Cookie("token")
		if err != nil {
			log.Println("error getting cookie", err)
		} else {
			c.Request().Header.Set("Authorization", "Bearer "+cookie.Value)
		}
		return next(c)
	}
}

var JWTMiddlewareProtectedNew = func(router *echo.Echo, signinRedirectRHN string) echo.MiddlewareFunc {
	var protectedConfig = echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(secret),
		ErrorHandler: func(c echo.Context, err error) error {
			// Redirect to login page on error
			log.Println("Error validating token: ", err)
			return c.Redirect(303, router.Reverse(signinRedirectRHN))
		},
		SuccessHandler: func(c echo.Context) {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(*JwtCustomClaims)
			expiration := claims.ExpiresAt.Time
			if time.Until(expiration) <= cushionTime {
				t, err := IssueToken(TokenParams{User: domain.User{
					ID:        claims.ID,
					FirstName: claims.First,
					LastName:  claims.Last,
					Email:     claims.Email,
					Picture:   claims.Picture,
				}})
				if err != nil {
					log.Println("Failed to issue token: ", err)
				}
				WriteToken(c, t)
				jsonString := fmt.Sprintf("{\"signin\":{\"expiration\":%d}}", time.Now().Add(SessionLifeSpan).UnixMilli())
				c.Response().Header().Set("Hx-Trigger", jsonString)
			} else {
				log.Println("more than a minute left")
			}
		},
	}
	return echojwt.WithConfig(protectedConfig)
}

func GetClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("running GetClaims middleware")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		c.Set("id", claims.ID)
		c.Set("first", claims.First)
		c.Set("last", claims.Last)
		c.Set("email", claims.Email)
		c.Set("picture", claims.Picture)
		return next(c)
	}
}

func GoogleAuth(c echo.Context) (*idtoken.Payload, error) {
	token, err := c.Cookie("g_csrf_token")
	if err != nil {
		return nil, errors.New("token not found")
	}
	bodyToken := c.FormValue("g_csrf_token")
	if token.Value != bodyToken {
		return nil, errors.New("token mismatch")
	}
	ctx := context.Background()
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		log.Println("Failed to create ID token validator: ", err)
		return nil, errors.New("failed to create ID token validator")
	}
	credential := c.FormValue("credential")
	payload, err := validator.Validate(ctx, credential, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		log.Println("Failed to validate ID token: ", err)
		return nil, errors.New("failed to validate ID token")
	}
	log.Println("Payload: ", payload)
	return payload, nil
}

type TokenParams struct {
	domain.User
}

func IssueToken(params TokenParams) (string, error) {
	claims := JwtCustomClaims{
		ID:      params.ID,
		First:   params.FirstName,
		Last:    params.LastName,
		Email:   params.Email,
		Picture: params.Picture,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(SessionLifeSpan)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func WriteToken(c echo.Context, t string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(SessionLifeSpan)
	log.Println("Setting cookie: ", cookie)
	c.SetCookie(cookie)
	c.Response().Header().Set("Authorization", "Bearer "+t)
}

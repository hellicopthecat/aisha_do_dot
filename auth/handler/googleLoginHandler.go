package handler

import (
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func (h AuthHandler) GoogleLoginHandler(e *echo.Context) error {
	uid := uuid.NewString()
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")

	e.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    uid,
		MaxAge:   300,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // 배포 후 true
		SameSite: http.SameSiteLaxMode,
	})

	params := url.Values{}
	params.Set("client_id", googleClientId)
	params.Set("redirect_uri", "http://localhost:8080/api/auth/google/login/response")
	params.Set("response_type", "code")
	params.Set("scope", "email openid profile")
	params.Set("state", uid)

	loginUrl := "https://accounts.google.com/o/oauth2/v2/auth?" + params.Encode()
	return e.Redirect(http.StatusFound, loginUrl)
}

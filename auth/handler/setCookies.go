package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func setCookies(at string, rt string, e *echo.Context) {
	e.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    at,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 3,
		HttpOnly: true,
		Secure:   false, // 배포시 true
		SameSite: http.SameSiteLaxMode,
	})
	e.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    rt,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
		Secure:   false, // 배포시 true
		SameSite: http.SameSiteLaxMode,
	})
}

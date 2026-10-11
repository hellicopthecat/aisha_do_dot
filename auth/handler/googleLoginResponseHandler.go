package handler

import (
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v5"
)

func (h AuthHandler) GoogleLoginResponseHandler(e *echo.Context) error {

	stateCookie, err := e.Cookie("state")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "잘못된 요청입니다.",
		})
	}
	stateParam := e.Param("state")
	codeParam := e.Param("code")

	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	secretKey := os.Getenv("GOOGLE_SECRET_KEY")

	e.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // 배포 시 true
		SameSite: http.SameSiteLaxMode,
	})

	if stateCookie.Value != stateParam {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "잘못된 요청입니다.",
		})
	}

	if codeParam == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "잘못된 요청입니다.",
		})
	}

	qp := url.Values{}
	qp.Add("client_id", clientId)
	qp.Add("client_secret", secretKey)
	qp.Add("code", codeParam)
	qp.Add("grant_type", "authorization_code")
	qp.Add("redirect_uri", "http://localhost:8080/api/auth/google/login/response")

	return nil
}

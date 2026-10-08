package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	aishaJWT "github.com/hellicopthecat/aisha_do_dot/jwt"
	dtos "github.com/hellicopthecat/aisha_do_dot/users/dtos/request"
	"github.com/labstack/echo/v5"
)

func (h UserHandler) UpdateRefreshTokenHandler(e *echo.Context) error {
	ctx := e.Request().Context()
	// TODO: jwt 복호화 및 id 추출

	token, err := echo.ContextGet[*jwt.Token](e, "user")
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"status":  "FAILED",
			"message": "인증된 유저가 아닙니다.",
		})
	}
	claims, ok := token.Claims.(*aishaJWT.AccessClaims)
	if !ok {
		return echo.ErrUnauthorized
	}

	var dto dtos.UpdateRefreshToken
	if err := e.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	dto.Id = claims.Id

	err = h.userHandler.UpdateRefreshTokenService(ctx, dto)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"status":  "FAILED",
			"message": "인증된 유저가 아닙니다.",
		})
	}
	return e.JSON(http.StatusOK, map[string]string{
		"status":  "SUCCESS",
		"message": "Refresh Token 발급 성공",
	})
}

package handler

import (
	"net/http"

	dtos "github.com/hellicopthecat/aisha_do_dot/users/dtos/request"
	"github.com/labstack/echo"
)

func (h UserHandler) UpdateRefreshTokenHandler(e echo.Context) error {
	var dto dtos.UpdateRefreshToken
	ctx := e.Request().Context()
	// TODO: jwt 복호화 및 id 추출
	e.Bind(&dto)
	err := h.userHandler.UpdateRefreshTokenService(ctx, dto)
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

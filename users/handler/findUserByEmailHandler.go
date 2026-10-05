package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (h UserHandler) FindUserByEmailHandler(e echo.Context) error {
	ctx := e.Request().Context()

	email := e.QueryParam("email")

	userEmail, err := h.userHandler.FindUserByEmailService(ctx, email)
	if err != nil {
		log.Printf("Find User Failed :: email=%s || error=%v", *userEmail, err)
		return e.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"status":  "FAILED",
				"message": "서버에 오류가 발생되었습니다..",
			},
		)
	}
	if userEmail == nil {
		return e.JSON(
			http.StatusNotFound,
			map[string]string{
				"status":  "FAILED",
				"message": "조회되는 유저의 이메일이 없습니다.",
			},
		)
	}
	return e.JSON(http.StatusOK,
		map[string]string{
			"status":  "SUCCESS",
			"message": "유저 조회에 성공했습니다.",
			"data":    *userEmail,
		})
}

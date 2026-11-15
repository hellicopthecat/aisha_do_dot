package handler

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	aishaJwt "github.com/hellicopthecat/aisha_do_dot/jwt"
	"github.com/labstack/echo/v5"
)

func (s RoutineHandler) CreateRoutineHandler(e *echo.Context) error {
	// ctx := e.Request().Context()
	uu, _ := echo.ContextGet[*jwt.Token](e, "user")
	// fmt.Println(ctx.)
	claim, _ := aishaJwt.ParseAccessToken(uu.Raw)
	fmt.Println("hoho :: ", claim.Id)
	fmt.Println("hoho :: ", claim.Email)

	// //

	// s.routineService.CreateRoutineService(ctx,)
	return nil
}

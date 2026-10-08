package users

import (
	"database/sql"

	"github.com/hellicopthecat/aisha_do_dot/jwt"
	"github.com/hellicopthecat/aisha_do_dot/users/handler"
	"github.com/hellicopthecat/aisha_do_dot/users/repo"
	"github.com/hellicopthecat/aisha_do_dot/users/service"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type UserModule struct {
	userHandler handler.UserHandler
}

func InitUserModule(db *sql.DB) *UserModule {

	userRepo := repo.InitUserRepo(db)
	userService := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userService)

	return &UserModule{
		userHandler: userHandler,
	}
}

func (m UserModule) UserGroupApi(r *echo.Echo) {
	userGroup := r.Group("/user")
	config := jwt.EchoJwtConfig()
	userGroup.GET("/", m.userHandler.FindUserByEmailHandler)
	userGroup.PATCH("/refresh-token", m.userHandler.UpdateRefreshTokenHandler, echojwt.WithConfig(config))
}

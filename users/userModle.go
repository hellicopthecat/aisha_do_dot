package users

import (
	"database/sql"

	"github.com/hellicopthecat/aisha_do_dot/users/handler"
	"github.com/hellicopthecat/aisha_do_dot/users/repo"
	"github.com/hellicopthecat/aisha_do_dot/users/service"
	"github.com/labstack/echo"
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
	userGroup.GET("/", m.userHandler.FindUserByEmailHandler)
}

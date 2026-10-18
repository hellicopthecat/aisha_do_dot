package auth

import (
	"database/sql"

	"github.com/hellicopthecat/aisha_do_dot/auth/handler"
	authRepo "github.com/hellicopthecat/aisha_do_dot/auth/repo"
	authService "github.com/hellicopthecat/aisha_do_dot/auth/service"
	"github.com/labstack/echo/v5"

	userRepo "github.com/hellicopthecat/aisha_do_dot/users/repo"
	userService "github.com/hellicopthecat/aisha_do_dot/users/service"
)

type AuthModule struct {
	authHandler handler.AuthHandler
}

func InitAuthModule(db *sql.DB) *AuthModule {
	userInitRepo := userRepo.InitUserRepo(db)
	userInitService := userService.InitUserService(userInitRepo)

	authInitRepo := authRepo.InitAuthRepo(db)
	authInitService := authService.InitAuthService(authInitRepo)

	authHandler := handler.InitAuthHandler(authInitService, userInitService)

	return &AuthModule{
		authHandler: *authHandler,
	}
}

func (m AuthModule) AuthGroupApi(r *echo.Echo) {
	authGroup := r.Group("/api/auth")
	authGroup.GET("/google/login", m.authHandler.GoogleLoginHandler)
	authGroup.GET("/google/login/response", m.authHandler.GoogleLoginResponseHandler)
}

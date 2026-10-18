package handler

import (
	authService "github.com/hellicopthecat/aisha_do_dot/auth/service"
	userService "github.com/hellicopthecat/aisha_do_dot/users/service"
)

type AuthHandler struct {
	userHandler *userService.UserService
	authHandler *authService.AuthService
}

func InitAuthHandler(
	authService *authService.AuthService,
	userService *userService.UserService) *AuthHandler {
	return &AuthHandler{
		userHandler: userService,
		authHandler: authService,
	}
}

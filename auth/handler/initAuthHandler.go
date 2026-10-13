package handler

import "github.com/hellicopthecat/aisha_do_dot/users/service"

type AuthHandler struct {
	authHandler service.UserService
}

func InitAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{
		authHandler: userService,
	}
}

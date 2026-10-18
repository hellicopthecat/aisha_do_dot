package service

import "github.com/hellicopthecat/aisha_do_dot/auth/repo"

type AuthService struct {
	authRepo *repo.AuthRepo
}

func InitAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{
		authRepo: repo,
	}
}

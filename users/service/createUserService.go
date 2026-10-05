package service

import (
	"context"

	"github.com/hellicopthecat/aisha_do_dot/users/entity"
)

func (s UserService) CreateUserSerivce(ctx context.Context, userEntity entity.User) error {
	err := s.userService.CreateUserRepo(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

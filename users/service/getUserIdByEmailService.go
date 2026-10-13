package service

import (
	"context"

	"github.com/google/uuid"
)

func (s UserService) GetUserIdByEmailService(ctx context.Context, email string) (*uuid.UUID, error) {
	return s.userService.GetUserIdByEmailRepo(ctx, email)
}

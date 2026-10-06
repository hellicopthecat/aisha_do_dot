package service

import (
	"context"

	dtos "github.com/hellicopthecat/aisha_do_dot/users/dtos/request"
)

func (s UserService) UpdateRefreshTokenService(ctx context.Context, dto dtos.UpdateRefreshToken) error {
	err := s.UpdateRefreshTokenService(ctx, dto)
	return err
}

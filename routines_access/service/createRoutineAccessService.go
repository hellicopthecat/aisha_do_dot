package service

import (
	"context"

	"github.com/hellicopthecat/aisha_do_dot/routines_access/dtos"
)

func (s RoutineAccessService) CreateRoutineAccessService(ctx context.Context, dto dtos.CreateRoutineAccessDto) error {
	return s.repo.CreateRoutineAccessRepo(ctx, dto)
}

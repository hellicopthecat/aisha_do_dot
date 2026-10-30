package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (s RoutineService) FindMyRoutineByUserIdService(ctx context.Context, userId uuid.UUID) (*[]dtos.FindMyRoutineByUserIdDto, error) {
	return s.routineRepo.FindMyRoutineByUserIdRepo(ctx, userId)
}

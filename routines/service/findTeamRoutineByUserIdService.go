package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (s RoutineService) FindTeamRoutineByUserIdService(ctx context.Context, userId uuid.UUID) (*[]dtos.FindMyRoutineByUserIdDto, error) {
	return s.routineRepo.FindTeamRoutineByUserIdRepo(ctx, userId)
}

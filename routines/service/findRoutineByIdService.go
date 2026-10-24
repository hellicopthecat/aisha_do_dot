package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
	"github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (r *RoutineService) FindRoutineByIdService(ctx context.Context, user_id uuid.UUID, routine_id uuid.UUID) (*dtos.FindRoutineByIdDto, error) {
	routine, err := r.routineRepo.FindRoutineByIdRepo(ctx, routine_id)
	if err != nil {
		return nil, err
	}

	for _, a_user := range routine.Access_User {
		if a_user.Access_User_Id == user_id && a_user.Can_Access == enums.AccessAccess {
			return routine, nil
		}
	}

	return nil, fmt.Errorf("접근할 수 없는 권한입니다.")
}

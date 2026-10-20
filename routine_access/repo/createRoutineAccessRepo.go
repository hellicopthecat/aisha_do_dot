package repo

import (
	"context"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routine_access/dtos"
)

func (r *RoutineAccessRepo) CreateRoutineAccessRepo(ctx context.Context, dto dtos.CreateRoutineAccessDto) error {
	file, err := os.ReadFile("db/ddls/routine_access/createRoutineAccess.sql")
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, string(file),
		dto.Id,
		dto.User_Email,
		dto.User_Email,
		dto.Access,
		dto.Editable,
		dto.Routine_Id,
		dto.Access_User_Id,
	)

	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		return err
	}

	return nil
}

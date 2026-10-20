package repo

import (
	"context"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (r *RoutineRepo) CreateRoutineRepo(ctx context.Context, dto dtos.CreateRoutineDto) error {
	file, err := os.ReadFile("db/ddls/routines/createRoutine.sql")
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, string(file),
		dto.Id,
		dto.Email,
		dto.Email,
		dto.Routine_Title,
		dto.Routine_Desc,
		dto.Priority_One_Depth,
		dto.Category,
		dto.Owner_Id,
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

	return err
}

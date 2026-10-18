package repo

import (
	"context"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routines/entity"
)

func (r RoutineRepo) CreateRoutineRepo(ctx context.Context, routine *entity.Routine) error {
	file, err := os.ReadFile("db/ddls/routines/createRoutine.sql")
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, string(file),
		routine.Id,
		routine.Owner.Email,
		routine.Owner.Email,
		routine.Routine_Title,
		routine.Routine_Desc,
		routine.Priority_One_Depth,
		routine.Category,
		routine.Owner.Id,
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

	return tx.Commit()
}

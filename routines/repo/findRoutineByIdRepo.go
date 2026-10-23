package repo

import (
	"context"
	"database/sql"
	"os"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (r *RoutineRepo) FindRoutineByIdRepo(ctx context.Context, id uuid.UUID) (*dtos.FindRoutineByIdDto, error) {
	file, err := os.ReadFile("db/ddls/routines/findRoutineById.sql")
	if err != nil {
		return nil, err
	}
	var routine dtos.FindRoutineByIdDto
	err = r.db.QueryRowContext(ctx, string(file), id).Scan(
		&routine.Id,
		&routine.Created_At,
		&routine.Updated_At,
		&routine.Created_By,
		&routine.Updated_By,
		&routine.Routine_Title,
		&routine.Routine_Desc,
		&routine.Priority_One_Depth,
		&routine.Category,
		&routine.User_Id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &routine, nil
}

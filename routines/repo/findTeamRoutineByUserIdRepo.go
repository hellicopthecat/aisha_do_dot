package repo

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (r RoutineRepo) FindTeamRoutineByUserIdRepo(ctx context.Context, userId uuid.UUID) (*[]dtos.FindMyRoutineByUserIdDto, error) {
	file, err := os.ReadFile("db/ddls/routines/findTeamRoutineByUserId.sql")
	if err != nil {
		return nil, err
	}
	result, err := r.db.QueryContext(ctx, string(file), userId.String())
	if err != nil {
		return nil, err
	}
	var routines []dtos.FindMyRoutineByUserIdDto
	for result.Next() {
		var dto dtos.FindMyRoutineByUserIdDto
		err := result.Scan(
			&dto.Id,
			&dto.Created_At,
			&dto.Updated_At,
			&dto.Created_By,
			&dto.Updated_By,
			&dto.Routine_Title,
			&dto.Routine_Desc,
			&dto.Priority_One_Depth,
			&dto.Category,
			&dto.User_Id,
		)
		if err != nil {
			return nil, err
		}
		routines = append(routines, dto)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return &routines, nil
}

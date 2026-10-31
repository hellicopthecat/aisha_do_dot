package repo

import (
	"context"
	"encoding/json"
	"os"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/routines_detail/dtos"
)

func (r *RoutineDetailRepo) FindRoutineDetailByRoutineIdRepo(ctx context.Context, routine_id uuid.UUID) (*[]dtos.FindRoutineDetailByRoutineIdDto, error) {
	file, err := os.ReadFile("db/ddls/routines_detail/findRoutinesDetailByRoutineId.sql")

	if err != nil {
		return nil, err
	}

	result, err := r.db.QueryContext(ctx, string(file), routine_id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	details := make([]dtos.FindRoutineDetailByRoutineIdDto, 0)

	for result.Next() {
		var detail dtos.FindRoutineDetailByRoutineIdDto
		var tagsJson, decriptionsJson []byte
		err := result.Scan(
			&detail.Id,
			&detail.Created_At,
			&detail.Updated_At,
			&detail.Created_By,
			&detail.Updated_By,
			&detail.Priority_Two_Depth,
			&detail.Pre_Event_Start_At,
			&detail.Pre_Event_End_At,
			&detail.Start_At,
			&detail.End_At,
			&detail.Routine_Id,
			&tagsJson,
			&decriptionsJson,
		)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(tagsJson, &detail.Tags); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(decriptionsJson, &detail.Detail_Desc); err != nil {
			return nil, err
		}
		details = append(details, detail)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return &details, nil
}

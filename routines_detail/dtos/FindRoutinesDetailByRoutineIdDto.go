package dtos

import (
	"time"

	"github.com/google/uuid"
)

type FindRoutineDetailByRoutineIdDto struct {
	Id                 uuid.UUID           `json:"id"`
	Created_At         time.Time           `json:"created_at"`
	Updated_At         time.Time           `json:"updated_at"`
	Created_By         string              `json:"created_by"`
	Updated_By         string              `json:"updated_by"`
	Priority_Two_Depth int                 `json:"priority_two_depth"`
	Pre_Event_Start_At time.Time           `json:"pre_event_start_at"`
	Pre_Event_End_At   time.Time           `json:"pre_event_end_at"`
	Start_At           time.Time           `json:"start_at"`
	End_At             time.Time           `json:"end_at"`
	Routine_Id         uuid.UUID           `json:"routine_id"`
	Tags               []routineTags       `json:"tags"`
	Detail_Desc        []routineDetailDesc `json:"detail_desc"`
}

type routineTags struct {
	Id  uuid.UUID `json:"id"`
	Tag string    `json:"tag"`
}

type routineDetailDesc struct {
	Id          uuid.UUID `json:"id"`
	Detail_Desc string    `json:"detail_desc"`
}

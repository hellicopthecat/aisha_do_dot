package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateRoutineDetailDto struct {
	Id                 uuid.UUID
	Init_User_Email    string
	Priority_Two_Depth int
	Pre_Event_Start_At time.Time
	Pre_Event_End_At   time.Time
	Start_At           time.Time
	End_At             time.Time
	Routine_Id         uuid.UUID
}

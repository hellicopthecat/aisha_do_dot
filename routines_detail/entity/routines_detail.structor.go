package entity

import (
	"time"

	"github.com/google/uuid"
)

type RoutineDetail struct {
	Id                 uuid.UUID
	Created_At         time.Time
	Updated_At         time.Time
	Created_By         string
	Updated_By         string
	Priority_Two_Depth int
	Pre_Event_Start_At time.Time
	Pre_Event_End_At   time.Time
	Start_At           time.Time
	End_At             time.Time
	Routine_Id         uuid.UUID
}

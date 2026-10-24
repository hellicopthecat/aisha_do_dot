package entity

import (
	"time"

	"github.com/google/uuid"
)

type RoutineDetailDesc struct {
	Id                uuid.UUID
	Created_At        time.Time
	Updated_At        time.Time
	Created_By        string
	Updated_By        string
	Detail_Desc       string
	Routine_Detail_Id uuid.UUID
}

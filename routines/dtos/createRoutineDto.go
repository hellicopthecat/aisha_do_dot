package dtos

import (
	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type CreateRoutineDto struct {
	Id                 uuid.UUID
	Email              string
	Routine_Title      string
	Routine_Desc       string
	Priority_One_Depth int64
	Category           enums.OneDepth
	Owner_Id           uuid.UUID
}

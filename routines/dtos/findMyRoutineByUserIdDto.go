package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type FindMyRoutineByUserIdDto struct {
	Id                 uuid.UUID
	Created_At         time.Time
	Updated_At         time.Time
	Created_By         string
	Updated_By         string
	Routine_Title      string
	Routine_Desc       string
	Priority_One_Depth int
	Category           enums.OneDepth
	User_Id            uuid.UUID
}

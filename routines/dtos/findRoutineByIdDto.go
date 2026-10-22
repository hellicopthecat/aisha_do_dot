package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type FindRoutineById struct {
	Id                 uuid.UUID      //id
	Created_At         time.Time      //created_at
	Updated_At         time.Time      //updated_at
	Created_By         string         //created_by
	Updated_By         string         //updated_by
	Routine_Title      string         //routine_title
	Routine_Desc       *string        //routine_desc
	Priority_One_Depth *int           //priorityOneDepth
	Category           enums.OneDepth //category
	User_Id            uuid.UUID      //user_id
}

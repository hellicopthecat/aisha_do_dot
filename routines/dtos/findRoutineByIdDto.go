package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type FindRoutineByIdDto struct {
	Id                 uuid.UUID      //id
	Created_At         time.Time      //created_at
	Updated_At         time.Time      //updated_at
	Created_By         string         //created_by
	Updated_By         string         //updated_by
	Routine_Title      string         //routine_title
	Routine_Desc       *string        //routine_desc
	Priority_One_Depth *int           //priority_one_depth
	Category           enums.OneDepth //category
	User_Id            uuid.UUID      //user_id
	User_Email         string
	User_Name          string
	Access_User        []FindRoutineByIdAccessDto
}

type FindRoutineByIdAccessDto struct {
	Can_Access     enums.Access     `json:"can_access"`
	Editable       enums.AccessRole `json:"editable"`
	Access_User_Id uuid.UUID        `json:"access_user_id"`
}

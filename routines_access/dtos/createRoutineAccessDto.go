package dtos

import (
	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type CreateRoutineAccessDto struct {
	Id             uuid.UUID
	User_Email     string
	Access         enums.Access
	Editable       enums.AccessRole
	Routine_Id     uuid.UUID
	Access_User_Id uuid.UUID
}

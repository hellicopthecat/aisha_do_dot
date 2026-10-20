package entity

import (
	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
	routineAccessEntity "github.com/hellicopthecat/aisha_do_dot/routine_access/entity"
	userEntity "github.com/hellicopthecat/aisha_do_dot/users/entity"
)

type Routine struct {
	Id                 uuid.UUID
	Routine_Title      string
	Routine_Desc       *string
	Priority_One_Depth *int
	Category           enums.OneDepth
	Owner              userEntity.User // Routine_Owner
	Routine_Access     []routineAccessEntity.RoutineAccessEntity
	// Routine_Detail
}

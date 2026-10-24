package dtos

import "github.com/google/uuid"

type CreateRoutineDetailDescDto struct {
	Id                uuid.UUID
	Init_User_Email   string
	Detail_Desc       string
	Routine_Detail_Id uuid.UUID
}

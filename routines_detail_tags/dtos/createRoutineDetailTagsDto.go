package dtos

import "github.com/google/uuid"

type CreateRoutineDetailTagsDto struct {
	Id                uuid.UUID
	Init_User_Email   string
	Tag               string
	Routine_Detail_Id uuid.UUID
}

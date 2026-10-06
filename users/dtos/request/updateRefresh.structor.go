package dtos

import "github.com/google/uuid"

type UpdateRefreshToken struct {
	Id            uuid.UUID
	Refresh_token string
}

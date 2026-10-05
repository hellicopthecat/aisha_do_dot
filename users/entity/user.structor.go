package entity

import (
	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type User struct {
	Id            uuid.UUID
	Email         string
	Social        enums.Social
	Provider_id   string
	Name          string
	Refresh_token string
}

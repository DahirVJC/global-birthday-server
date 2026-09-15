package models

import (
	"github.com/google/uuid"
)

type UserReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
}

type UserRes struct {
	ID       uuid.UUID `json:"_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Timezone string    `json:"timezone"`
}

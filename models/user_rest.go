package models

import (
	"github.com/google/uuid"
)

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Timezone string `json:"timezone" binding:"required"`
}

type UserRes struct {
	ID       uuid.UUID `json:"_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Timezone string    `json:"timezone"`
}

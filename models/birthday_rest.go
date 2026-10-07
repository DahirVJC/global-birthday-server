package models

import (
	"time"

	"github.com/google/uuid"
)

type BirthdayReq struct {
	Name         string     `json:"name" binding:"required"`
	Birthdate    time.Time  `json:"birthdate" binding:"required"`
	Timezone     string     `json:"timezone" binding:"required"`
	RemindMeIn   *time.Time `json:"remindMeIn,omitempty"`
	ContactLinks []string   `json:"contactLinks,omitempty"`
	Wishlists    []string   `json:"wishlists,omitempty"`
}

type BirthdayRes struct {
	ID           uuid.UUID   `json:"_id"`
	Name         string      `json:"name"`
	Birthdate    time.Time   `json:"birthdate"`
	Timezone     string      `json:"timezone"`
	RemindMeIn   *time.Time  `json:"remindMeIn,omitempty"`
	ContactLinks []string    `json:"contactLinks"`
	Wishlists    []string    `json:"wishlists"`
	Events       []uuid.UUID `json:"events"`
	StartDate    time.Time   `json:"startDate"`
	EndDate      time.Time   `json:"endDate"`
}

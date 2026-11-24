package models

import (
	"time"

	"github.com/google/uuid"
)

type EventReq struct {
	BirthdayID     uuid.UUID  `json:"birthdayId" binding:"required"`
	EventDateStart time.Time  `json:"eventDateStart" binding:"required"`
	EventDateEnd   *time.Time `json:"eventDateEnd,omitempty"`
	RemindMeIn     *time.Time `json:"remindMeIn,omitempty"`
	Description    *string    `json:"description,omitempty"`
}

type EventRes struct {
	ID             uuid.UUID  `json:"_id"`
	BirthdayID     uuid.UUID  `json:"birthdayId"`
	EventDateStart time.Time  `json:"eventDateStart"`
	EventDateEnd   *time.Time `json:"eventDateEnd,omitempty"`
	RemindMeIn     *time.Time `json:"remindMeIn,omitempty"`
	Description    *string    `json:"description,omitempty"`
	StartDate      time.Time  `json:"startDate"`
	EndDate        time.Time  `json:"endDate"`
}

package events

import "time"

type Event struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description" validate:"required"`
	Type        string    `json:"type" validate:"required"`
}

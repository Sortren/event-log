package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Event struct {
	bun.BaseModel `bun:"table:events"`

	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description" validate:"required"`
	Type        string    `json:"type" validate:"required"`
}

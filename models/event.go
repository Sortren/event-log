package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Event struct {
	bun.BaseModel `bun:"table:events"`

	ID          uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt   time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	Description string    `json:"description"`
	Type        string    `json:"type" `
}

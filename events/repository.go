package events

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) WithCreatedAtRange(start string, end string) *Repository {
	r.db.Where("created_at BETWEEN ? AND ?", start, end)
	return r
}

func (r *Repository) WithType(eventType string) *Repository {
	r.db.Where("type = ?", eventType)
	return r
}

func (r *Repository) OrderBy(by string) *Repository {
	r.db.Order(fmt.Sprintf("%s DESC", by))
	return r
}

func (r *Repository) Find() []Event {
	var event []Event
	r.db.Find(&event)
	return event
}

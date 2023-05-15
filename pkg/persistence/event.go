package persistence

import (
	"github.com/Sortren/event-log/models"
	"github.com/uptrace/bun"
)

type Event interface {
	IRepository[models.Event]
}

func NewEvent(db bun.IDB) Event {
	return NewRepository[models.Event](db)
}

func EventWithCreatedAtRange(start string, end string) SelectCriteria {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("created_at BETWEEN ? AND ?", start, end)
	}
}

func EventWithType(eventType string) SelectCriteria {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("type = ?", eventType)
	}
}

package persistence

import (
	"github.com/Sortren/event-log/models"
	"github.com/uptrace/bun"
	"time"
)

type Event interface {
	IRepository[models.Event]
}

func NewEvent(db bun.IDB) Event {
	return NewRepository[models.Event](db)
}

func EventWithCreatedAtRange(start time.Time, end time.Time) SelectCriteria {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		if end.IsZero() {
			end = time.Now()
		}

		return q.Where("created_at BETWEEN ? AND ?", start, end)
	}
}

func EventWithType(eventType string) SelectCriteria {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("type = ?", eventType)
	}
}

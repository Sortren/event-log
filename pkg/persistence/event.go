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

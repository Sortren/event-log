package services

import (
	"context"
	"fmt"
	"github.com/Sortren/event-log/models"
	"github.com/Sortren/event-log/pkg/persistence"
	"time"
)

type IEventService interface {
	GetEvents(ctx context.Context, start time.Time, end time.Time, eventType string, limit int, offset int) ([]models.Event, error)
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
}

var _ IEventService = &EventService{}

type EventService struct {
	repo persistence.IRepository[models.Event]
}

func NewEventService(repo persistence.IRepository[models.Event]) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (e *EventService) GetEvents(ctx context.Context, start time.Time, end time.Time, eventType string, limit int, offset int) ([]models.Event, error) {
	events, err := e.repo.FindAll(
		ctx,
		persistence.EventWithCreatedAtRange(start, end),
		persistence.EventWithType(eventType),
		persistence.WithLimit(limit),
		persistence.WithOffset(offset),
	)

	if err != nil {
		return nil, fmt.Errorf("can't get events from repository, err: %w", err)
	}

	return events, nil
}

func (e *EventService) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := e.repo.Save(ctx, event); err != nil {
		return nil, fmt.Errorf("can't save event to the database")
	}

	return event, nil
}

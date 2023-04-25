package services

import (
	"github.com/Sortren/event-log/events"
	"github.com/Sortren/event-log/models"
)

type IEventService interface {
	GetEvents(start string, end string, eventType string, limit int, offset int) ([]events.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
}

var _ IEventService = &EventService{}

type EventService struct {
	repo events.Repository
}

func NewEventService(repo events.Repository) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (e *EventService) GetEvents(start string, end string, eventType string, limit int, offset int) ([]events.Event, error) {
	events := e.repo.
		WithCreatedAtRange(start, end).
		WithType(eventType).
		OrderBy("created_at").
		Find()

	return events, nil
}

func (e *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	//db := database.DBConn
	//
	//db.Create(&event)
	//
	//log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	return event, nil
}

package services

import (
	"log"

	"github.com/Sortren/event-log/src/database"
	"github.com/Sortren/event-log/src/events"
	"github.com/Sortren/event-log/src/models"
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
	// db := database.DBConn

	// var events []models.Event
	events := e.repo.
		WithCreatedAtRange(start, end).
		WithType(eventType).
		OrderBy("created_at").
		Find()

	// if utils.IsFilterPresent(eventType) {
	// 	db = db.Where("type = ?", eventType)
	// }
	// if utils.IsFilterPresent(start) && utils.IsFilterPresent(end) {
	// 	db = db.Where("created_at BETWEEN ? AND ?", start, end)
	// }

	// db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

	return events, nil
}

func (e *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	db := database.DBConn

	db.Create(&event)

	log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	return event, nil
}

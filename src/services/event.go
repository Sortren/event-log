package services

import (
	"log"

	"github.com/Sortren/event-log/src/database"
	"github.com/Sortren/event-log/src/models"
	"github.com/Sortren/event-log/src/utils"
)

type IEventService interface {
	GetEvents(start string, end string, eventType string, limit int, offset int) ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
}

type EventService struct{}

func (e *EventService) GetEvents(start string, end string, eventType string, limit int, offset int) ([]models.Event, error) {
	db := database.DBConn

	var events []models.Event

	if utils.IsFilterPresent(eventType) {
		db = db.Where("type = ?", eventType)
	}
	if utils.IsFilterPresent(start) && utils.IsFilterPresent(end) {
		db = db.Where("created_at BETWEEN ? AND ?", start, end)
	}

	db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

	return events, nil
}

func (e *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	db := database.DBConn

	db.Create(&event)

	log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	return event, nil
}

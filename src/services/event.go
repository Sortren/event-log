package services

import (
	"log"

	"github.com/Sortren/event-log/src/database"
	"github.com/Sortren/event-log/src/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type EventService interface {
	GetEvents(filters map[string]string) ([]map[string]interface{}, error)
	CreateEvent(event *models.Event) (*models.Event, error)
}

func GetEvents(filters map[string]string) ([]map[string]interface{}, error) {
	db := database.DBConn

	start := filters["start"]
	end := filters["end"]
	eventType := filters["type"]

	isStartFilterPresent := start != ""
	isEndFilterPresent := end != ""
	isTypeFilterPresent := eventType != ""

	// Providing start without end or end without start resolves Bad Request
	if isStartFilterPresent != isEndFilterPresent {
		return nil, fiber.ErrBadRequest
	}

	var err error
	var events []map[string]interface{}

	if isTypeFilterPresent && (isStartFilterPresent && isEndFilterPresent) {
		db.Model(&models.Event{}).Where("type = ? AND created_at BETWEEN ? AND ?", eventType, start, end).Find(&events)

	} else if isStartFilterPresent && isEndFilterPresent {
		db.Model(&models.Event{}).Where("created_at BETWEEN ? AND ?", start, end).Find(&events)

	} else if isTypeFilterPresent {
		db.Model(&models.Event{}).Where("type = ?", eventType).Find(&events)
	}

	return events, err
}

func CreateEvent(event *models.Event) (*models.Event, error) {
	db := database.DBConn

	validate := validator.New()
	if err := validate.Struct(event); err != nil {
		return nil, err
	}

	db.Create(&event)

	log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	var err error

	return event, err
}

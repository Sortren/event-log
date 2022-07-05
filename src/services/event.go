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

func GetEvents(filters map[string]interface{}) ([]models.Event, error) {
	db := database.DBConn

	start := filters["Start"]
	end := filters["End"]
	eventType := filters["Type"]

	limit, ok := filters["Limit"].(int)

	if !ok {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Wrong limit queryparam")
	}

	offset, ok := filters["Offset"].(int)

	if !ok {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Wrong offset queryparam")
	}

	isStartFilterPresent := start != ""
	isEndFilterPresent := end != ""
	isTypeFilterPresent := eventType != ""

	if isStartFilterPresent != isEndFilterPresent {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Can't provide start without end and end without start")
	}

	var events []models.Event

	db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

	if isTypeFilterPresent {
		db.Where("type = ?", eventType)
	}
	if isStartFilterPresent && isEndFilterPresent {
		db.Where("created_at BETWEEN ? AND ?", start, end)
	}

	return events, nil
}

func CreateEvent(event *models.Event) (*models.Event, error) {
	db := database.DBConn

	validate := validator.New()
	if err := validate.Struct(event); err != nil {
		return nil, err
	}

	db.Create(&event)

	log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	return event, nil
}

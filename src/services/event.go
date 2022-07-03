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

func GetEvents(filters map[string]interface{}) ([]map[string]interface{}, error) {
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

	var events []map[string]interface{}

	if isTypeFilterPresent && (isStartFilterPresent && isEndFilterPresent) {
		db.Model(&models.Event{}).Where("type = ? AND created_at BETWEEN ? AND ?", eventType, start, end).Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

	} else if isStartFilterPresent && isEndFilterPresent {
		db.Model(&models.Event{}).Where("created_at BETWEEN ? AND ?", start, end).Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

	} else if isTypeFilterPresent {
		db.Model(&models.Event{}).Where("type = ?", eventType).Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

	} else {
		db.Model(&models.Event{}).Order("created_at DESC").Limit(limit).Offset(offset).Find(&events)

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

package services

import (
	"fmt"
	"github.com/Sortren/event-log/models"
	"github.com/Sortren/event-log/pkg/persistence"
)

type IEventService interface {
	GetEvents(start string, end string, eventType string, limit int, offset int) ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
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

func (e *EventService) GetEvents(start string, end string, eventType string, limit int, offset int) ([]models.Event, error) {
	fmt.Println("success get events tmp")

	return nil, nil
}

func (e *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	//db := database.DBConn
	//
	//db.Create(&event)
	//
	//log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	return event, nil
}

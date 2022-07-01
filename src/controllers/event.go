package controllers

import (
	"log"

	"github.com/Sortren/event-log/src/database"
	"github.com/Sortren/event-log/src/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func EventController(api fiber.Router) {
	eventGroup := api.Group("/event")

	eventGroup.Get("/", GetEvents)
	eventGroup.Post("/", CreateEvent)
}

func GetEvents(c *fiber.Ctx) error {
	return c.SendString("Getting an event")
}

func CreateEvent(c *fiber.Ctx) error {
	db := database.DBConn

	event := new(models.Event)

	if err := c.BodyParser(event); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(event); err != nil {
		return err
	}

	db.Create(&event)

	log.Printf("Event[%s] (%s) added to the database", event.Type, event.Description)

	return c.JSON(event)
}

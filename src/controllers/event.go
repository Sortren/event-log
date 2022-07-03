package controllers

import (
	"github.com/Sortren/event-log/src/models"
	"github.com/Sortren/event-log/src/services"
	"github.com/gofiber/fiber/v2"
)

func RestEventController(api fiber.Router) {
	eventGroup := api.Group("/event")

	eventGroup.Post("/", func(c *fiber.Ctx) error {
		event := new(models.Event)

		if err := c.BodyParser(event); err != nil {
			return fiber.ErrBadRequest
		}

		event, err := services.CreateEvent(event)

		if err != nil {
			return fiber.ErrBadRequest
		}

		return c.JSON(event)
	})

	eventGroup.Get("/", func(c *fiber.Ctx) error {
		if string(c.Request().URI().QueryString()) == "" {
			return fiber.ErrBadRequest
		}

		filters := map[string]string{
			"start": c.Query("start"),
			"end":   c.Query("end"),
			"type":  c.Query("type"),
		}

		events, err := services.GetEvents(filters)

		if err != nil {
			return err
		}

		return c.JSON(events)
	})
}

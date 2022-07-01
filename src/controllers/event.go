package controllers

import "github.com/gofiber/fiber/v2"

func EventController(api fiber.Router) {
	eventGroup := api.Group("/event")

	eventGroup.Get("/", GetEvents)
	eventGroup.Post("/", CreateEvent)
}

func GetEvents(c *fiber.Ctx) error {
	return c.SendString("Getting an event")
}

func CreateEvent(c *fiber.Ctx) error {
	return c.SendString("Creating an event")
}

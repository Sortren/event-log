package controllers

import (
	_ "github.com/Sortren/event-log/src/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Event Logger Service API
// @version 1.0
// @description HTTP Service that allows saving and getting logs to/from database"
// @contact.name Sortren
// @contact.email sortren.dev@gmail.com
// @host event-log:3000
// @BasePath /api/v1
func RegisterRestControllers(app *fiber.App) {
	api_v1 := app.Group("/api/v1")

	events := api_v1.Group("/events")
	restEventController := &RestEventController{}
	{
		events.Post("/", restEventController.CreateEvent)
		events.Get("/", restEventController.GetEvents)
	}

	docs := api_v1.Group("/docs")
	{
		docs.Get("/*", swagger.HandlerDefault)
	}

}

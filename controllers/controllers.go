package controllers

import (
	"fmt"

	"github.com/Sortren/event-log/database"
	_ "github.com/Sortren/event-log/docs"
	"github.com/Sortren/event-log/events"
	"github.com/Sortren/event-log/services"
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

	eventsGroup := api_v1.Group("/events")
	db := database.DBConn
	eventRepo := events.NewRepository(db)
	eventService := services.NewEventService(*eventRepo)
	// start=2022-07-06T18:29:00.000Z&end=2022-07-06T18:33:00.000Z&limit=0&offset=0
	e, err := eventService.GetEvents("2022-07-06T18:29:00.000Z", "2022-07-06T18:33:00.000Z", "", 0, 0)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(e)

	restEventController := NewRestEventController(eventService)
	{
		eventsGroup.Post("/", restEventController.CreateEvent)
		eventsGroup.Get("/", restEventController.GetEvents)
	}

	docs := api_v1.Group("/docs")
	{
		docs.Get("/*", swagger.HandlerDefault)
	}
}

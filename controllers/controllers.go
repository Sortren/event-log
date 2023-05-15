package controllers

import (
	"github.com/Sortren/event-log/database"
	_ "github.com/Sortren/event-log/docs"
	"github.com/Sortren/event-log/pkg/config"
	"github.com/Sortren/event-log/pkg/persistence"
	"github.com/Sortren/event-log/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

// RegisterRestControllers godoc
// @title Event Logger Service API
// @version 1.0
// @description HTTP Service that allows saving and getting logs to/from database"
// @contact.name Sortren
// @contact.email sortren.dev@gmail.com
// @host event-log:3000
// @BasePath /api/v1
func RegisterRestControllers(app *fiber.App) {
	postgresConfig := config.NewPostgres()

	db, err := database.Connect(postgresConfig)
	if err != nil {
		log.Fatalf("can't get db connection, err: %v", err)
	}

	eventRepo := persistence.NewEvent(db)
	eventService := services.NewEventService(eventRepo)

	v1 := app.Group("/api/v1")
	eventsGroup := v1.Group("/events")

	restEventController := NewRestEventController(eventService)
	{
		eventsGroup.Post("/", restEventController.CreateEvent)
		eventsGroup.Get("/", restEventController.GetEvents)
	}

	docs := v1.Group("/docs")
	{
		docs.Get("/*", swagger.HandlerDefault)
	}
}

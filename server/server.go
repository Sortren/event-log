package server

import (
	"github.com/Sortren/event-log/database"
	"github.com/Sortren/event-log/pkg/config"
	"github.com/Sortren/event-log/pkg/persistence"
	"github.com/Sortren/event-log/services"
	"github.com/gofiber/swagger"
	"log"
	"os"
	"os/signal"

	"github.com/Sortren/event-log/controllers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func (s *Server) Setup() {
	if s.App == nil {
		log.Fatal("can't setup the server")
	}

	postgresConfig := config.NewPostgres()

	db, err := database.Connect(postgresConfig)
	if err != nil {
		log.Fatalf("can't get db connection, err: %v", err)
	}

	eventRepo := persistence.NewEvent(db)
	eventService := services.NewEventService(eventRepo)

	restControllers := []controllers.RestController{
		controllers.NewRestEventController(eventService),
	}

	v1 := s.App.Group("/api/v1")
	for _, controller := range restControllers {
		controller.RegisterRoutes(v1)
	}

	docs := v1.Group("/docs")
	docs.Get("/*", swagger.HandlerDefault)
}

func (s *Server) Start() <-chan os.Signal {
	s.Setup()

	exitSignal := make(chan os.Signal, 1)

	signal.Notify(exitSignal, os.Interrupt)

	go func() {
		if err := s.App.Listen(os.Getenv("API_SERVER_URL")); err != nil {
			log.Print("Provided wrong API Server URL")
			log.Fatal(err)
		}
	}()

	return exitSignal
}

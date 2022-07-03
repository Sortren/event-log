package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/Sortren/event-log/src/controllers"
	"github.com/Sortren/event-log/src/database"
	"github.com/gofiber/fiber/v2"
)

func RegisterRestControllers(app *fiber.App) {
	api_v1 := app.Group("/api/v1")

	controllers.RestEventController(api_v1)
}

type Server struct {
	App *fiber.App
}

func (s *Server) Setup() {
	// All middlewares and controllers will be registered here

	if s.App == nil {
		log.Fatal("Can't setup the server")
	}
	database.InitDatabaseConn()

	RegisterRestControllers(s.App)
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

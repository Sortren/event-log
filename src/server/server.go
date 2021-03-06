package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/Sortren/event-log/src/controllers"
	"github.com/Sortren/event-log/src/database"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func (s *Server) Setup() {
	// All middlewares and controllers will be registered here

	if s.App == nil {
		log.Fatal("Can't setup the server")
	}
	database.InitDatabaseConn()

	controllers.RegisterRestControllers(s.App)
}

func (s *Server) Start() <-chan os.Signal {
	s.Setup()
	database.MakeAutoMigrations()

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

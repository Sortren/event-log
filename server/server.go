package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/Sortren/event-log/controllers"
	"github.com/Sortren/event-log/database"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func (s *Server) Setup() {
	if s.App == nil {
		log.Fatal("can't setup the server")
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("can't initialize database, err: %v", err)
	}

	controllers.RegisterRestControllers(s.App)
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

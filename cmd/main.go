package main

import (
	"log"
	"os"

	"github.com/Sortren/event-log/pkg/config"
	"github.com/Sortren/event-log/server"
)

func main() {
	srv := &server.Server{
		App: config.InitializeFiberApp(),
	}

	signal := srv.Start()

	<-signal

	log.Print("Shutting down the server")
	os.Exit(0)
}

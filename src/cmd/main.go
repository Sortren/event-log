package main

import (
	"log"
	"os"

	"github.com/Sortren/event-log/src/config"
	"github.com/Sortren/event-log/src/server"
)

func main() {
	server := &server.Server{
		App: config.InitializeFiberApp(),
	}

	signal := server.Start()

	<-signal

	log.Print("Shutting down the server")
	os.Exit(0)
}

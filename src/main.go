package main

import (
	"github.com/Sortren/event-log/src/controllers"
	"github.com/Sortren/event-log/src/database"
	"github.com/gofiber/fiber/v2"
)

func registerControllers(api fiber.Router) {
	controllers.EventController(api)
}

func main() {
	database.InitDatabaseConn()

	app := fiber.New()
	api_v1 := app.Group("/api/v1")

	registerControllers(api_v1)

	app.Listen(":3000")
}

package controllers

import (
	_ "github.com/Sortren/event-log/docs"
	"github.com/gofiber/fiber/v2"
)

type RestController interface {
	RegisterRoutes(router fiber.Router)
}

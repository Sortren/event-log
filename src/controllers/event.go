package controllers

import (
	_ "net/http/httputil"

	_ "github.com/Sortren/event-log/src/docs"
	"github.com/Sortren/event-log/src/models"
	"github.com/Sortren/event-log/src/services"
	"github.com/gofiber/fiber/v2"
)

type RestEventController struct{}

// CreateEvent godoc
// @Summary      Create an event
// @Description  Create an event and save it to the database
// @Tags         event
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Event
// @Failure      400  {object}  fiber.Error
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /events [post]
func (ctr *RestEventController) CreateEvent(c *fiber.Ctx) error {
	event := new(models.Event)

	if err := c.BodyParser(event); err != nil {
		return fiber.ErrBadRequest
	}

	event, err := services.CreateEvent(event)

	if err != nil {
		return fiber.ErrBadRequest
	}

	return c.JSON(event)
}

// GetEvents godoc
// @Summary      Show a list of filtered events
// @Description  Get objects by query filters, at least one filter is required (type || (start && end))
// @Tags         events
// @Produce      json
// @Param        type   query      string  false  "Event type"
// @Param        start  query      string  false  "Start date" Format(date)
// @Param        end   	query      string  false  "End date" Format(date)
// @Success      200  {array}  models.Event
// @Failure      400  {object}  fiber.Error
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /events [get]
func (ctr *RestEventController) GetEvents(c *fiber.Ctx) error {
	if string(c.Request().URI().QueryString()) == "" {
		return fiber.ErrBadRequest
	}

	filters := map[string]string{
		"start": c.Query("start"),
		"end":   c.Query("end"),
		"type":  c.Query("type"),
	}

	events, err := services.GetEvents(filters)

	if err != nil {
		return err
	}

	return c.JSON(events)
}

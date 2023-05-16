package controllers

import (
	"fmt"
	_ "net/http/httputil"
	"time"

	_ "github.com/Sortren/event-log/docs"
	"github.com/Sortren/event-log/models"
	"github.com/Sortren/event-log/services"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type RestEventController struct {
	eventService services.IEventService
}

func NewRestEventController(service services.IEventService) *RestEventController {
	return &RestEventController{eventService: service}
}

func (c *RestEventController) RegisterRoutes(router fiber.Router) {
	eventsGroup := router.Group("/events")

	eventsGroup.Post("/", c.CreateEvent)
	eventsGroup.Get("/", c.GetEvents)
}

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
func (c *RestEventController) CreateEvent(ctx *fiber.Ctx) error {
	event := new(models.Event)

	if err := ctx.BodyParser(event); err != nil {
		msg := &ErrorMessage{
			Message: "can't bind request body",
			Details: fmt.Sprintf("%s", err),
		}

		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(msg)
	}

	if err := validator.New().Struct(event); err != nil {
		msg := &ErrorMessage{
			Message: "can't validate request body",
			Details: fmt.Sprintf("%s", err),
		}

		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(msg)
	}

	event, err := c.eventService.CreateEvent(ctx.Context(), event)

	if err != nil {
		msg := &ErrorMessage{
			Message: "can't create event",
			Details: fmt.Sprintf("%s", err),
		}

		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(msg)
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(event)
}

// GetEvents godoc
// @Summary      Show a list of filtered events
// @Description  Get objects by query filters, at least one filter is required (type || (start && end))
// @Tags         events
// @Produce      json
// @Param        type   query      string  false  "Event type"
// @Param        start  query      string  false  "Start date" Format(date)
// @Param        end   	query      string  false  "End date" Format(date)
// @Param        limit  query      int     false  "Limit"
// @Param        offset query      int     false  "Offset"
// @Success      200  {array}  models.Event
// @Failure      400  {object}  fiber.Error
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /events [get]
func (c *RestEventController) GetEvents(ctx *fiber.Ctx) error {
	type EventParams struct {
		Type   string    `query:"type"`
		Start  time.Time `query:"start"`
		End    time.Time `query:"end"`
		Limit  int       `query:"limit,required" validate:"gt=0"`
		Offset int       `query:"offset,required" validate:"gte=0"`
	}

	filters := &EventParams{}

	if err := ctx.QueryParser(filters); err != nil {
		msg := &ErrorMessage{
			Message: "required fields are not provided in the queryparams",
			Details: fmt.Sprintf("%s", err),
		}

		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(msg)
	}

	if err := validator.New().Struct(filters); err != nil {
		msg := &ErrorMessage{
			Message: "can't validate queryparams, limit and offset can't be negative",
			Details: fmt.Sprintf("%s", err),
		}

		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(msg)
	}

	events, err := c.eventService.GetEvents(
		ctx.Context(),
		filters.Start,
		filters.End,
		filters.Type,
		filters.Limit,
		filters.Offset,
	)

	if len(events) == 0 {
		msg := ErrorMessage{Message: "events with provided params not found"}
		return ctx.
			Status(fiber.StatusNotFound).
			JSON(msg)
	}

	if err != nil {
		msg := ErrorMessage{Message: "can't get events"}
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(msg)
	}

	return ctx.JSON(events)
}

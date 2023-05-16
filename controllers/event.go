package controllers

import (
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
		return fiber.ErrBadRequest
	}

	validate := validator.New()
	if err := validate.Struct(event); err != nil {
		return err
	}

	event, err := c.eventService.CreateEvent(event)

	if err != nil {
		return err
	}

	return ctx.JSON(event)
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
		Limit  int       `query:"limit,required"`
		Offset int       `query:"offset,required"`
	}

	filters := &EventParams{}

	var validationErrorMessage *ErrorMessage

	if err := ctx.QueryParser(filters); err != nil {
		validationErrorMessage = &ErrorMessage{Message: "Required fields are not provided in the queryparams"}
	}

	if filters.Limit < 0 || filters.Offset < 0 {
		validationErrorMessage = &ErrorMessage{Message: "Limit and Offset can't be negative"}
	}

	if validationErrorMessage != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(validationErrorMessage)
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
		msg := ErrorMessage{Message: "Events with provided params not found"}
		return ctx.
			Status(fiber.StatusNotFound).
			JSON(msg)
	}

	if err != nil {
		msg := ErrorMessage{Message: "Can't get events"}
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(msg)
	}

	return ctx.JSON(events)
}

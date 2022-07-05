package controllers

import (
	_ "net/http/httputil"

	_ "github.com/Sortren/event-log/src/docs"
	"github.com/Sortren/event-log/src/models"
	"github.com/Sortren/event-log/src/services"
	"github.com/Sortren/event-log/src/utils"
	"github.com/gofiber/fiber/v2"
)

type RestEventController struct {
	services.IEventService
}

func NewRestEventController(service services.IEventService) *RestEventController {
	return &RestEventController{IEventService: service}
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
func (ctr *RestEventController) CreateEvent(c *fiber.Ctx) error {
	event := new(models.Event)

	if err := c.BodyParser(event); err != nil {
		return fiber.ErrBadRequest
	}

	event, err := ctr.IEventService.CreateEvent(event)

	if err != nil {
		return err
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
// @Param        limit  query      int     false  "Limit"
// @Param        offset query      int     false  "Offset"
// @Success      200  {array}  models.Event
// @Failure      400  {object}  fiber.Error
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /events [get]
func (ctr *RestEventController) GetEvents(c *fiber.Ctx) error {
	type EventParams struct {
		Type   string `query:"type"`
		Start  string `query:"start"`
		End    string `query:"end"`
		Limit  int    `query:"limit,required"`
		Offset int    `query:"offset,required"`
	}

	filters := &EventParams{}
	if err := c.QueryParser(filters); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Required fields are not provided in the queryparams")
	}

	if utils.IsFilterPresent(filters.Start) != utils.IsFilterPresent(filters.End) {
		return fiber.NewError(fiber.StatusBadRequest, "Can't provide start without end and end without start")
	}

	events, err := ctr.IEventService.GetEvents(filters.Start, filters.End, filters.Type, filters.Limit, filters.Offset)

	if err != nil {
		return err
	}

	return c.JSON(events)
}

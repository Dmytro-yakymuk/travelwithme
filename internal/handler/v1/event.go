package v1

import (
	"net/http"
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initEventRoutes(api *gin.RouterGroup) {
	events := api.Group("/events")
	{
		events.GET("", h.getEvents)
		events.GET("/:id", h.getEventByID)

		eventsPermission := events.Group("", h.userIdentity, h.admin)
		{
			eventsPermission.POST("", h.createEvent)
			eventsPermission.PUT("/:id", h.updateEvent)
			eventsPermission.DELETE("/:id", h.deleteEvent)
		}

	}
}

func (h *Handler) getEvents(c *gin.Context) {

	events, err := h.services.Event.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newEvents := make([]models.EventOutput, 0, len(events))
	for _, v := range events {

		toursEvent, err := h.services.Tour.GetByEventID(v.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		event := models.EventOutput{
			Id:         v.Id,
			Text:       v.Text,
			CountTours: len(toursEvent),
		}
		newEvents = append(newEvents, event)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"events": newEvents,
	})
}

func (h *Handler) getEventByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	event, err := h.services.Event.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *Handler) createEvent(c *gin.Context) {
	var input models.Event
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Event.Create(input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Подію '" + input.Text + "' створено успішно!",
		"status":  true,
	})
}

func (h *Handler) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	var input models.UpdateEventInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Event.Update(id, input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Подію '" + *input.Text + "' оновлено успішно!",
		"status":  true,
	})
}

func (h *Handler) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	if err := h.services.Event.Delete(id); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Подію видалено успішно!",
		"status":  true,
	})
}

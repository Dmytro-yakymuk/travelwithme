package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initTripRoutes(api *gin.RouterGroup) {
	trips := api.Group("/trips", h.userIdentity)
	{
		trips.GET("", h.getAllTrips)
		trips.GET("/:id", h.getOneByID)

		tripsPermission := trips.Group("", h.owner)
		{
			tripsPermission.POST("", h.createTrip)
			tripsPermission.PUT("/:id", h.updateTrip)
			tripsPermission.DELETE("/:id", h.deleteTrip)
		}
	}
}

func (h *Handler) getAllTrips(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	trips, err := h.services.Trip.GetAllTripsByUserId(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, trips)
}

func (h *Handler) getOneByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	trip, err := h.services.Trip.GetOneByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, trip)
}

func (h *Handler) createTrip(c *gin.Context) {
	var input models.CreateTripInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tour, err := h.services.Tour.GetById(input.TourId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	trip := models.Trip{
		Start:  time.Unix(input.Start, 0),
		End:    time.Unix(input.End, 0),
		Price:  input.Price,
		Count:  input.Count,
		TourId: input.TourId,
	}

	if err := h.services.Trip.Create(&trip); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Поїздку для туру '" + tour.Title + "' додано!",
		"status":  true,
	})
}

func (h *Handler) updateTrip(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	var input models.UpdateTripInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	trip := models.Trip{
		Start:  time.Unix(input.Start, 0),
		End:    time.Unix(input.End, 0),
		Price:  input.Price,
		Count:  input.Count,
		TourId: input.TourId,
	}

	if err := h.services.Trip.Update(id, &trip); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	newSuccessResponse(c, http.StatusCreated,
		"Поїздку оновлено успішно!",
		nil,
	)
}

func (h *Handler) deleteTrip(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	if err := h.services.Trip.Delete(id); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	newSuccessResponse(c, http.StatusOK,
		"Поїздку видалено успішно!",
		nil,
	)
}

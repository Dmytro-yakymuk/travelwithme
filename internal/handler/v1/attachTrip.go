package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAttachTripRoutes(api *gin.RouterGroup) {
	trips := api.Group("/attach_trips", h.userIdentity)
	{
		trips.GET("", h.getAllTripsWhichAttach)
		tripsPermssion := trips.Group("", h.client)
		{
			tripsPermssion.POST("", h.createAttach)
			tripsPermssion.POST("/decr", h.decrAttach)
			tripsPermssion.DELETE("/:id", h.deleteAttach)
		}
	}
}

func (h *Handler) getAllTripsWhichAttach(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	nameMap := fmt.Sprintf("trips_%v", userID)
	tripKeys, err := rdb.HKeys(c, nameMap).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	trips, err := h.services.Trip.GetAllTripsWhichAttach(tripKeys)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tripData, err := rdb.HGetAll(c, nameMap).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var newTrips []models.TripToCart
	for _, trip := range trips {
		for tripID, tripCount := range tripData {
			id, err := strconv.Atoi(tripID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			count, err := strconv.Atoi(tripCount)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			if trip.Id == id {
				trip.Count = count
			}
		}

		trip.Image, err = h.services.Image.GetMainNameImageByTourId(trip.TourId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		newTrips = append(newTrips, trip)
	}

	c.JSON(http.StatusOK, newTrips)
}

type TripForAttach struct {
	Id string
}

func (h *Handler) createAttach(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input TripForAttach
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	nameMap := fmt.Sprintf("trips_%v", userID)
	ok, err := rdb.HExists(c, nameMap, input.Id).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	count, err := h.services.Trip.CheckCountTrip(c, rdb, input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if count >= 1 {
		if ok {
			rdb.HIncrBy(c, nameMap, input.Id, 1)
		} else {

			// 1 - це кількість
			var trip = map[string]interface{}{
				input.Id: 1,
			}

			err = rdb.HSet(c, nameMap, trip).Err()
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Недостатньо кількості вільних поїздок для бронювання!",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Поїздку було успішно додано до корзини!",
		"status":  true,
	})
}

func (h *Handler) decrAttach(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input TripForAttach
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var trip = map[string]interface{}{
		input.Id: 1,
	}

	nameMap := fmt.Sprintf("trips_%v", userID)
	ok, err := rdb.HExists(c, nameMap, input.Id).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	count, err := strconv.Atoi(rdb.HGet(c, nameMap, input.Id).Val())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if count > 1 {
		if ok {
			rdb.HSet(c, nameMap, input.Id, count-1)
		} else {
			err = rdb.HSet(c, nameMap, trip).Err()
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Тур №" + input.Id + " вже має мінімальну кількість 1!",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Кількість туру №" + input.Id + " було зменшено на 1!",
		"status":  true,
	})

}

func (h *Handler) deleteAttach(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	nameMap := fmt.Sprintf("trips_%v", userID)
	ok, err := rdb.HExists(c, nameMap, id).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if ok {
		rdb.HDel(c, nameMap, id)
	} else {
		c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Не можливо видалети поїздки з корзини!",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Поїздки видалені з корзини!",
		"status":  true,
	})
}

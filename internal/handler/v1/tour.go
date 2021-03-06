package v1

import (
	"net/http"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func (h *Handler) initTourRoutes(api *gin.RouterGroup) {
	tours := api.Group("/tours")
	{
		tours.GET("", h.getAll)
		tours.GET("/:slug", h.getBySlug)

		toursPermission := tours.Group("", h.userIdentity, h.ownerAndAdmin)
		{
			toursPermission.POST("", h.create)
			toursPermission.PUT("/:slug", h.update)
			toursPermission.DELETE("/:slug", h.delete)
		}

	}
}

func (h *Handler) getAll(c *gin.Context) {
	urlQuery := c.Request.URL.Query()

	tours, err := h.services.Tour.GetAll(urlQuery)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var newTours []models.Tour
	for _, tour := range tours {

		images, err := h.services.Image.GetAllImageByTourId(tour.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		tour.Images = images

		trips, err := h.services.Trip.GetAllTripByTourId(tour.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		tour.Trips = trips
		newTours = append(newTours, tour)

	}

	c.JSON(http.StatusOK, newTours)
}

func (h *Handler) getBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	images, err := h.services.Image.GetAllImageByTourId(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tour.Images = images

	trips, err := h.services.Trip.GetAllTripByTourId(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tour.Trips = trips

	c.JSON(http.StatusOK, tour)
}

func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Tour

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.UserId = userId

	input.Slug = slug.Make(input.Title)

	if err = h.services.Tour.Create(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"slug": input.Slug,
	})
}

func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userId != tour.UserId {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied. not your record")
		return
	}

	var input models.UpdateTourInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Tour.Update(slug, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userId != tour.UserId {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied. not your record")
		return
	}

	if err := h.services.Tour.Delete(slug); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

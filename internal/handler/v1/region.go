package v1

import (
	"net/http"
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func (h *Handler) initRegionRoutes(api *gin.RouterGroup) {
	regions := api.Group("/regions")
	{
		regions.GET("", h.getAllRegions)
		regions.GET("/id/:id", h.getRegionById)
		regions.GET("/slug/:slug", h.getRegionBySlug)

		regionsPermission := regions.Group("", h.userIdentity, h.admin)
		{
			regionsPermission.POST("", h.createRegion)
			regionsPermission.PUT("/:slug", h.updateRegion)
			regionsPermission.DELETE("/:slug", h.deleteRegion)
		}

	}
}

func (h *Handler) getAllRegions(c *gin.Context) {
	regions, err := h.services.Region.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newRegions := make([]models.RegionOutput, 0, len(regions))
	for _, v := range regions {

		tours, err := h.services.Tour.GetByRegionID(v.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		category := models.RegionOutput{
			Id:         v.Id,
			Name:       v.Name,
			Slug:       v.Slug,
			CountTours: len(tours),
		}
		newRegions = append(newRegions, category)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"regions": newRegions,
	})
}

func (h *Handler) getRegionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	region, err := h.services.Region.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, region)
}

func (h *Handler) getRegionBySlug(c *gin.Context) {
	slug := c.Param("slug")
	region, err := h.services.Region.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, region)
}

func (h *Handler) createRegion(c *gin.Context) {
	var input models.Region
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Slug = slug.Make(input.Name)

	if err := h.services.Region.Create(input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Регіон '" + input.Name + "' створено успішно!",
		"status":  true,
	})
}

func (h *Handler) updateRegion(c *gin.Context) {
	slugRegion := c.Param("slug")
	if slugRegion == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slugRegion param")
		return
	}

	var input models.UpdateRegionInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Slug = slug.Make(*input.Name)
	if err := h.services.Region.Update(slugRegion, input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Регіон '" + *input.Name + "' оновлено успішно!",
		"status":  true,
	})
}

func (h *Handler) deleteRegion(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slug param")
		return
	}

	if err := h.services.Region.Delete(slug); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Регіон '" + slug + "' видалено успішно!",
		"status":  true,
	})
}

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initRegionRoutes(api *gin.RouterGroup) {
	regions := api.Group("/regions")
	{
		regions.GET("", h.getAllRegions)
		//regions.GET("/:slug", h.getBySlug)

		regionsPermission := regions.Group("", h.userIdentity, h.owner)
		{
			regionsPermission.POST("", h.create)
			regionsPermission.PUT("/:slug", h.update)
			regionsPermission.DELETE("/:slug", h.delete)
		}

	}
}

func (h *Handler) getAllRegions(c *gin.Context) {
	regions, err := h.services.Region.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, regions)
}

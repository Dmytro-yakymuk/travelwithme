package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initCategoryRoutes(api *gin.RouterGroup) {
	categories := api.Group("/categories")
	{
		categories.GET("", h.getAllCategories)
		//categories.GET("/:slug", h.getBySlug)

		categoriesPermission := categories.Group("", h.userIdentity, h.owner)
		{
			categoriesPermission.POST("", h.create)
			categoriesPermission.PUT("/:slug", h.update)
			categoriesPermission.DELETE("/:slug", h.delete)
		}

	}
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

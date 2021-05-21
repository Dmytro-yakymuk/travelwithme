package v1

import (
	"net/http"
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func (h *Handler) initCategoryRoutes(api *gin.RouterGroup) {
	categories := api.Group("/categories")
	{
		categories.GET("", h.getAllCategories)
		categories.GET("/id/:id", h.getCategoryById)
		categories.GET("/slug/:slug", h.getCategoryBySlug)

		categoriesPermission := categories.Group("", h.userIdentity, h.admin)
		{
			categoriesPermission.POST("", h.createCategory)
			categoriesPermission.PUT("/:slug", h.updateCategory)
			categoriesPermission.DELETE("/:slug", h.deleteCategory)
		}

	}
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newCategories := make([]models.CategoryOutput, 0, len(categories))
	for _, v := range categories {

		tours, err := h.services.Tour.GetByCategoryID(v.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		category := models.CategoryOutput{
			Id:         v.Id,
			Name:       v.Name,
			Slug:       v.Slug,
			Icon:       v.Icon,
			CountTours: len(tours),
		}
		newCategories = append(newCategories, category)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"categories": newCategories,
	})
}

func (h *Handler) getCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	category, err := h.services.Category.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) getCategoryBySlug(c *gin.Context) {
	slug := c.Param("slug")
	category, err := h.services.Category.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) createCategory(c *gin.Context) {
	var input models.Category
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Slug = slug.Make(input.Name)

	if err := h.services.Category.Create(input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Категорію '" + input.Name + "' створено успішно!",
		"status":  true,
	})
}

func (h *Handler) updateCategory(c *gin.Context) {
	slugCategory := c.Param("slug")
	if slugCategory == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slugCategory param")
		return
	}

	var input models.UpdateCategoryInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Slug = slug.Make(*input.Name)
	if err := h.services.Category.Update(slugCategory, input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Категорію '" + *input.Name + "' оновлено успішно!",
		"status":  true,
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty slug param")
		return
	}

	if err := h.services.Category.Delete(slug); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Категорію '" + slug + "' видалено успішно!",
		"status":  true,
	})
}

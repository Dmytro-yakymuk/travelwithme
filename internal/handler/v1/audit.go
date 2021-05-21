package v1

import (
	"net/http"
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuditRoutes(api *gin.RouterGroup) {
	audits := api.Group("/audits")
	{
		audits.GET("/:tour_slug", h.getAllAuditsByTourId)

		auditsPermission := audits.Group("", h.userIdentity, h.admin)
		{
			auditsPermission.POST("", h.createAudit)
			auditsPermission.DELETE("/:id", h.deleteAudit)
		}

	}
}

func (h *Handler) getAllAuditsByTourId(c *gin.Context) {
	slug := c.Param("tour_slug")

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	audits, err := h.services.Audit.GetAllAuditsByTourId(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, audits)
}

func (h *Handler) createAudit(c *gin.Context) {

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole != "admin" {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Створювати вказівку може лише адмін",
			"result":  false,
		})
	}

	var input models.Audit

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tour, err := h.services.Tour.GetById(input.TourId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	audit, err := h.services.Audit.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Вказівку для туру '" + tour.Title + "' вказані!",
		"status":  true,
		"result":  audit,
	})
}

func (h *Handler) deleteAudit(c *gin.Context) {
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	if userRole != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied. not your record")
		return
	}

	if err := h.services.Audit.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Вказівку видалено успішно!",
		"status":  true,
		"result":  id,
	})
}

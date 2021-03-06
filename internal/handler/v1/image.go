package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initImageRoutes(api *gin.RouterGroup) {
	images := api.Group("/images")
	{
		// images.GET("", h.getAllImages)
		// images.GET("/:slug", h.getBySlug)

		imagesPermission := images.Group("", h.userIdentity, h.ownerAndAdmin)
		{
			// imagesPermission.POST("/", h.create)
			imagesPermission.PUT("/:slug", h.updateImage)
			imagesPermission.DELETE("/:id", h.deleteImage)
		}

	}
}

func (h *Handler) updateImage(c *gin.Context) {
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

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error read MultipartForm")
		return
	}

	for i := 0; i < len(form.File); i++ {
		file := form.File[fmt.Sprintf("file[%v]", i)][0]
		imageName := fmt.Sprintf("%v_%v", time.Now().Unix(), file.Filename)

		h.services.Image.Create(tour.Id, imageName)

		err := c.SaveUploadedFile(file, "static/img/tours/"+imageName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error convert or empty id param")
		return
	}

	if err := h.services.Image.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

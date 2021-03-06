package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	// users := api.Group("/users")
	// {
	// 	users.GET("", h.getAllUsers)
	// users.GET("/:slug", h.getBySlug)

	// usersPermission := users.Group("", h.userIdentity, h.ownerAndAdmin)
	// {
	// 	usersPermission.POST("/", h.create)
	// 	usersPermission.PUT("/:slug", h.updateImage)
	// 	usersPermission.DELETE("/:id", h.deleteImage)
	// }

	// }
}

// func (h *Handler) getAllUsers(c *gin.Context) {
// 	urlQuery := c.Request.URL.Query()
// 	users, err := h.services.User.GetAll(urlQuery)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, users)
// }

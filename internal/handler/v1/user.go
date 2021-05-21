package v1

import (
	"net/http"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		// users.GET("", h.getAllUsers)
		users.GET("/byToken", h.userIdentity, h.getOneUserByToken)

		usersPermission := users.Group("", h.userIdentity)
		{
			//sersPermission.POST("/", h.create)
			usersPermission.PUT("/:id", h.updateUser)
			//usersPermission.DELETE("/:id", h.deleteImage)
		}

	}
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

func (h *Handler) getOneUserByToken(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := h.services.User.GetOneByID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetOneByID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !h.services.Authorization.IsEqualPasswordHash(user.Password, input.OldPassword) {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied. not equal password")
		return
	}

	userUpdate := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: h.services.Authorization.GeneratePasswordHash(*input.NewPassword),
		Phone:    *input.Phone,
	}

	if err := h.services.User.Update(userId, userUpdate); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Юзер " + userUpdate.Name + " обновлений!",
		"status":  true,
	})

}

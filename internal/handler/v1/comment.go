package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initCommentRoutes(api *gin.RouterGroup) {
	comments := api.Group("/comments")
	{

		comments.GET("", h.getComments)
		comments.GET("/:id", h.getOneComment)

		commentsPermission := comments.Group("", h.userIdentity, h.client)
		{
			commentsPermission.POST("", h.createComment)
			commentsPermission.PUT("/:id", h.updateComment)
			commentsPermission.DELETE("/:id", h.deleteComment)
		}

	}
}

func (h *Handler) getComments(c *gin.Context) {
	urlQuery := c.Request.URL.Query()

	if urlQuery["user_id"] != nil && urlQuery["user_id"][0] != "0" {

		userID, err := strconv.Atoi(urlQuery["user_id"][0])
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
			return
		}

		user, err := h.services.User.GetOneByID(userID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		role, err := h.services.Role.GetOneByID(userID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if role.Name != "client" {
			if role.Name == "owner" {
				userID = user.Id
			} else {
				userID = 0
			}

			tours, err := h.services.Tour.GetByUserID(userID)
			if err != nil {
				newErrorResponse(c, http.StatusOK, err.Error())
				return
			}

			var toursIds []string
			for _, v := range tours {
				toursIds = append(toursIds, strconv.Itoa(v.Id))
			}

			urlQuery["tour_id"] = toursIds
			delete(urlQuery, "user_id")
		}
	}

	comments, err := h.services.Comment.Get(urlQuery)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newComments := make([]map[string]interface{}, 0, len(comments))
	for _, v := range comments {

		tour, err := h.services.Tour.GetById(v.TourId)
		if err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}

		user, err := h.services.User.GetOneByID(v.UserId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		newComment := map[string]interface{}{
			"id":         v.Id,
			"message":    v.Message,
			"star":       v.Star,
			"tour":       tour,
			"user":       user,
			"created_at": v.CreatedAt,
		}

		newComments = append(newComments, newComment)
	}

	c.JSON(http.StatusOK, newComments)
}

func (h *Handler) getOneComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	comment, err := h.services.Comment.GetOneByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (h *Handler) createComment(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	var input models.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserId = userId
	input.CreatedAt = time.Now()
	if err := h.services.Comment.Create(input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Комент створено успішно!",
		"status":  true,
	})
}

func (h *Handler) updateComment(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	var input models.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserId = userId
	if err := h.services.Comment.Update(id, input); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Комент оновлено успішно!",
		"status":  true,
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid convert or empty param")
		return
	}

	if err := h.services.Comment.Delete(id); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Комент видалено успішно!",
		"status":  true,
	})
}

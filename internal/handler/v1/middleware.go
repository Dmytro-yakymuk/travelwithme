package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPars := strings.Split(header, " ")
	if len(headerPars) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userCtx, err := h.services.Authorization.ParseToken(headerPars[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("id", userCtx["id"])
	c.Set("role", userCtx["role"])
}

func (h *Handler) client(c *gin.Context) {

	role, ok := c.Get("role")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id and role not found")
		return
	}

	if role != "client" {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied")
		return
	}
}

func (h *Handler) owner(c *gin.Context) {

	role, ok := c.Get("role")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id and role not found")
		return
	}

	if role != "owner" {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied")
		return
	}
}

func (h *Handler) admin(c *gin.Context) {

	role, ok := c.Get("role")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id and role not found")
		return
	}

	if role != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied")
		return
	}
}

func (h *Handler) ownerAndAdmin(c *gin.Context) {

	role, ok := c.Get("role")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id and role not found")
		return
	}

	if role != "admin" && role != "owner" {
		newErrorResponse(c, http.StatusInternalServerError, "permission denied")
		return
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("id")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

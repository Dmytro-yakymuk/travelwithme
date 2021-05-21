package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type successResponse struct {
	Message string                 `json:"message"`
	Status  bool                   `json:"status"`
	Result  map[string]interface{} `json:"result"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message, false})
}

func newSuccessResponse(c *gin.Context, statusCode int, message string, result map[string]interface{}) {
	logrus.Error(message)
	c.JSON(statusCode, successResponse{message, true, result})
}

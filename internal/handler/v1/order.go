package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initOrderRoutes(api *gin.RouterGroup) {
	orders := api.Group("/orders")
	{
		ordersPermssion := orders.Group("", h.userIdentity, h.client)
		{
			ordersPermssion.POST("", h.createOrder)
		}
	}
}

func (h *Handler) createOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	fmt.Printf("%v\n", reqBody)

	order := models.Order{
		Count:     1,
		Status:    "paid",
		CreatedAt: time.Now(),
		UserId:    userId,
		TripId:    1,
	}
	fmt.Printf("%v\n", order)
}

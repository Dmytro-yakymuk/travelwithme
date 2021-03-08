package v1

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func (h *Handler) initOrderRoutes(api *gin.RouterGroup, rdb *redis.Client) {
	orders := api.Group("/orders", h.userIdentity, h.client)
	{
		orders.GET("", h.getAllOrders(c*gin.Context, rdb))
		orders.POST("", h.createOrder)
	}
}

func (h *Handler) getAllOrders(c *gin.Context, rdb *redis.Client) {

	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// orders, err := h.services.Order.GetAll(userId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// err := rdb.Set(c, "orders", {
	// 	"user_id": 
	// }).Err()
    // if err != nil {
    //     panic(err)
    // }

	c.JSON(http.StatusOK, orders)
}

type CreateOrderInput struct {
	Id int `json:"id"`
}

func (h *Handler) createOrder(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}


	var input CreateOrderInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// orders := map[string]interface{}{
	// 	"id":         input.Id,
	// 	"created_at": time.Now(),
	// }

	err := rdb.Set(c, "orders", {
		"user_id": userId,
		"trip_id": input.id,
		"count": 1,
	}).Err()
    if err != nil {
        panic(err)
    }


	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	//
	// buf := make([]byte, 1024)
	// num, _ := c.Request.Body.Read(buf)
	// reqBody := string(buf[0:num])
	// fmt.Printf("%v\n", reqBody)

	// order := models.Order{
	// 	Count:     1,
	// 	Status:    "paid",
	// 	CreatedAt: time.Now(),
	// 	UserId:    userId,
	// 	TripId:    1,
	// }

	// if err = h.services.Order.Create(order); err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.Status(http.StatusCreated)
}

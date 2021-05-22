package v1

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/pkg/database/redis"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/mitchellh/mapstructure"
)

var rdb = redis.ConnectRedis("localhost:6379", "", 0)

const (
	checkoutUrl      = "https://pay.fondy.eu/api/checkout/url/"
	merchantId       = "1396424"
	merchantPassword = "test"
	currency         = "UAH"
)

type CheckoutRequest struct {
	OrderId           string `json:"order_id"`
	MerchantId        string `json:"merchant_id"`
	OrderDesc         string `json:"order_desc"`
	Signature         string `json:"signature"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ResponseURL       string `json:"response_url,omitempty"`
	ServerCallbackURL string `json:"server_callback_url,omitempty"`
	SenderEmail       string `json:"sender_email,omitempty"`
	Language          string `json:"lang,omitempty"`
	ProductId         string `json:"product_id,omitempty"`
}

type APIRequest struct {
	Request interface{} `json:"request"`
}

type APIResponse struct {
	Response MapResp `json:"response"`
}

type MapResp struct {
	CheckoutUrl string `json:"checkout_url"`
}

type Response struct {
	OrderID     string `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

func (h *Handler) initOrderRoutes(api *gin.RouterGroup) {
	orders := api.Group("/orders")
	{
		orders.POST("/callback", h.callback)

		orderOnerAndAdmin := orders.Group("", h.userIdentity)
		{
			orderOnerAndAdmin.GET("", h.getAllOrdersByUserIdTrip)
			orderOnerAndAdmin.GET("/:id", h.getOneOrderByUserIdTrip)
		}

		orderAdmin := orders.Group("", h.userIdentity, h.admin)
		{
			orderAdmin.PUT("/paid/:id", h.updatePaid)
		}

		orderClient := orders.Group("", h.userIdentity)
		{
			orderClient.POST("", h.client, h.owner, h.createOrder)
			orderClient.POST("/download/:id", h.client, h.downloadOrder)
		}

	}
}

func (h *Handler) getAllOrdersByUserIdTrip(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	role, ok := c.Get("role")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id and role not found")
		return
	}

	orders, err := h.services.Order.GetAllOrdersByUserIdTrip(userID, fmt.Sprintf("%v", role))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) getOneOrderByUserIdTrip(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	order, err := h.services.Order.GetOneOrderByUserIdTrip(userID, orderID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetOneByID(order.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	trips_orders, err := h.services.TripsOrder.GetAllTripsOrderByOrderId(order.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var general_price float64
	for _, v := range trips_orders {
		general_price += v.Price * float64(v.Count)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":            order.Id,
		"paid":          order.Paid,
		"phone":         order.Phone,
		"text":          order.Text,
		"created_at":    order.CreatedAt,
		"user_id":       order.UserId,
		"trips_orders":  trips_orders,
		"general_price": general_price,
		"user_name":     user.Name,
	})
}

type OrderPaidInput struct {
	Paid bool `json:"paid"`
}

func (h *Handler) updatePaid(c *gin.Context) {
	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input OrderPaidInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Order.UpdatePaid(orderID, !input.Paid); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusCreated)
}

func (r *CheckoutRequest) SetSignature(password string) {
	params := structs.Map(r)

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	values := []string{}

	for _, key := range keys {
		value := params[key].(string)
		if value == "" {
			continue
		}

		values = append(values, value)
	}

	r.Signature = generateSignature(values, password)
}

func generateSignature(values []string, password string) string {
	newValues := []string{password}
	newValues = append(newValues, values...)

	signatureString := strings.Join(newValues, "|")

	hash := sha1.New()
	hash.Write([]byte(signatureString))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (h *Handler) callback(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if resp.OrderStatus == "approved" {
		orderID, err := uuid.Parse(resp.OrderID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if err = h.services.Order.UpdatePaid(orderID, true); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) createOrder(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var orderRequest models.OrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var order models.Order
	if orderRequest.Id == "" {
		order.Id = uuid.New()
	} else {
		orderID, err := uuid.Parse(orderRequest.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		order.Id = orderID
	}

	order.Paid = false
	order.Phone = orderRequest.Phone
	order.Text = orderRequest.Text
	order.CreatedAt = time.Now()
	order.UserId = userID

	checkoutReq := &CheckoutRequest{
		OrderId:           order.Id.String(),
		MerchantId:        merchantId,
		OrderDesc:         "order #" + order.Id.String(),
		Amount:            strconv.Itoa(int(orderRequest.TotalPrice * float64(100))),
		Currency:          currency,
		ServerCallbackURL: "https://49da8245b408.ngrok.io/api/v1/orders/callback",
	}

	checkoutReq.SetSignature(merchantPassword)

	request := APIRequest{Request: checkoutReq}
	requestBody, err := json.Marshal(request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := http.Post(checkoutUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	apiResp := APIResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if orderRequest.Id == "" {

		orderID, err := h.services.Order.Create(order)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		nameMap := fmt.Sprintf("trips_%v", userID)
		orderData, err := rdb.HGetAll(c, nameMap).Result()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		for tripID, count := range orderData {
			countInt, err := strconv.Atoi(count)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			tripIDInt, err := strconv.Atoi(tripID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			tripsOrder := models.TripsOrder{
				Count:   countInt,
				TripId:  tripIDInt,
				OrderId: orderID,
			}

			if err := h.services.TripsOrder.Create(tripsOrder); err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			rdb.HDel(c, nameMap, tripID)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Замовлення додано! В адмін панелі доступний перегляд замовлень.",
			"status":  true,
			"result":  apiResp.Response.CheckoutUrl,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Запит на оплату надіслано. Після оплати обновіть сторінку",
			"status":  true,
			"result":  apiResp.Response.CheckoutUrl,
		})
	}
}

func (h *Handler) downloadOrder(c *gin.Context) {

	orderID := c.Param("id")

	nameFile := ""
	if orderID != "" {

		orderID, err := uuid.Parse(orderID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		order, err := h.services.Order.GetOneByID(orderID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		trips, err := h.services.TripsOrder.GetAllTripsOrderByOrderId(order.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		nameFile, err = h.services.Order.DownloadOrder(order, trips)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Замовлення додано! В адмін панелі доступний перегляд замовлень.",
		"status":  true,
		"result":  "http://localhost:8000/assets/pdfs/" + nameFile,
	})
}

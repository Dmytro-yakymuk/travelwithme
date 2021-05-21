package handler

import (
	"net/http"
	"time"

	v1 "github.com/Dmytro-yakymuk/travelwithme/internal/handler/v1"
	"github.com/Dmytro-yakymuk/travelwithme/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Handler ...
type Handler struct {
	authService       service.Authorization
	tourService       service.Tour
	categoryService   service.Category
	regionService     service.Region
	imageService      service.Image
	userService       service.User
	tripService       service.Trip
	tripsOrderService service.TripsOrder
	orderService      service.Order
	eventService      service.Event
	roleService       service.Role
	services          *service.Service
	rdb               *redis.Client
}

// NewHandler ...
func NewHandler(services *service.Service, rdb *redis.Client) *Handler {
	return &Handler{
		services: services,
		rdb:      rdb,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST, OPTIONS, PUT, DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Static file
	router.Static("/assets", "/media/dmytro/Disk_D1/Web/golang/travelwithme/static/img")

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.rdb)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}

package v1

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	services *service.Service
	rdb      *redis.Client
}

func NewHandler(services *service.Service, rdb *redis.Client) *Handler {
	return &Handler{
		services: services,
		rdb:      rdb,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initTourRoutes(v1)
		h.initCategoryRoutes(v1)
		h.initRegionRoutes(v1)
		h.initImageRoutes(v1)
		h.initUserRoutes(v1)
		h.initAttachTripRoutes(v1)
		h.initTripRoutes(v1)
		h.initOrderRoutes(v1)
		h.initCommentRoutes(v1)
		h.initAuditRoutes(v1)
		h.initEventRoutes(v1)

	}
}

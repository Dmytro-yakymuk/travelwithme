package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initTripRoutes(api *gin.RouterGroup) {
	trips := api.Group("/trips")
	{
		tripsPermssion := trips.Group("", h.userIdentity, h.ownerAndAdmin)
		{
			tripsPermssion.POST("/", h.createTrip)
		}
	}
}

func (h *Handler) createTrip(c *gin.Context) {

}

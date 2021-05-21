package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrderTrip(db *gorm.DB, count, tripID int, orderIDStr string) error {
	orderID, _ := uuid.Parse(orderIDStr)
	return db.Create(&models.TripsOrder{Count: count, TripId: tripID, OrderId: orderID}).Error
}

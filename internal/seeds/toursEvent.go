package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateToursEvent(db *gorm.DB, tourId, eventId int) error {
	return db.Create(&models.ToursEvent{TourId: tourId, EventId: eventId}).Error
}

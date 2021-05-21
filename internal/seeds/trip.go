package seeds

import (
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateTrip(db *gorm.DB, start, end time.Time, price float64, count, tour_id int) error {
	return db.Create(&models.Trip{Start: start, End: end, Price: price, Count: count, TourId: tour_id}).Error
}

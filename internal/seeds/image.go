package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateImage(db *gorm.DB, name string, tour_id int) error {
	return db.Create(&models.Image{Name: name, TourId: tour_id}).Error
}

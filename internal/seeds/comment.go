package seeds

import (
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB, message string, star, tourID, userID int, createdAt time.Time) error {
	return db.Create(&models.Comment{Message: message, Star: star, TourId: tourID, UserId: userID, CreatedAt: createdAt}).Error
}

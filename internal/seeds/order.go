package seeds

import (
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, idStr string, paid bool, phone, text string, createdAt time.Time, userID int) error {
	id, _ := uuid.Parse(idStr)
	return db.Create(&models.Order{Id: id, Paid: paid, Phone: phone, Text: text, CreatedAt: createdAt, UserId: userID}).Error
}

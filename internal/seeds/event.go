package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateEvent(db *gorm.DB, text string) error {
	return db.Create(&models.Event{Text: text}).Error
}

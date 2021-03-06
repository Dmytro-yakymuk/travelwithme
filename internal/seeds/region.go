package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateRegion(db *gorm.DB, name, slug string) error {
	return db.Create(&models.Region{Name: name, Slug: slug}).Error
}

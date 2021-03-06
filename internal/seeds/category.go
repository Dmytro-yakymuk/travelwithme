package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name, slug, icon string) error {
	return db.Create(&models.Category{Name: name, Slug: slug, Icon: icon}).Error
}

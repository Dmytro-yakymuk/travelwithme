package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateTour(db *gorm.DB, title, slug, description, status string, category_id, region_id, user_id int) error {
	return db.Create(&models.Tour{Title: title, Slug: slug, Description: description, Status: status, CategoryId: category_id, RegionId: region_id, UserId: user_id}).Error
}

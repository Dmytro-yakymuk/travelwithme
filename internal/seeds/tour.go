package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateTour(db *gorm.DB, title, slug, description, text string, public, activ bool, category_id, region_id, user_id int, x, y float64) error {
	return db.Create(&models.Tour{Title: title, Slug: slug, Description: description, Text: text, Public: public, Activ: activ, CategoryId: category_id, RegionId: region_id, UserId: user_id, X: x, Y: y}).Error
}

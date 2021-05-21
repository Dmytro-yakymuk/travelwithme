package seeds

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB, name string) error {
	return db.Create(&models.Role{Name: name}).Error
}

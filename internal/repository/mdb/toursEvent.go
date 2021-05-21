package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type ToursEventRepository struct {
	db *gorm.DB
}

func NewToursEventRepository(db *gorm.DB) *ToursEventRepository {
	return &ToursEventRepository{db: db}
}

func (r *ToursEventRepository) Create(toursEvent []*models.ToursEvent) error {
	result := r.db.Model(&models.ToursEvent{}).Create(&toursEvent)
	return result.Error
}

func (r *ToursEventRepository) Delete(tourID int) error {
	result := r.db.Where("tour_id = ?", tourID).Delete(&models.ToursEvent{})
	return result.Error
}

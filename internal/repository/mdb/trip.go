package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) GetAllTripByTourId(tour_id int) ([]models.Trip, error) {
	var trips []models.Trip

	result := r.db.Where("tour_id = ?", tour_id).Find(&trips)
	return trips, result.Error
}

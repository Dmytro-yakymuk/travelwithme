package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TripsOrderRepository struct {
	db *gorm.DB
}

func NewTripsOrderRepository(db *gorm.DB) *TripsOrderRepository {
	return &TripsOrderRepository{db: db}
}

func (r *TripsOrderRepository) Create(tripsOrder models.TripsOrder) error {
	result := r.db.Model(&models.TripsOrder{}).Create(&tripsOrder)
	return result.Error
}

func (r *TripsOrderRepository) GetAllTripsOrderByOrderId(orderID uuid.UUID) ([]models.TripsOrderOutput, error) {
	var output []models.TripsOrderOutput
	result := r.db.Model(
		&models.TripsOrder{},
	).Select(
		"trips.id, tours.title, tours.slug, trips_orders.count, trips.start, trips.end, trips.price",
	).Joins(
		"left join trips on trips_orders.trip_id = trips.id",
	).Joins(
		"left join tours on trips.tour_id = tours.id",
	).Where(
		"trips_orders.order_id = ?", orderID,
	).Find(
		&output,
	)
	return output, result.Error
}

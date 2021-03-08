package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAll(id int) ([]models.GetOrderToCart, error) {
	var orders []models.GetOrderToCart

	result := r.db.Table("orders").Joins("left join trips on orders.trip_id = trips.id").Joins("left join tours on trips.tour_id = tours.id").Select("orders.id", "orders.count", "orders.created_at", "trips.price", "tours.title").Where("orders.user_id = ?", id).Take(&orders)
	return orders, result.Error
}

func (r *OrderRepository) GetAllOrderByTourId(tour_id int) ([]models.Order, error) {
	var orders []models.Order

	result := r.db.Where("tour_id = ?", tour_id).Find(&orders)
	return orders, result.Error
}

func (r *OrderRepository) Create(order models.Order) error {
	result := r.db.Model(&models.Order{}).Create(&order)
	return result.Error
}

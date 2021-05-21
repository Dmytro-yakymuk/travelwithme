package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAllOrdersByUserIdTrip(userID int, role string) ([]models.OrderOutput, error) {
	var orders []models.OrderOutput

	result := r.db.Table(
		"orders",
	).Select(
		"orders.id, orders.paid, orders.phone, orders.text, orders.created_at, orders.user_id, users.name as user_name, SUM(trips.price*trips_orders.count) as general_price",
	).Joins(
		"left join users on users.id = orders.user_id",
	).Joins(
		"left join trips_orders on orders.id = trips_orders.order_id",
	).Joins(
		"left join trips on trips_orders.trip_id = trips.id",
	).Joins(
		"left join tours on trips.tour_id = tours.id",
	)

	if role == "admin" {
		ids := []string{}
		r.db.Table("roles").Select("id").Where("name = ? OR name = ?", "admin", "owner").Find(&ids)
		result.Where(
			"tours.user_id IN ?", ids,
		)
	} else if role == "owner" {
		result.Where(
			"tours.user_id = ?", userID,
		)
	} else {
		result.Where(
			"orders.user_id = ?", userID,
		)
	}

	result.Group(
		"orders.id",
	).Find(
		&orders,
	)

	return orders, result.Error
}

func (r *OrderRepository) GetOneByID(orderID uuid.UUID) (models.OrderOutput, error) {
	var orders models.OrderOutput

	result := r.db.Table(
		"orders",
	).Select(
		"orders.id, orders.paid, orders.phone, orders.text, orders.created_at, orders.user_id, users.name as user_name, users.surname as user_surname, users.patronymic as user_patronymic, users.email as user_email, SUM(trips.price*trips_orders.count) as general_price",
	).Joins(
		"left join users on users.id = orders.user_id",
	).Joins(
		"left join trips_orders on orders.id = trips_orders.order_id",
	).Joins(
		"left join trips on trips_orders.trip_id = trips.id",
	).Joins(
		"left join tours on trips.tour_id = tours.id",
	)

	result.Group(
		"orders.id",
	).Find(
		&orders,
	)

	return orders, result.Error
}

func (r *OrderRepository) GetOneOrderByUserIdTrip(userID int, orderID uuid.UUID) (models.Order, error) {
	var order models.Order

	result := r.db.Model(
		&order,
	).Joins(
		"left join trips_orders on orders.id = trips_orders.order_id",
	).Where(
		"trips_orders.order_id = ?", orderID,
	).Find(
		&order,
	)

	return order, result.Error
}

func (r *OrderRepository) UpdatePaid(id uuid.UUID, paid bool) error {
	result := r.db.Model(&models.Order{}).Where("id = ?", id).Update("paid", paid)
	return result.Error
}

func (r *OrderRepository) Create(order models.Order) (uuid.UUID, error) {
	result := r.db.Model(&models.Order{}).Create(&order)
	return order.Id, result.Error
}

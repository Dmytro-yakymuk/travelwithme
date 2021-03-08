package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type OrderService struct {
	orderRepository repository.Order
}

func NewOrderService(orderRepository repository.Order) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) GetAll(id int) ([]models.GetOrderToCart, error) {
	return s.orderRepository.GetAll(id)
}

func (s *OrderService) GetAllOrderByTourId(tour_id int) ([]models.Order, error) {
	return s.orderRepository.GetAllOrderByTourId(tour_id)
}

func (s *OrderService) Create(order models.Order) error {
	return s.orderRepository.Create(order)
}

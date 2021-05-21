package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
	"github.com/google/uuid"
)

type TripsOrderService struct {
	tripsOrderRepository repository.TripsOrder
}

func NewTripsOrderService(tripsOrderRepository repository.TripsOrder) *TripsOrderService {
	return &TripsOrderService{tripsOrderRepository: tripsOrderRepository}
}

func (s *TripsOrderService) Create(tripsOrder models.TripsOrder) error {
	return s.tripsOrderRepository.Create(tripsOrder)
}

func (s *TripsOrderService) GetAllTripsOrderByOrderId(orderID uuid.UUID) ([]models.TripsOrderOutput, error) {
	return s.tripsOrderRepository.GetAllTripsOrderByOrderId(orderID)
}

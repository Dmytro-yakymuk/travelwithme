package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type TripService struct {
	tripRepository repository.Trip
}

func NewTripService(tripRepository repository.Trip) *TripService {
	return &TripService{tripRepository: tripRepository}
}

func (s *TripService) GetAllTripByTourId(tour_id int) ([]models.Trip, error) {
	return s.tripRepository.GetAllTripByTourId(tour_id)
}

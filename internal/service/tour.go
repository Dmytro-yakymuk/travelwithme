package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type TourService struct {
	tourRepository repository.Tour
}

func NewTourService(tourRepository repository.Tour) *TourService {
	return &TourService{tourRepository: tourRepository}
}

func (s *TourService) GetAll(urlQuery map[string][]string) ([]models.Tour, error) {
	return s.tourRepository.GetAll(urlQuery)
}

func (s *TourService) GetBySlug(slug string) (models.Tour, error) {
	return s.tourRepository.GetBySlug(slug)
}

func (s *TourService) Create(tour models.Tour) error {
	return s.tourRepository.Create(tour)
}

func (s *TourService) Update(slug string, tour models.UpdateTourInput) error {
	if err := tour.Validate(); err != nil {
		return err
	}
	return s.tourRepository.Update(slug, tour)
}

func (s *TourService) Delete(slug string) error {
	return s.tourRepository.Delete(slug)
}

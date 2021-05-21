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
func (s *TourService) GetById(id int) (models.Tour, error) {
	return s.tourRepository.GetById(id)
}
func (s *TourService) GetByCategoryID(category_id int) ([]models.Tour, error) {
	return s.tourRepository.GetByCategoryID(category_id)
}
func (s *TourService) GetByRegionID(region_id int) ([]models.Tour, error) {
	return s.tourRepository.GetByRegionID(region_id)
}
func (s *TourService) GetByEventID(event_id int) ([]models.ToursEvent, error) {
	return s.tourRepository.GetByEventID(event_id)
}
func (s *TourService) GetByUserID(user_id int) ([]models.Tour, error) {
	return s.tourRepository.GetByUserID(user_id)
}
func (s *TourService) Create(tour models.Tour) (int, error) {
	return s.tourRepository.Create(tour)
}
func (s *TourService) Update(slug string, tour models.UpdateTourInput) error {
	if err := tour.Validate(); err != nil {
		return err
	}
	return s.tourRepository.Update(slug, tour)
}
func (s *TourService) Public(slug string, public bool) error {
	return s.tourRepository.Public(slug, public)
}
func (s *TourService) Activ(slug string, activ bool) error {
	return s.tourRepository.Activ(slug, activ)
}
func (s *TourService) Delete(slug string) error {
	return s.tourRepository.Delete(slug)
}

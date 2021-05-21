package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type ToursEventService struct {
	toursEventRepository repository.ToursEvent
}

func NewToursEventService(toursEventRepository repository.ToursEvent) *ToursEventService {
	return &ToursEventService{toursEventRepository: toursEventRepository}
}

func (s *ToursEventService) Create(toursEvent []*models.ToursEvent) error {
	return s.toursEventRepository.Create(toursEvent)
}

func (s *ToursEventService) Delete(tourID int) error {
	return s.toursEventRepository.Delete(tourID)
}

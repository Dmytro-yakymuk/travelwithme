package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type EventService struct {
	eventRepository repository.Event
}

func NewEventService(eventRepository repository.Event) *EventService {
	return &EventService{eventRepository: eventRepository}
}
func (s *EventService) GetAllEventByTourId(tourId int) ([]models.Event, error) {
	return s.eventRepository.GetAllEventByTourId(tourId)
}
func (s *EventService) GetAll() ([]models.Event, error) {
	return s.eventRepository.GetAll()
}
func (s *EventService) GetByID(id int) (models.Event, error) {
	return s.eventRepository.GetByID(id)
}
func (s *EventService) Create(event models.Event) error {
	return s.eventRepository.Create(event)
}
func (s *EventService) Update(id int, event models.UpdateEventInput) error {
	return s.eventRepository.Update(id, event)
}
func (s *EventService) Delete(id int) error {
	return s.eventRepository.Delete(id)
}

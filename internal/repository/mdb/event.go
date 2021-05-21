package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetAllEventByTourId(tourId int) ([]models.Event, error) {
	var events []models.Event
	result := r.db.Select(
		"events.id, events.text",
	).Joins(
		"left join tours_events on events.id = tours_events.event_id",
	).Where(
		"tours_events.tour_id = ?", tourId,
	).Find(
		&events,
	)
	return events, result.Error
}

func (r *EventRepository) GetAll() ([]models.Event, error) {
	var events []models.Event
	result := r.db.Model(&events).Find(&events)
	return events, result.Error
}

func (r *EventRepository) GetByID(id int) (models.Event, error) {
	var event models.Event
	result := r.db.Where("id = ?", id).First(&event)
	return event, result.Error
}

func (r *EventRepository) Create(event models.Event) error {
	result := r.db.Model(&models.Event{}).Create(&event)
	return result.Error
}

func (r *EventRepository) Update(id int, event models.UpdateEventInput) error {
	result := r.db.Model(models.Event{}).Where("id = ?", id).Updates(&models.Event{Text: *event.Text})
	return result.Error
}

func (r *EventRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Delete(&models.Event{})
	return result.Error
}

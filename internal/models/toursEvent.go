package models

// ToursEvent ...
type ToursEvent struct {
	TourId  int `gorm:"primaryKey; not null" json:"tour_id" binding:"required"`
	EventId int `gorm:"primaryKey; not null" json:"event_id" binding:"required"`
}

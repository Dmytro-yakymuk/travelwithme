package models

import (
	"time"

	"github.com/google/uuid"
)

// TripsOrder ...
type TripsOrder struct {
	Count int `gorm:"not null" json:"count" binding:"required"`

	TripId  int       `gorm:"primaryKey; not null" json:"trip_id" binding:"required"`
	OrderId uuid.UUID `gorm:"primaryKey; not null" json:"order_id" binding:"required"`
}

type TripsOrderOutput struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Slug  string    `json:"slug"`
	Count int       `json:"count"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Price float64   `json:"price"`
}

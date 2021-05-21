package models

import "time"

type Trip struct {
	Id    int       `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Start time.Time `gorm:"type:date; not null" json:"start" binding:"required"`
	End   time.Time `gorm:"type:date; not null" json:"end" binding:"required"`
	Price float64   `gorm:"not null" json:"price" binding:"required"`
	Count int       `gorm:"not null" json:"count" binding:"required"`

	TourId int `gorm:"not null" json:"tour_id"`

	TripsOrder []TripsOrder `gorm:"ForeignKey:TripId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type CreateTripInput struct {
	Start  int64   `json:"start" binding:"required"`
	End    int64   `json:"end" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Count  int     `json:"count" binding:"required"`
	TourId int     `json:"tour_id" binding:"required"`
}

type UpdateTripInput struct {
	Start  int64   `json:"start" binding:"required"`
	End    int64   `json:"end" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Count  int     `json:"count" binding:"required"`
	TourId int     `json:"tour_id" binding:"required"`
}

type TripToCart struct {
	Id     int       `json:"id"`
	TourId int       `json:"tour_id"`
	Title  string    `json:"title"`
	Image  string    `json:"image"`
	Slug   string    `json:"slug"`
	Count  int       `json:"count"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Price  float64   `json:"price"`
}

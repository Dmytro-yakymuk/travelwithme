package models

import "time"

type Trip struct {
	Id    int       `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Start time.Time `gorm:"type:date; not null" json:"start" binding:"required"`
	End   time.Time `gorm:"type:date; not null" json:"end" binding:"required"`
	Price float64   `gorm:"not null" json:"price" binding:"required"`

	TourId int `gorm:"not null" json:"tour_id"`

	Order []Order `gorm:"ForeignKey:TripId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

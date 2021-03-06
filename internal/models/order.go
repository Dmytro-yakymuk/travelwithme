package models

import "time"

type Order struct {
	Id        int       `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"-"`
	Count     int       `gorm:"not null" json:"count" binding:"required"`
	Status    string    `gorm:"type:varchar(20); not null" json:"status" binding:"required"`
	CreatedAt time.Time `gorm:"type:date; not null" json:"created_at"`
	UserId    int       `gorm:"not null" json:"user_id"`
	TripId    int       `gorm:"not null" json:"trip_id"`
}

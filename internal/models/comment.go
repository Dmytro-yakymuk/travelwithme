package models

import "time"

type Comment struct {
	Id        int       `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Message   string    `gorm:"type:text; not null" json:"message" binding:"required"`
	Star      int       `gorm:"not null" json:"star" binding:"required"`
	TourId    int       `gorm:"not null" json:"tour_id" binding:"required"`
	UserId    int       `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

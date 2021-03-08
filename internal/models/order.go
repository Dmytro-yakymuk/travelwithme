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

type GetOrderToCart struct {
	Id        int       `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Title     string    `gorm:"type:varchar(255); not null" json:"title" binding:"required"`
	Count     int       `gorm:"not null" json:"count" binding:"required"`
	CreatedAt time.Time `gorm:"type:date; not null" json:"created_at"`
	Price     float64   `gorm:"not null" json:"price" binding:"required"`
}

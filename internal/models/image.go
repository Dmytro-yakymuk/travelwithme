package models

type Image struct {
	Id     int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Name   string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	TourId int    `gorm:"not null" json:"tour_id" binding:"required"`
}

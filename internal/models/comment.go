package models

type Comment struct {
	Id      int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"-"`
	Message string `gorm:"type:text; not null" json:"message" binding:"required"`
	TourId  int    `gorm:"not null" json:"tour_id"`
	UserId  int    `gorm:"not null" json:"user_id"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id        uuid.UUID `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Paid      bool      `gorm:"not null" json:"paid"`
	Phone     string    `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
	Text      string    `gorm:"type:text" json:"text"`
	CreatedAt time.Time `gorm:"type:datetime; not null" json:"created_at"`
	UserId    int       `gorm:"not null" json:"user_id"`

	TripsOrder []TripsOrder `gorm:"ForeignKey:OrderId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type OrderRequest struct {
	Id         string  `json:"id"`
	Phone      string  `json:"phone" binding:"required"`
	Text       string  `json:"text"`
	TotalPrice float64 `json:"totalPrice" binding:"required"`
}

type OrderDownload struct {
	Id string `json:"id" binding:"required"`
}

type OrderOutput struct {
	Id             uuid.UUID `json:"id"`
	Paid           bool      `json:"paid"`
	Phone          string    `json:"phone"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
	UserId         int       `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserSurname    string    `json:"user_surname"`
	UserPatronymic string    `json:"user_patronymic"`
	UserEmail      string    `json:"user_email"`
	GeneralPrice   float64   `json:"general_price"`
}

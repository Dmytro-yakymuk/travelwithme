package models

type Audit struct {
	Id      int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Message string `gorm:"type:text; not null" json:"message"`
	TourId  int    `gorm:"not null" json:"tour_id"`
}

// type AuditInput struct {
// 	Message string `json:"message" binding:"required"`
// 	Slug    string `json:"slug" binding:"required"`
// }

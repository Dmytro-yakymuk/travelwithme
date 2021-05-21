package models

type Event struct {
	Id   int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Text string `gorm:"type:text; not null" json:"text" binding:"required"`

	ToursEvent []ToursEvent `gorm:"ForeignKey:EventId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type UpdateEventInput struct {
	Text *string `json:"text" binding:"required"`
}

type EventOutput struct {
	Id         int    `json:"id"`
	Text       string `json:"text"`
	CountTours int    `json:"count_tours"`
}

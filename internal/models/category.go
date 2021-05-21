package models

type Category struct {
	Id   int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Name string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Slug string `gorm:"type:varchar(255); not null; unique" json:"slug"`
	Icon string `gorm:"type:varchar(255); not null" json:"icon" binding:"required"`
	Tour []Tour `gorm:"ForeignKey:CategoryId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type UpdateCategoryInput struct {
	Name *string `json:"name" binding:"required"`
	Slug string  `json:"slug"`
	Icon *string `json:"icon" binding:"required"`
}

type CategoryOutput struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Icon       string `json:"icon"`
	CountTours int    `json:"count_tours"`
}

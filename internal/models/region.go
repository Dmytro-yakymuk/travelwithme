package models

type Region struct {
	Id   int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Name string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Slug string `gorm:"type:varchar(255); not null; unique" json:"slug"`
	Tour []Tour `gorm:"ForeignKey:RegionId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type UpdateRegionInput struct {
	Name *string `json:"name" binding:"required"`
	Slug string  `json:"slug"`
}

type RegionOutput struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	CountTours int    `json:"count_tours"`
}

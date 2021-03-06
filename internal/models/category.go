package models

type Category struct {
	Id   int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Name string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Slug string `gorm:"type:varchar(255); not null; unique" json:"slug" binding:"required"`
	Icon string `gorm:"type:varchar(255); not null" json:"icon" binding:"required"`
	Tour []Tour `gorm:"ForeignKey:CategoryId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

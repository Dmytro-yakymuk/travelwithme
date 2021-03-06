package models

import "errors"

type Tour struct {
	Id          int    `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Title       string `gorm:"type:varchar(255); not null" json:"title" binding:"required"`
	Slug        string `gorm:"type:varchar(255); not null; unique" json:"slug"`
	Description string `gorm:"type:text; not null" json:"description" binding:"required"`
	Status      string `gorm:"type:varchar(20); not null" json:"status"`

	CategoryId int `gorm:"not null" json:"category_id" binding:"required"`
	RegionId   int `gorm:"not null" json:"region_id" binding:"required"`
	UserId     int `gorm:"not null" json:"user_id"`

	Images   []Image   `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images"`
	Comments []Comment `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Trips    []Trip    `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"trips"`
}

// указетели застосовуємо тому щоб поля які не вели будуть nil в не залежності від типу даних, це допоможе при фільтрації полів в repository.Update
type UpdateTourInput struct {
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
	CategoryId  *int    `json:"category_id"`
	RegionId    *int    `json:"region_id"`
}

func (i UpdateTourInput) Validate() error {
	if i.Title == nil && i.Slug == nil && i.Description == nil && i.CategoryId == nil && i.RegionId == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

package models

import "errors"

type Tour struct {
	Id          int     `gorm:"primaryKey; autoIncrement:true; not null; unique" json:"id"`
	Title       string  `gorm:"type:varchar(255); not null" json:"title" binding:"required"`
	Slug        string  `gorm:"type:varchar(255); not null; unique" json:"slug"`
	Description string  `gorm:"type:text; not null" json:"description" binding:"required"`
	Text        string  `gorm:"type:text; not null" json:"text" binding:"required"`
	Public      bool    `gorm:"not null" json:"public"`
	Activ       bool    `gorm:"not null" json:"activ"`
	X           float64 `gorm:"not null;" json:"x" binding:"required"`
	Y           float64 `gorm:"not null;" json:"y" binding:"required"`

	CategoryId int `gorm:"not null" json:"category_id" binding:"required"`
	RegionId   int `gorm:"not null" json:"region_id" binding:"required"`
	UserId     int `gorm:"not null" json:"user_id"`

	Images   []Image   `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images"`
	Comments []Comment `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	Trips    []Trip    `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"trips"`
	Audit    Audit     `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"audit"`

	ToursEvent []ToursEvent `gorm:"ForeignKey:TourId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

type CreateTourInput struct {
	Id             int     `json:"id"`
	Title          string  `json:"title" binding:"required"`
	Slug           string  `json:"slug"`
	Description    string  `json:"description" binding:"required"`
	Text           string  `json:"text" binding:"required"`
	Public         bool    `json:"public"`
	Activ          bool    `json:"activ"`
	X              float64 `json:"x" binding:"required"`
	Y              float64 `json:"y" binding:"required"`
	CategoryId     int     `json:"category_id" binding:"required"`
	RegionId       int     `json:"region_id" binding:"required"`
	UserId         int     `json:"user_id"`
	SelectedEvents []int   `json:"selected_events"`
}

// указетели застосовуємо тому щоб поля які не вели будуть nil в не залежності від типу даних, це допоможе при фільтрації полів в repository.Update
type UpdateTourInput struct {
	Title          *string  `json:"title"`
	Slug           *string  `json:"slug"`
	Description    *string  `json:"description"`
	Text           *string  `json:"text"`
	X              *float64 `json:"x"`
	Y              *float64 `json:"y"`
	CategoryId     *int     `json:"category_id"`
	RegionId       *int     `json:"region_id"`
	SelectedEvents []int    `json:"selected_events"`
}

func (i UpdateTourInput) Validate() error {
	if i.Title == nil && i.Slug == nil && i.Description == nil && i.Text == nil && i.X == nil && i.Y == nil && i.CategoryId == nil && i.RegionId == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

func (r *ImageRepository) GetAllImageByTourId(tour_id int) ([]models.Image, error) {
	var images []models.Image

	result := r.db.Where("tour_id = ?", tour_id).Find(&images)
	return images, result.Error
}

func (r *ImageRepository) GetMainNameImageByTourId(tour_id int) (string, error) {
	var image models.Image
	result := r.db.Where("tour_id = ?", tour_id).First(&image)
	return image.Name, result.Error
}

func (r *ImageRepository) Create(tour_id int, name string) error {
	result := r.db.Model(&models.Image{}).Create(map[string]interface{}{
		"name": name, "tour_id": tour_id,
	})
	return result.Error
}

func (r *ImageRepository) Delete(id int) (models.Image, error) {
	image := models.Image{}
	err := r.db.Select("name").First(&image, id)
	if err.Error != nil {
		return image, err.Error
	}
	result := r.db.Where("id = ?", id).Delete(image)
	return image, result.Error
}

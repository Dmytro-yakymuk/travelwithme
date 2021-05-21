package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type RegionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) *RegionRepository {
	return &RegionRepository{db: db}
}

func (r *RegionRepository) GetAll() ([]models.Region, error) {
	var regions []models.Region

	result := r.db.Find(&regions)
	return regions, result.Error
}

func (r *RegionRepository) GetById(id int) (models.Region, error) {
	var region models.Region
	result := r.db.Find(&region, id)
	return region, result.Error
}

func (r *RegionRepository) GetBySlug(slug string) (models.Region, error) {
	var region models.Region
	result := r.db.Where("slug = ?", slug).First(&region)
	return region, result.Error
}

func (r *RegionRepository) Create(region models.Region) error {
	result := r.db.Model(&models.Region{}).Create(&region)
	return result.Error
}

func (r *RegionRepository) Update(slug string, region models.UpdateRegionInput) error {
	result := r.db.Model(models.Region{}).Where("slug = ?", slug).Updates(&models.Region{Name: *region.Name, Slug: region.Slug})
	return result.Error
}

func (r *RegionRepository) Delete(slug string) error {
	result := r.db.Where("slug = ?", slug).Delete(&models.Region{})
	return result.Error
}

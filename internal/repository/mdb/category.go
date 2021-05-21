package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	result := r.db.Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) GetById(id int) (models.Category, error) {
	var category models.Category
	result := r.db.Find(&category, id)
	return category, result.Error
}

func (r *CategoryRepository) GetBySlug(slug string) (models.Category, error) {
	var category models.Category
	result := r.db.Where("slug = ?", slug).First(&category)
	return category, result.Error
}

func (r *CategoryRepository) Create(category models.Category) error {
	result := r.db.Model(&models.Category{}).Create(&category)
	return result.Error
}

func (r *CategoryRepository) Update(slug string, category models.UpdateCategoryInput) error {
	result := r.db.Model(models.Category{}).Where("slug = ?", slug).Updates(&models.Category{Name: *category.Name, Slug: category.Slug, Icon: *category.Icon})
	return result.Error
}

func (r *CategoryRepository) Delete(slug string) error {
	result := r.db.Where("slug = ?", slug).Delete(&models.Category{})
	return result.Error
}

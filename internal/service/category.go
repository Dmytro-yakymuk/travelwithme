package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type CategoryService struct {
	categoryRepository repository.Category
}

func NewCategoryService(categoryRepository repository.Category) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.categoryRepository.GetAll()
}
func (s *CategoryService) GetById(id int) (models.Category, error) {
	return s.categoryRepository.GetById(id)
}
func (s *CategoryService) GetBySlug(slug string) (models.Category, error) {
	return s.categoryRepository.GetBySlug(slug)
}
func (s *CategoryService) Create(category models.Category) error {
	return s.categoryRepository.Create(category)
}
func (s *CategoryService) Update(slug string, category models.UpdateCategoryInput) error {
	return s.categoryRepository.Update(slug, category)
}
func (s *CategoryService) Delete(slug string) error {
	return s.categoryRepository.Delete(slug)
}

package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type RegionService struct {
	regionRepository repository.Region
}

func NewRegionService(regionRepository repository.Region) *RegionService {
	return &RegionService{regionRepository: regionRepository}
}

func (s *RegionService) GetAll() ([]models.Region, error) {
	return s.regionRepository.GetAll()
}
func (s *RegionService) GetById(id int) (models.Region, error) {
	return s.regionRepository.GetById(id)
}
func (s *RegionService) GetBySlug(slug string) (models.Region, error) {
	return s.regionRepository.GetBySlug(slug)
}
func (s *RegionService) Create(region models.Region) error {
	return s.regionRepository.Create(region)
}
func (s *RegionService) Update(slug string, region models.UpdateRegionInput) error {
	return s.regionRepository.Update(slug, region)
}
func (s *RegionService) Delete(slug string) error {
	return s.regionRepository.Delete(slug)
}

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

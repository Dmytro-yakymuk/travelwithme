package service

import (
	"os"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type ImageService struct {
	imageRepository repository.Image
}

func NewImageService(imageRepository repository.Image) *ImageService {
	return &ImageService{imageRepository: imageRepository}
}

func (s *ImageService) GetAllImageByTourId(tour_id int) ([]models.Image, error) {
	return s.imageRepository.GetAllImageByTourId(tour_id)
}

func (s *ImageService) GetMainNameImageByTourId(tour_id int) (string, error) {
	return s.imageRepository.GetMainNameImageByTourId(tour_id)
}

func (s *ImageService) Create(tour_id int, name string) error {
	return s.imageRepository.Create(tour_id, name)
}

func (s *ImageService) Delete(id int) error {
	image, err := s.imageRepository.Delete(id)
	if err != nil && image.Name != "" {
		return err
	}

	pwd, _ := os.Getwd()
	err = os.Remove(pwd + "/static/img/tours/" + image.Name)
	if err != nil {
		return err
	}

	return nil
}

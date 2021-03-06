package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(email, password string) (map[string]interface{}, error)
	ParseToken(token string) (map[string]interface{}, error)
}

type Tour interface {
	GetAll(urlQuery map[string][]string) ([]models.Tour, error)
	GetBySlug(slug string) (models.Tour, error)
	Create(tour models.Tour) error
	Update(slug string, tour models.UpdateTourInput) error
	Delete(slug string) error
}

type Category interface {
	GetAll() ([]models.Category, error)
}

type Region interface {
	GetAll() ([]models.Region, error)
}

type Image interface {
	GetAllImageByTourId(tour_id int) ([]models.Image, error)
	Create(tour_id int, name string) error
	Delete(id int) error
}

type User interface {
	GetAll(urlQuery map[string][]string) ([]models.User, error)
}

type Trip interface {
	GetAllTripByTourId(tour_id int) ([]models.Trip, error)
}

type Service struct {
	Authorization Authorization
	Tour          Tour
	Category      Category
	Region        Region
	Image         Image
	User          User
	Trip          Trip
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repository.Authorization),
		Tour:          NewTourService(repository.Tour),
		Category:      NewCategoryService(repository.Category),
		Region:        NewRegionService(repository.Region),
		Image:         NewImageService(repository.Image),
		User:          NewUserService(repository.User),
		Trip:          NewTripService(repository.Trip),
	}
}

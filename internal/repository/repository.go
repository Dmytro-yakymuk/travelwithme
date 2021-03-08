package repository

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository/mdb"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(email, password string) (models.User, error)
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
	Delete(id int) (models.Image, error)
}

type User interface {
	GetAll(urlQuery map[string][]string) ([]models.User, error)
}

type Trip interface {
	GetAllTripByTourId(tour_id int) ([]models.Trip, error)
}

type Order interface {
	GetAll(id int) ([]models.GetOrderToCart, error)
	GetAllOrderByTourId(tour_id int) ([]models.Order, error)
	Create(order models.Order) error
}

type Repository struct {
	Authorization Authorization
	Tour          Tour
	Category      Category
	Region        Region
	Image         Image
	User          User
	Trip          Trip
	Order         Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: mdb.NewAuthorizationRepository(db),
		Tour:          mdb.NewTourRepository(db),
		Category:      mdb.NewCategoryRepository(db),
		Region:        mdb.NewRegionRepository(db),
		Image:         mdb.NewImageRepository(db),
		User:          mdb.NewUserRepository(db),
		Trip:          mdb.NewTripRepository(db),
		Order:         mdb.NewOrderRepository(db),
	}
}

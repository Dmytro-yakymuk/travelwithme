package repository

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository/mdb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(email, password string) models.UserOutput
}

type Tour interface {
	GetAll(urlQuery map[string][]string) ([]models.Tour, error)
	GetById(id int) (models.Tour, error)
	GetBySlug(slug string) (models.Tour, error)
	GetByCategoryID(category_id int) ([]models.Tour, error)
	GetByRegionID(region_id int) ([]models.Tour, error)
	GetByEventID(event_id int) ([]models.ToursEvent, error)
	GetByUserID(user_id int) ([]models.Tour, error)
	Create(tour models.Tour) (int, error)
	Update(slug string, tour models.UpdateTourInput) error
	Public(slug string, public bool) error
	Activ(slug string, activ bool) error
	Delete(slug string) error
}

type Category interface {
	GetAll() ([]models.Category, error)
	GetById(id int) (models.Category, error)
	GetBySlug(slug string) (models.Category, error)
	Create(category models.Category) error
	Update(slug string, category models.UpdateCategoryInput) error
	Delete(slug string) error
}

type Region interface {
	GetAll() ([]models.Region, error)
	GetById(id int) (models.Region, error)
	GetBySlug(slug string) (models.Region, error)
	Create(region models.Region) error
	Update(slug string, region models.UpdateRegionInput) error
	Delete(slug string) error
}

type Image interface {
	GetAllImageByTourId(tour_id int) ([]models.Image, error)
	GetMainNameImageByTourId(tour_id int) (string, error)
	Create(tour_id int, name string) error
	Delete(id int) (models.Image, error)
}

type User interface {
	GetAll(urlQuery map[string][]string) ([]models.User, error)
	GetOneByID(id int) (models.User, error)
	Update(userID int, user models.User) error
}

type Trip interface {
	GetAllTripsWhichAttach(id []string) ([]models.TripToCart, error)
	GetAllTripByTourId(tour_id int) ([]models.Trip, error)
	CheckFreeCountTrip(id string) (int, error)

	// adminka
	GetAllTripsByUserId(userId int) ([]models.Trip, error)
	GetOneByID(id int) (models.Trip, error)
	Create(trip *models.Trip) error
	Update(id int, trip *models.Trip) error
	Delete(id int) error
}

type Order interface {
	GetAllOrdersByUserIdTrip(userID int, role string) ([]models.OrderOutput, error)
	GetOneByID(orderID uuid.UUID) (models.OrderOutput, error)
	GetOneOrderByUserIdTrip(userID int, orderID uuid.UUID) (models.Order, error)
	UpdatePaid(id uuid.UUID, paid bool) error
	Create(order models.Order) (uuid.UUID, error)
}

type TripsOrder interface {
	Create(countTrip models.TripsOrder) error
	GetAllTripsOrderByOrderId(orderID uuid.UUID) ([]models.TripsOrderOutput, error)
}

type Event interface {
	GetAllEventByTourId(tourId int) ([]models.Event, error)
	GetAll() ([]models.Event, error)
	GetByID(id int) (models.Event, error)
	Create(event models.Event) error
	Update(id int, event models.UpdateEventInput) error
	Delete(id int) error
}

type Comment interface {
	Get(urlQuery map[string][]string) ([]models.Comment, error)
	GetAllCommentsByTourId(tourId int) ([]models.Comment, error)
	GetOneByID(id int) ([]models.Comment, error)
	Create(comment models.Comment) error
	Update(id int, comment models.Comment) error
	Delete(id int) error
}

type Audit interface {
	GetAllAuditsByTourId(tourId int) ([]models.Audit, error)
	Create(models.Audit) (*models.Audit, error)
	Delete(id int) error
}

type ToursEvent interface {
	Create(toursEvent []*models.ToursEvent) error
	Delete(tourID int) error
}

type Role interface {
	GetOneByID(id int) (models.Role, error)
}

type Repository struct {
	Authorization Authorization
	Tour          Tour
	Category      Category
	Region        Region
	Image         Image
	User          User
	Trip          Trip
	TripsOrder    TripsOrder
	Order         Order
	Event         Event
	Comment       Comment
	Audit         Audit
	ToursEvent    ToursEvent
	Role          Role
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
		TripsOrder:    mdb.NewTripsOrderRepository(db),
		Order:         mdb.NewOrderRepository(db),
		Event:         mdb.NewEventRepository(db),
		Comment:       mdb.NewCommentRepository(db),
		Audit:         mdb.NewAuditRepository(db),
		ToursEvent:    mdb.NewToursEventRepository(db),
		Role:          mdb.NewRoleRepository(db),
	}
}

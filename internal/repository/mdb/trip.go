package mdb

import (
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) GetAllTripsWhichAttach(id []string) ([]models.TripToCart, error) {
	var trips []models.TripToCart
	result := r.db.Table(
		"trips",
	).Select(
		"trips.*", "tours.title", "tours.slug", "tours.id as tour_id",
	).Joins(
		"LEFT JOIN tours ON trips.tour_id = tours.id",
	).Where(
		"trips.id IN ?", id,
	).Find(
		&trips,
	)
	return trips, result.Error
}

func (r *TripRepository) GetAllTripByTourId(tour_id int) ([]models.Trip, error) {
	var trips []models.Trip

	result := r.db.Where("tour_id = ? AND start <= ?", tour_id, time.Now().AddDate(0, 0, 7)).Find(&trips)
	return trips, result.Error
}

func (r *TripRepository) CheckFreeCountTrip(id string) (int, error) {
	var trip models.Trip
	var tripsOrder []models.TripsOrder
	result := r.db.Select("count").First(&trip, id)
	result2 := r.db.Select("sum(trips_orders.count) as count").Where("trip_id = ?", id).Find(&tripsOrder)
	if result2.Error != nil {
		return trip.Count, result2.Error
	}

	freeCount := trip.Count - tripsOrder[0].Count
	return freeCount, result.Error
}

func (r *TripRepository) GetAllTripsByUserId(userId int) ([]models.Trip, error) {
	var trips []models.Trip

	result := r.db.Joins("left join tours on trips.tour_id = tours.id")
	if userId != 0 {
		result.Where("tours.user_id = ?", userId)
	}
	result.Find(&trips)
	return trips, result.Error
}

func (r *TripRepository) GetOneByID(id int) (models.Trip, error) {
	var trip models.Trip
	result := r.db.First(&trip, id)
	return trip, result.Error
}

func (r *TripRepository) Create(trip *models.Trip) error {
	result := r.db.Model(&models.Trip{}).Create(&trip)
	return result.Error
}

func (r *TripRepository) Update(id int, trip *models.Trip) error {
	result := r.db.Model(models.Trip{}).Where("id = ?", id).Updates(&trip)
	return result.Error
}

func (r *TripRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Delete(&models.Trip{})
	return result.Error
}

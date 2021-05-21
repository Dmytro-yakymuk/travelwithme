package service

import (
	"strconv"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type TripService struct {
	tripRepository repository.Trip
}

func NewTripService(tripRepository repository.Trip) *TripService {
	return &TripService{tripRepository: tripRepository}
}

func (s *TripService) GetAllTripsWhichAttach(id []string) ([]models.TripToCart, error) {
	return s.tripRepository.GetAllTripsWhichAttach(id)
}

func (s *TripService) GetAllTripByTourId(c *gin.Context, rdb *redis.Client, tour_id int) (map[string]interface{}, error) {

	trips, err := s.tripRepository.GetAllTripByTourId(tour_id)
	if err != nil {
		return nil, err
	}
	var mapTrips []interface{}
	minPrice := 0.0
	countGroup := 0
	for _, trip := range trips {
		trip_id := strconv.Itoa(trip.Id)
		freeCount, err := s.CheckCountTrip(c, rdb, trip_id)
		if err != nil {
			return nil, err
		}
		mapTrips = append(mapTrips, map[string]interface{}{
			"id":        trip.Id,
			"start":     trip.Start,
			"end":       trip.End,
			"price":     trip.Price,
			"count":     trip.Count,
			"freeCount": freeCount,
		})

		minPrice += trip.Price
		if countGroup < trip.Count {
			countGroup = trip.Count
		}
	}
	lenTrips := float64(len(trips))
	if lenTrips <= 1 {
		minPrice = minPrice / 1
	} else {
		minPrice = minPrice / lenTrips
	}

	return map[string]interface{}{
		"mapTrips":   mapTrips,
		"minPrice":   minPrice,
		"countGroup": countGroup,
	}, nil
}

// по id trip перевіряємо і повертаємо кількість вільних місць
func (s *TripService) CheckCountTrip(c *gin.Context, rdb *redis.Client, id string) (int, error) {
	// дивимось скільки взагалі місць має поїздка
	count, err := s.tripRepository.CheckFreeCountTrip(id)
	if err != nil {
		return 0, err
	}

	// вибираєм клієнтів
	keys, err := rdb.Keys(c, "trips_*").Result()
	if err != nil {
		return 0, err
	}

	for _, key := range keys {

		ok, err := rdb.HExists(c, key, id).Result()
		if err != nil {
			return 0, err
		}

		if ok {
			// дивимось кількість зарезервованих місць клієнтом для кожної вибраної ним поїздки
			countAttach, err := rdb.HGet(c, key, id).Result()
			if err != nil {
				return 0, err
			}

			countAttachInt, err := strconv.Atoi(countAttach)
			if err != nil {
				return 0, err
			}

			// відмінусовуємо від загальної кількості
			count -= countAttachInt

		}
	}

	return count, nil
}

func (s *TripService) GetAllTripsByUserId(userId int) ([]models.Trip, error) {
	return s.tripRepository.GetAllTripsByUserId(userId)
}
func (s *TripService) GetOneByID(id int) (models.Trip, error) {
	return s.tripRepository.GetOneByID(id)
}
func (s *TripService) Create(trip *models.Trip) error {
	return s.tripRepository.Create(trip)
}
func (s *TripService) Update(id int, trip *models.Trip) error {
	return s.tripRepository.Update(id, trip)
}
func (s *TripService) Delete(id int) error {
	return s.tripRepository.Delete(id)
}

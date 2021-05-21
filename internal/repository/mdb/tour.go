package mdb

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type TourRepository struct {
	db *gorm.DB
}

func NewTourRepository(db *gorm.DB) *TourRepository {
	return &TourRepository{db: db}
}

func (r *TourRepository) GetAll(urlQuery map[string][]string) ([]models.Tour, error) {
	var tours []models.Tour

	result := r.db.Model(&tours)

	if urlQuery["category_id"] != nil {
		result.Where("category_id = ?", urlQuery["category_id"][0])
	}

	if urlQuery["region_id"] != nil {
		result.Where("region_id = ?", urlQuery["region_id"][0])
	}

	if urlQuery["user_id"] != nil && urlQuery["user_id"][0] != "0" {
		result.Where("user_id = ?", urlQuery["user_id"][0])
	}

	if urlQuery["price"] != nil {
		price := urlQuery["price"][0]

		priceSplit := strings.Split(price, "-")
		if len(priceSplit) != 2 {
			return nil, errors.New("bad request, error split")
		}

		result.Select("tours.*").Joins("left join trips on tours.id = trips.tour_id").Where("trips.price >= ? AND trips.price < ?", priceSplit[0], priceSplit[1]).Group("trips.tour_id")
	}

	if urlQuery["date"] != nil {
		date := urlQuery["date"][0]

		dateSplit := strings.Split(date, " ")
		if len(dateSplit) != 2 {
			return nil, errors.New("bad request, error split")
		}
		result.Select("tours.*").Joins("left join trips on tours.id = trips.tour_id").Where("trips.start >= ? AND trips.end < ?", dateSplit[0], dateSplit[1]).Group("trips.tour_id")
	}

	if urlQuery["order_by"] != nil {
		order := urlQuery["order_by"][0]
		result.Order(order)
	}

	if urlQuery["limit"] != nil {
		limit, err := strconv.Atoi(urlQuery["limit"][0])
		if err != nil {
			return nil, err
		}
		result.Limit(limit)
	}

	result.Find(&tours)
	return tours, result.Error
}

func (r *TourRepository) GetBySlug(slug string) (models.Tour, error) {
	var tour models.Tour
	result := r.db.First(&tour, "slug = ?", slug)
	return tour, result.Error
}

func (r *TourRepository) GetById(id int) (models.Tour, error) {
	var tour models.Tour
	result := r.db.First(&tour, "id = ?", id)
	return tour, result.Error
}

func (r *TourRepository) GetByCategoryID(category_id int) ([]models.Tour, error) {
	var tours []models.Tour
	result := r.db.Find(&tours, "category_id = ?", category_id)
	return tours, result.Error
}

func (r *TourRepository) GetByRegionID(region_id int) ([]models.Tour, error) {
	var tours []models.Tour
	result := r.db.Find(&tours, "region_id = ?", region_id)
	return tours, result.Error
}

func (r *TourRepository) GetByEventID(event_id int) ([]models.ToursEvent, error) {
	var toursEvents []models.ToursEvent
	result := r.db.Find(&toursEvents, "event_id = ?", event_id)
	return toursEvents, result.Error
}

func (r *TourRepository) GetByUserID(user_id int) ([]models.Tour, error) {
	var tours []models.Tour
	result := r.db.Find(&tours, "user_id = ?", user_id)
	return tours, result.Error
}

func (r *TourRepository) Create(tour models.Tour) (int, error) {
	result := r.db.Model(&models.Tour{}).Create(&tour)
	return tour.Id, result.Error
}

func (r *TourRepository) Update(slug string, tour models.UpdateTourInput) error {
	var tourModel models.Tour
	result := r.db.Model(tourModel).Where("slug = ?", slug).Updates(&models.Tour{Title: *tour.Title, Slug: *tour.Slug, Description: *tour.Description, Text: *tour.Text, X: *tour.X, Y: *tour.Y, CategoryId: *tour.CategoryId, RegionId: *tour.RegionId})
	return result.Error
}

func (r *TourRepository) Public(slug string, public bool) error {
	result := r.db.Model(models.Tour{}).Where("slug = ?", slug).Update("public", public)
	return result.Error
}

func (r *TourRepository) Activ(slug string, activ bool) error {
	result := r.db.Model(models.Tour{}).Where("slug = ?", slug).Update("activ", activ)
	return result.Error
}

func (r *TourRepository) Delete(slug string) error {
	result := r.db.Where("slug = ?", slug).Delete(&models.Tour{})
	return result.Error
}

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

	if urlQuery["where"] != nil {
		where := urlQuery["where"][0]

		whereSplit := strings.Split(where, "=")
		if len(whereSplit) != 2 {
			return nil, errors.New("bad request, error split")
		}

		result.Where(whereSplit[0]+"= ?", whereSplit[1])
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

func (r *TourRepository) Create(tour models.Tour) error {
	result := r.db.Model(&models.Tour{}).Create(&tour)
	return result.Error
}

func (r *TourRepository) Update(slug string, tour models.UpdateTourInput) error {
	result := r.db.Model(models.Tour{}).Where("slug = ?", slug).Updates(&models.Tour{Title: *tour.Title, Slug: *tour.Slug, Description: *tour.Description, CategoryId: *tour.CategoryId, RegionId: *tour.RegionId})
	return result.Error
}

func (r *TourRepository) Delete(slug string) error {
	result := r.db.Where("slug = ?", slug).Delete(&models.Tour{})
	return result.Error
}

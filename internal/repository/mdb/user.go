package mdb

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(urlQuery map[string][]string) ([]models.User, error) {
	var users []models.User

	result := r.db.Model(&users)

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

	result.Find(&users)
	return users, result.Error
}

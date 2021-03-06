package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type AuthorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) CreateUser(user models.User) error {
	// result := r.db.Create(user)
	result := r.db.Model(&models.User{}).Create(map[string]interface{}{
		"name": user.Name, "email": user.Email, "password": user.Password, "role": user.Role,
	})
	return result.Error
}

func (r *AuthorizationRepository) GetUser(email, password string) (models.User, error) {
	var user models.User
	result := r.db.Select("id", "role").Where("email = ? AND password = ?", email, password).First(&user)
	return user, result.Error
}

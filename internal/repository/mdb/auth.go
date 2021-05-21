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

	// result := r.db.Model(&models.User{}).Create(map[string]interface{}{
	// 	"name": user.Name, "email": user.Email, "password": user.Password, "role": user.Role,
	// })
	// return result.Error
	return nil
}

func (r *AuthorizationRepository) GetUser(email, password string) models.UserOutput {
	var userOutput models.UserOutput
	r.db.Select(
		"users.id", "roles.name as role",
	).Table(
		"users",
	).Where(
		"users.email = ? AND users.password = ?", email, password,
	).Joins(
		"left join roles on users.role_id = roles.id",
	).First(
		&userOutput,
	)

	return userOutput
}

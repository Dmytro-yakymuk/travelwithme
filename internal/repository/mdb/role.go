package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetOneByID(id int) (models.Role, error) {
	var role models.Role
	result := r.db.First(&role, id)
	return role, result.Error
}

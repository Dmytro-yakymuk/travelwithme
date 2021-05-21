package mdb

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"gorm.io/gorm"
)

type AuditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{db: db}
}

func (r *AuditRepository) GetAllAuditsByTourId(tourId int) ([]models.Audit, error) {
	var audits []models.Audit

	result := r.db.Find(&audits).Where("tour_id = ?", tourId)
	return audits, result.Error
}

func (r *AuditRepository) Create(audit models.Audit) (*models.Audit, error) {
	result := r.db.Model(&models.Audit{}).Create(&audit)
	return &audit, result.Error
}

func (r *AuditRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Delete(&models.Audit{})
	return result.Error
}

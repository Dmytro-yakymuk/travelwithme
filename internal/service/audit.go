package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type AuditService struct {
	auditRepository repository.Audit
}

func NewAuditService(auditRepository repository.Audit) *AuditService {
	return &AuditService{auditRepository: auditRepository}
}

func (s *AuditService) GetAllAuditsByTourId(tourId int) ([]models.Audit, error) {
	return s.auditRepository.GetAllAuditsByTourId(tourId)
}

func (s *AuditService) Create(audit models.Audit) (*models.Audit, error) {
	return s.auditRepository.Create(audit)
}

func (s *AuditService) Delete(id int) error {
	return s.auditRepository.Delete(id)
}

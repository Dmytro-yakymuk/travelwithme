package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type RoleService struct {
	roleRepository repository.Role
}

func NewRoleService(roleRepository repository.Role) *RoleService {
	return &RoleService{roleRepository: roleRepository}
}

func (s *RoleService) GetOneByID(id int) (models.Role, error) {
	return s.roleRepository.GetOneByID(id)
}

package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type UserService struct {
	userRepository repository.User
}

func NewUserService(userRepository repository.User) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAll(urlQuery map[string][]string) ([]models.User, error) {
	return s.userRepository.GetAll(urlQuery)
}

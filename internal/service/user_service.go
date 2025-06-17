package service

import (
	"final-project/internal/domain"
	"final-project/internal/repository"
)

type UserService interface {
	GetUserByID(userID uint) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetUserByID(userID uint) (*domain.User, error) {
	return s.repo.FindByID(userID)
}

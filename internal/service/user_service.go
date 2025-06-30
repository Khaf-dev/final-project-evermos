package service

import (
	"final-project/internal/domain"
	"final-project/internal/helpers"
	"final-project/internal/repository"
)

type UserService interface {
	GetUserByID(userID uint) (*domain.User, error)
	UpdateUser(userID uint, input domain.UpdateUserRequest) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{userRepo: r}
}

func (s *userService) GetUserByID(userID uint) (*domain.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *userService) UpdateUser(userID uint, input domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	if input.Password != "" {
		hashed, _ := helpers.HashedPassword(input.Password)
		user.Password = hashed
	}

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

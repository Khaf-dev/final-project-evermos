package service

import (
	"final-project/internal/domain"
	"final-project/internal/repository"
)

type AddressService interface {
	CreateAddress(userID uint, input domain.Address) error
	GetUserAddresses(userID uint) ([]domain.Address, error)
}

type addressService struct {
	repo repository.AddressRepository
}

func NewAddressService(r repository.AddressRepository) AddressService {
	return &addressService{repo: r}
}

func (s *addressService) CreateAddress(userID uint, input domain.Address) error {
	input.UserID = userID
	return s.repo.Create(&input)
}

func (s *addressService) GetUserAddresses(userID uint) ([]domain.Address, error) {
	return s.repo.FindByUserID(userID)
}

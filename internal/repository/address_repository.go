package repository

import (
	"final-project/internal/domain"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *domain.Address) error
	FindByUserID(userID uint) ([]domain.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

func (r *addressRepository) Create(address *domain.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) FindByUserID(userID uint) ([]domain.Address, error) {
	var addresses []domain.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

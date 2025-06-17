package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmailOrPhone(email, phone string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("email = ? OR phone = ?", email, phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmailOrPhoneLogin(input string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("email = ? OR phone = ?", input, input).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type userRepository interface {
	FindByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmailOrPhone(email, phone string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ? OR phone = ?", email, phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmailOrPhoneLogin(input string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ? OR phone = ?", input, input).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

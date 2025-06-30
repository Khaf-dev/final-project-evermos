package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) Create(store *domain.Store) error {
	return r.db.Create(store).Error
}

func (r *StoreRepository) FindByUserID(userID uint) (*domain.Store, error) {
	var store domain.Store
	if err := r.db.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepository) FindByID(id uint) (*domain.Store, error) {
	var store domain.Store
	err := r.db.First(&store, id).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type StoreRepository struct {
	DB *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{DB: db}
}

func (r *StoreRepository) Create(store *domain.Store) error {
	return r.DB.Create(store).Error
}

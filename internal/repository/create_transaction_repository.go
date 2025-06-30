package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type TransactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) *TransactionDetailRepository {
	return &TransactionDetailRepository{db: db}
}

func (r *TransactionDetailRepository) Create(detail *domain.TransactionDetail) error {
	return r.db.Create(detail).Error
}

func (r *TransactionDetailRepository) BulkCreate(details []domain.TransactionDetail) error {
	return r.db.Create(&details).Error
}

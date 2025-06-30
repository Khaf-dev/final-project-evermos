package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(tx *domain.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *TransactionRepository) CreateDetails(details []domain.TransactionDetail) error {
	return r.db.Create(&details).Error
}

func (r *TransactionRepository) FindByID(id uint) (*domain.Transaction, error) {
	var tx domain.Transaction

	if err := r.db.
		Preload("User").
		Preload("Store").
		Preload("Details").
		Preload("Details.Product").
		Preload("Details.Product.Store").
		Preload("Details.Product.Category").
		First(&tx, id).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *TransactionRepository) FindByUserID(userID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := r.db.
		Where("user_id = ?", userID).
		Preload("Details").
		Preload("Details.Product").
		Preload("Details.Product.Category").
		Preload("Details.Product.Store").
		Preload("User").
		Preload("Store").
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) FindByStoreID(storeID uint) ([]domain.Transaction, error) {
	var transaction []domain.Transaction

	if err := r.db.
		Preload("User").
		Preload("Store").
		Preload("Details").
		Preload("Details.Product").
		Preload("Details.Product.Store").
		Preload("Details.Product.Category").
		Where("store_id = ?", storeID).
		Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) FindALl() ([]domain.Transaction, error) {
	var transaction []domain.Transaction
	err := r.db.
		Preload("Details.Product").
		Find(&transaction).Error
	return transaction, err
}

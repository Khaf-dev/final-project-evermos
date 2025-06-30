package repository

import (
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type ProductLogRepository interface {
	Create(log *domain.ProductLog) error
	FindAll() ([]domain.ProductLog, error)
	FindByProductID(productID uint) ([]domain.ProductLog, error)
}

type productLogRepository struct {
	db *gorm.DB
}

func NewProductLogRepository(db *gorm.DB) ProductLogRepository {
	return &productLogRepository{db}
}

func (r *productLogRepository) Create(log *domain.ProductLog) error {
	return r.db.Create(log).Error
}
func (r *productLogRepository) FindAll() ([]domain.ProductLog, error) {
	var logs []domain.ProductLog
	err := r.db.Find(&logs).Error
	return logs, err
}
func (r *productLogRepository) FindByProductID(productID uint) ([]domain.ProductLog, error) {
	var logs []domain.ProductLog
	err := r.db.Where("product_id = ?", productID).Find(&logs).Error
	return logs, err
}

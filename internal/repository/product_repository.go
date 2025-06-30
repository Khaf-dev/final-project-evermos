package repository

import (
	"errors"
	"final-project/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.
		Preload("Store").
		Preload("Category").
		First(&product, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindAll(filter string, categoryID uint, limit int, offset int) ([]domain.Product, error) {
	var products []domain.Product

	if err := r.db.Preload("Store").Find(&products).Error; err != nil {
		return nil, err
	}

	query := r.db.Preload("Category")

	if filter != "" {
		query = query.Where("name LIKE ?", "%"+filter+"%")
	}
	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindBy(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) FindAllFiltered(name string, categoryID uint, page int, limit int) ([]domain.Product, error) {
	var products []domain.Product
	query := r.db

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	offset := (page + 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&products).Error

	return products, err
}

func (r *ProductRepository) GetStoreIDByProduct(productID uint) (uint, error) {
	var product domain.Product
	if err := r.db.Select("store_id").First(&product, productID).Error; err != nil {
		return 0, err
	}
	return product.StoreID, nil
}

func (r *ProductRepository) FindByStoreID(storeID uint) ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Where("store_id = ?", storeID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Update(product *domain.Product) error {
	return r.db.Save(&product).Error
}

func (r *ProductRepository) Delete(product *domain.Product) error {
	return r.db.Delete(product).Error
}

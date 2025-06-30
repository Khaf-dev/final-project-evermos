package service

import (
	"errors"
	"final-project/dto/request"
	"final-project/internal/domain"
	"final-project/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(input request.CreateProductRequest, storeID uint) (*domain.Product, error) {
	product := domain.Product{

		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		CategoryID:  input.CategoryID,
		ImageURL:    input.ImageURL,
		StoreID:     storeID,
	}
	err := s.repo.Create(&product)
	if err != nil {
		return nil, err
	}

	fullProduct, err := s.repo.FindByID(product.ID)
	if err != nil {
		return nil, err
	}

	return fullProduct, nil
}

func (s *ProductService) GetAllProduct(name string, categoryID uint, page int, limit int) ([]domain.Product, error) {
	offset := (page - 1) * limit
	return s.repo.FindAll(name, categoryID, offset, limit)
}

func (s *ProductService) GetProductByStore(storeID uint) ([]domain.Product, error) {
	return s.repo.FindByStoreID(storeID)
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetAllProductFiltered(name string, categoryID uint, page int, limit int) ([]domain.Product, error) {
	return s.repo.FindAllFiltered(name, categoryID, page, limit)
}

func (s *ProductService) IsOwner(productId, storeID uint) (bool, error) {
	productStoreID, err := s.repo.GetStoreIDByProduct(productId)
	if err != nil {
		return false, err
	}
	return productStoreID == storeID, nil
}

func (s *ProductService) UpdateProduct(id uint, input request.UpdateProductRequest, storeID uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("product tidak ditemukan")
	}

	// Validasi kepemilikan
	if product.StoreID != storeID {
		return errors.New("akses ditolak")
	}

	// Update data
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.ImageURL = input.ImageURL

	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint, storeID uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	if product.StoreID != storeID {
		return errors.New("akses ditolak")
	}

	return s.repo.Delete(product)
}

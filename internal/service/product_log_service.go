package service

import (
	"final-project/internal/domain"
	"final-project/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type ProductLogService interface {
	GetAll() ([]domain.ProductLog, error)
	GetLogsByProductID(productID uint, userID uint) ([]domain.ProductLog, error)
}

type productLogService struct {
	repo        repository.ProductLogRepository
	productRepo repository.ProductRepository
	storeRepo   repository.StoreRepository
}

func NewProductLogService(repo repository.ProductLogRepository, p repository.ProductRepository, s repository.StoreRepository) ProductLogService {
	return &productLogService{
		repo:        repo,
		productRepo: p,
		storeRepo:   s,
	}
}

func (s *productLogService) GetAll() ([]domain.ProductLog, error) {
	return s.repo.FindAll()
}
func (s *productLogService) GetLogsByProductID(productID uint, userID uint) ([]domain.ProductLog, error) {
	product, err := s.productRepo.FindByID(productID)
	if err != nil || product == nil {
		return nil, err
	}

	store, err := s.storeRepo.FindByID(product.StoreID)
	if err != nil || store == nil {
		return nil, err
	}

	if store.UserID != userID {
		return nil, fiber.NewError(fiber.StatusForbidden, "Akses ditolak ke log produk ini")
	}

	return s.repo.FindByProductID(productID)
}

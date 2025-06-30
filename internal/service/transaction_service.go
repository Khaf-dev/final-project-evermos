package service

import (
	"final-project/dto/request"
	"final-project/internal/domain"
	"final-project/internal/repository"
	"fmt"
)

type TransactionService struct {
	repo        repository.TransactionRepository
	productRepo repository.ProductRepository
	detailRepo  repository.TransactionDetailRepository
	logRepo     repository.ProductLogRepository
}

func NewTransactionService(
	repo repository.TransactionRepository,
	productRepo repository.ProductRepository,
	detailRepo repository.TransactionDetailRepository,
	logRepo repository.ProductLogRepository,
) *TransactionService {
	return &TransactionService{
		repo:        repo,
		productRepo: productRepo,
		detailRepo:  detailRepo,
		logRepo:     logRepo,
	}
}

func (s *TransactionService) CreateTransaction(userID uint, req request.CreateTransactionRequest) (*domain.Transaction, error) {
	// 1. Buat object transaction (masih kosong ID-nya)
	transaction := &domain.Transaction{
		UserID:  userID,
		StoreID: req.StoreID,
		Status:  "pending",
	}

	var details []domain.TransactionDetail
	var logs []domain.ProductLog
	var total float64

	for _, item := range req.Items {
		product, err := s.productRepo.FindByID(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("error saat cari produk ID %d: %v", item.ProductID, err)
		}
		if product == nil {
			return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", item.ProductID)
		}

		// Validasi: produk harus dari store yang sesuai
		if product.StoreID != req.StoreID {
			return nil, fmt.Errorf("semua produk harus berasal dari store yang sama dengan StoreID yang dikirim")
		}

		// Validasi stok cukup
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("stok tidak mencukupi untuk produk: %s", product.Name)
		}

		// Update stok
		prevStock := product.Stock
		product.Stock -= item.Quantity

		if err := s.productRepo.Update(product); err != nil {
			return nil, fmt.Errorf("gagal update stock produk: %v", err)
		}

		subTotal := float64(item.Quantity) * product.Price
		total += subTotal

		// Buat detail transaksi
		detail := domain.TransactionDetail{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			Subtotal:  subTotal,
		}
		details = append(details, detail)

		// Siapkan log
		logs = append(logs, domain.ProductLog{
			ProductID:      product.ID,
			ProductName:    product.Name,
			PreviousStock:  prevStock,
			PurchasedQty:   item.Quantity,
			RemainingStock: product.Stock,
			Activity:       "Purchase",
			Detail:         fmt.Sprintf("Pembelian produk %s sebanyak %d unit", product.Name, item.Quantity),
		})
	}

	// Set total transaksi
	transaction.Total = total

	// Simpan transaksi (baru punya ID)
	if err := s.repo.Create(transaction); err != nil {
		return nil, err
	}

	// Set transaction ID ke detail & log, lalu simpan semuanya
	for i := range details {
		details[i].TransactionID = transaction.ID
	}
	if err := s.detailRepo.BulkCreate(details); err != nil {
		return nil, err
	}

	for i := range logs {
		logs[i].TransactionID = transaction.ID
		if err := s.logRepo.Create(&logs[i]); err != nil {
			return nil, fmt.Errorf("gagal simpan log produk: %v", err)
		}
	}

	// Ambil transaksi lengkap
	fullTransaction, err := s.repo.FindByID(transaction.ID)
	if err != nil {
		return nil, err
	}

	return fullTransaction, nil
}

func (s *TransactionService) GetTransactionByID(id uint) (*domain.Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *TransactionService) GetStoreTransactions(storeID uint) ([]domain.Transaction, error) {
	return s.repo.FindByStoreID(storeID)
}

func (s *TransactionService) GetUserTransactions(userID uint) ([]domain.Transaction, error) {
	return s.repo.FindByUserID(userID)
}

func (s *TransactionService) GetAllTransactions() ([]domain.Transaction, error) {
	return s.repo.FindALl()
}

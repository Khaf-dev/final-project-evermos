package handler

import (
	"final-project/dto/request"
	"final-project/internal/repository"
	"final-project/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Service   *service.TransactionService
	storeRepo *repository.StoreRepository
}

func NewTransactionHandler(s *service.TransactionService, storeRepo *repository.StoreRepository) *TransactionHandler {
	return &TransactionHandler{
		Service:   s,
		storeRepo: storeRepo,
	}
}

// POST /transactions
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req request.CreateTransactionRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := req.Validate(); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Transaksi ID tidak valid"})
	}

	transaction, err := h.Service.CreateTransaction(userID, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":     "Transaksi berhasil dibuat",
		"transaction": transaction,
	})
}

// GET /transactions/admin
func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	role, ok := c.Locals("role").(string)
	if !ok || role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied, hanya admin yang bisa mengakses",
		})
	}

	transactions, err := h.Service.GetAllTransactions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data transaksi",
		})
	}

	return c.JSON(fiber.Map{
		"message":      "Berhasil mengambil semua transaksi",
		"transactions": transactions,
	})
}

// GET /transactions/:id
func (h *TransactionHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	transaction, err := h.Service.GetTransactionByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaksi tidak ditemukan"})
	}

	return c.JSON(transaction)
}

// GET /store/transactions
func (h *TransactionHandler) GetByStore(c *fiber.Ctx) error {
	storeID, ok := c.Locals("store_id").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Store ID tidak valid"})
	}

	transactions, err := h.Service.GetStoreTransactions(storeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message":      "Daftar transaksi toko",
		"transactions": transactions,
	})
}

func (h *TransactionHandler) GetUserTransactions(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID tidak valid dari token",
		})
	}

	transactions, err := h.Service.GetUserTransactions(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mengambil data transaksi",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":      "Berhasil mengambil semua transaksi user",
		"transactions": transactions,
	})
}

func (h *TransactionHandler) GetStoreTransactions(c *fiber.Ctx) error {

	storeID, ok := c.Locals("store_id").(uint)
	fmt.Printf("[HANDLER] store_id: %v (ok: %v)\n", storeID, ok)

	if !ok || storeID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid store ID",
		})
	}

	transactions, err := h.Service.GetStoreTransactions(storeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil transaksi dari toko",
		})
	}

	return c.JSON(fiber.Map{
		"message":      "Berhasil mengambil transaksi dari toko",
		"transactions": transactions,
	})
}

package handler

import (
	"final-project/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductLogHandler struct {
	service service.ProductLogService
}

func NewProductLogHandler(s service.ProductLogService) *ProductLogHandler {
	return &ProductLogHandler{service: s}
}

func (h *ProductLogHandler) GetAll(c *fiber.Ctx) error {
	logs, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil semua log",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil ambil semua log",
		"data":    logs,
	})
}
func (h *ProductLogHandler) GetLogsByProductID(c *fiber.Ctx) error {
	productIDStr := c.Params("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID produk tidak valid",
		})
	}

	userID := c.Locals("user_id").(uint)
	logs, err := h.service.GetLogsByProductID(uint(productID), userID)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(fiber.Map{"error": fiberErr.Message})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal ambil log",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil ambil log produk",
		"data":    logs,
	})
}

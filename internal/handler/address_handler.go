package handler

import (
	"final-project/internal/domain"
	"final-project/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	service service.AddressService
}

func NewAddressHandler(s service.AddressService) *AddressHandler {
	return &AddressHandler{service: s}
}

func (h *AddressHandler) Create(c *fiber.Ctx) error {
	var input domain.Address
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}
	userID := c.Locals("user_id").(uint)
	err := h.service.CreateAddress(userID, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menyimpan alamat",
		})
	}
	return c.JSON(fiber.Map{"message": "Alamat berhasil disimpan"})
}

func (h *AddressHandler) GetByUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	addresses, err := h.service.GetUserAddresses(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil alamat",
		})
	}
	return c.JSON(fiber.Map{"data": addresses})
}

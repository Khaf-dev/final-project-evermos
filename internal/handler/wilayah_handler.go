package handler

import (
	"final-project/internal/service"

	"github.com/gofiber/fiber/v2"
)

type WilayahHandler struct {
	service service.WilayahService
}

func NewWilayahHandler(s service.WilayahService) *WilayahHandler {
	return &WilayahHandler{service: s}
}

func (h *WilayahHandler) GetProvinces(c *fiber.Ctx) error {
	data, err := h.service.GetProvinces()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

func (h *WilayahHandler) GetRegencies(c *fiber.Ctx) error {
	id := c.Params("province_id")
	data, err := h.service.GetRegencies(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

func (h *WilayahHandler) GetDistricts(c *fiber.Ctx) error {
	id := c.Params("regency_id")
	data, err := h.service.GetDistricts(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

func (h *WilayahHandler) GetVillages(c *fiber.Ctx) error {
	id := c.Params("district_id")
	data, err := h.service.GetVillages(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

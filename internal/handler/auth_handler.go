package handler

import (
	"final-project/dto/request"
	"final-project/internal/service"
	"final-project/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error { // Fungsi Registrasi Akun
	var input request.RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  "Invalid Input",
			"debug":  err.Error(),
			"detail": err.Error()})
	}

	if err := utils.Validator.Struct(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Validation failed",
			"debug": err.Error()})
	}

	user, err := h.Service.Register(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Registrasi akun berhasil",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
			"role":  user.Role,
		},
	})

}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input request.LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := utils.Validator.Struct(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.Service.Login(input)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

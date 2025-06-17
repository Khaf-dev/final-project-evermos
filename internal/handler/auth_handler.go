package handler

import (
	"final-project/dto/request"
	"final-project/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

var validate = validator.New()

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input request.RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Registrasi akun berhasil"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input request.LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.Service.Login(input)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

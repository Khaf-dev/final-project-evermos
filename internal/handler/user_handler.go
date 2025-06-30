package handler

import (
	"final-project/internal/domain"
	"final-project/internal/service"
	"final-project/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service service.UserService
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	fmt.Println("JWT Claims:", c.Locals("user_id"))
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: user_id not found in token",
		})
	}
	userID := userIDRaw.(uint)

	var input domain.UpdateUserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := utils.Validator.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	user, err := h.Service.UpdateUser(userID, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update profil",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
		"data":    user,
	})
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not Found"})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"role":  user.Role,
	})
}

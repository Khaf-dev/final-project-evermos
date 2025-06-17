package handler

import (
	"final-project/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service service.UserService
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

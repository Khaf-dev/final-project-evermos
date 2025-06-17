package middleware

import "github.com/gofiber/fiber/v2"

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole")

		if userRole != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Akses ditolak, hanya ditujukan kepada admin saja",
			})
		}

		return c.Next()
	}
}

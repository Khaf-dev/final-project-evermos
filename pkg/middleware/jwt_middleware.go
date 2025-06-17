package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("secretkey_final_project") // sama dengan yang ada di utils/jwt.go

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or Invalid Authorization header",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse Token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or Expired Token",
			})
		}

		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Token Claims",
			})
		}

		userID := uint(claims["user_id"].(float64))
		role := claims["role"].(string)

		// Simpan user_id dan role ke context
		c.Locals("user_id", userID)
		c.Locals("role", role)
		c.Locals("userRole", claims["role"])

		return c.Next()
	}
}
